package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

func worldHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	log.Println(vars["id"])

	upgrader := websocket.Upgrader {
		ReadBufferSize: 	DEFAULT_READ_BUFFER_SIZE,
		WriteBufferSize: 	DEFAULT_WRITE_BUFFER_SIZE,
	}

	c, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		log.Println(err)
	}
	
	for {
		
		_, buf, err := c.ReadMessage()

		if err != nil {
			
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway) {
				log.Println(err)
			}

		}

		var msg PlayerAction

		err = json.Unmarshal(buf, &msg)

		if err != nil {
			log.Println(err)
		}

		switch(msg.Resource) {
		case RESOURCE_PLAYER:
			playerAction(msg)
		case RESOURCE_KINGDOM:

		default:
			log.Println("Unknown client event resource: ", msg.Resource)
		}
		
	}
	
} // worldHandler
