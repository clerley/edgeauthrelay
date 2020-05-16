package controller

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

import (
	"com/novare/auth/model"
	"com/novare/auth/sse"
	"encoding/json"
	"log"
	"net/http"

	"novacity.org/controller"
)

type eventsReq struct {
	EventIDs []string `json:"events"`
}

//ServerSentEvents ... This will always just poll for events.
func ServerSentEvents(w http.ResponseWriter, r *http.Request) {
	//This will keep the connection alive.
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	var events eventsReq
	usr := r.Context().Value(controller.CtxUser).(model.User)

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&events)
	if err != nil {
		log.Printf("No events requested. The event IDs must be known")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	subscriber := sse.NewSubscriber(usr.ID.Hex())
	for i := range events.EventIDs {
		subscriber.AddEvent(events.EventIDs[i])
	}

	broker := sse.MessageBroker

	//Make sure there isn't a listening user with the same ID
	broker.Unsubscribe(subscriber)

	//Subscribe to the queue
	broker.Subscribe(subscriber)

	// We need to be able to flush for SSE
	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "Flushing not supported", http.StatusNotImplemented)
		return
	}

	// Returns a channel that blocks until the connection is closed
	close := r.Context().Done()

	for {
		select {
		case <-close:
			log.Printf("The connection was closed, unsubscribing now")
			// Disconnect the client when the connection is closed
			broker.Unsubscribe(subscriber)
			return

		case event := <-subscriber.EventChannel:
			buf, err := json.Marshal(event)
			if err != nil {
				continue
			}
			w.Write(buf)
			flusher.Flush()
		}
	}

}
