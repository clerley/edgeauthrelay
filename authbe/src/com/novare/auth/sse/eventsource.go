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
	"log"
	"sync"
	"time"
)

//Event - An Event that can be sent to the site
type Event struct {
	EventID   string      `json:"eventID"`
	TimeStamp string      `json:"timeStamp"`
	Data      interface{} `json:"data"`
}

//NewEvent - ...
func NewEvent(eventID string, timeStamp time.Time, data interface{}) *Event {
	e := new(Event)
	e.Data = data
	e.EventID = eventID
	e.TimeStamp = timeStamp.Format(time.RFC3339)
	return e
}

//Subscriber - Used to subscribe to event
type Subscriber struct {
	ID           string
	EventIDs     []string
	EventChannel chan Event
	lock         sync.Mutex
	closed       bool
}

//NewSubscriber ...
func NewSubscriber(ID string) *Subscriber {
	sub := new(Subscriber)
	sub.ID = ID
	sub.EventChannel = make(chan Event)
	sub.closed = false
	return sub
}

//IsListeningForEvent - Looks for the event, if it is found return true otherwise, false
func (subscriber *Subscriber) IsListeningForEvent(event Event) bool {

	for i := range subscriber.EventIDs {

		if subscriber.EventIDs[i] == EventAll {
			log.Printf("the subscriber with ID:[%s] is subscribed to all events", subscriber.ID)
			return true
		}

		if subscriber.EventIDs[i] == event.EventID {
			return true
		}
	}

	return false
}

//Publish ... Returns true or false
func (subscriber *Subscriber) Publish(event Event) bool {
	subscriber.lock.Lock()
	defer subscriber.lock.Unlock()

	if subscriber.closed {
		return false
	}

	if subscriber.IsListeningForEvent(event) {
		subscriber.EventChannel <- event
		return true
	}

	return false
}

//AddEvent ... Events the subscriber is interested in listening for request
func (subscriber *Subscriber) AddEvent(eventID string) {
	subscriber.EventIDs = append(subscriber.EventIDs, eventID)
}

//DelEvent ...
func (subscriber *Subscriber) DelEvent(eventID string) {

	idx := -1
	for idx = range subscriber.EventIDs {
		if subscriber.EventIDs[idx] == eventID {
			break
		}
	}

	if idx >= len(subscriber.EventIDs) || idx < 0 {
		return
	}

	subscriber.EventIDs = append(subscriber.EventIDs[0:idx], subscriber.EventIDs[idx+1:]...)

}

//Close ...
func (subscriber *Subscriber) Close() {
	subscriber.lock.Lock()
	defer subscriber.lock.Unlock()

	if subscriber.closed {
		log.Printf("The Subscriber with ID: [%s] has already been closed", subscriber.ID)
		return
	}
	subscriber.closed = true
	close(subscriber.EventChannel)
}

//EventBroker - ...
type EventBroker struct {
	inGressEvents chan Event    //Channel for events coming in
	subscribers   []*Subscriber //Subscribers list
	running       bool          //Is it running
	lock          sync.Mutex    //This is required to protect the changes made to subscribers
}

//Run - This will be self started by the
func (broker *EventBroker) Run() {

	go func() {
		for broker.running {
			select {
			case event := <-broker.inGressEvents:
				for idx := range broker.subscribers {
					subscriber := broker.subscribers[idx]
					//Subscriber will know if it should publish it or not
					subscriber.Publish(event)
				}
			}
		}
	}()
}

//Close - Close the event broker and release all teh channels
func (broker *EventBroker) Close() {
	for i := range broker.subscribers {
		broker.Unsubscribe(broker.subscribers[i])
	}

	close(broker.inGressEvents)
}

//Publish an Event
func (broker *EventBroker) Publish(event *Event) bool {
	broker.inGressEvents <- *event

	return true
}

//Subscribe to Events
func (broker *EventBroker) Subscribe(subscriber *Subscriber) {
	broker.lock.Lock()
	defer broker.lock.Unlock()

	broker.subscribers = append(broker.subscribers, subscriber)

}

//Unsubscribe from Event notifications
func (broker *EventBroker) Unsubscribe(subscriber *Subscriber) {
	broker.lock.Lock()
	defer broker.lock.Unlock()

	idx := -1
	for idx = range broker.subscribers {
		if broker.subscribers[idx].ID == subscriber.ID {
			break
		}
	}

	if idx < 0 || idx >= len(broker.subscribers) {
		log.Printf("The subscriber with ID:[%s] was not found!", subscriber.ID)
		return
	}

	subs := broker.subscribers[idx]
	broker.subscribers = append(broker.subscribers[0:idx], broker.subscribers[idx:]...)
	subs.Close()
}

//MessageBroker - Used to post events to a queue
var MessageBroker *EventBroker

//NewMessageBroker ...
func NewMessageBroker() *EventBroker {
	mb := new(EventBroker)
	mb.inGressEvents = make(chan Event)
	mb.running = true
	return mb
}

func init() {
	MessageBroker = NewMessageBroker()
}
