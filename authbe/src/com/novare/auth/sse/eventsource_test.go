/*
MIT License

Copyright (c) 2020 Clerley Silveira

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/
package sse

import (
	"encoding/json"
	"log"
	"sync"
	"testing"
	"time"
)

func TestMessageEventBroker(t *testing.T) {

	MessageBroker.Run()
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()

		subscriber := NewSubscriber("123456")
		subscriber.AddEvent("0001")

		//Unsubscribe first. Let's see if it works
		MessageBroker.Unsubscribe(subscriber)

		log.Printf("Working on the unsubscribed channel!")

		MessageBroker.Subscribe(subscriber)
		event := <-subscriber.EventChannel
		buf, err := json.Marshal(event)
		if err != nil {
			t.Errorf("The following error occurred:[%s]", err)
		}

		log.Printf("ROUTINE 1 - The event: %s", string(buf))
		//Unsubscribe first. Let's see if it works
		MessageBroker.Unsubscribe(subscriber)

		subscriber = NewSubscriber("123456")
		subscriber.AddEvent("0001")
		MessageBroker.Subscribe(subscriber)

		event = <-subscriber.EventChannel

		buf, err = json.Marshal(event)
		if err != nil {
			t.Errorf("The following error occurred:[%s]", err)
		}

		log.Printf("ROUTINE 1 - The event: %s", string(buf))

	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		event := NewEvent("0001", time.Now(), "This is my data")
		buf, err := json.Marshal(event)
		if err != nil {
			t.Logf("The following error occurred:[%s]", err)
		}
		log.Printf("ROUTINE 2 - Sending : %s", string(buf))

		if !MessageBroker.Publish(event) {
			t.Logf("There was an error publishing the event")
		}

		event = NewEvent("0002", time.Now(), "This is my data")
		buf, err = json.Marshal(event)
		if err != nil {
			t.Errorf("The following error occurred:[%s]", err)
		}
		log.Printf("ROUTINE 2 - Sending : %s", string(buf))
		if !MessageBroker.Publish(event) {
			t.Logf("The second message was not properly transmitted")
		}

		time.Sleep(2 * time.Second)
		event = NewEvent("0001", time.Now(), "This is my data 2")
		buf, err = json.Marshal(event)
		if err != nil {
			t.Errorf("The following error occurred:[%s]", err)
		}
		t.Logf("ROUTINE 2 - Sending : %s", string(buf))

		if !MessageBroker.Publish(event) {
			t.Error("The second message was not properly transmitted")
		}

	}()

	wg.Wait()
	MessageBroker.Close()

}
