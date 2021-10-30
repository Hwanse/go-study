package main

import (
	"github.com/gorilla/pat"
	"github.com/gorilla/websocket"
	"github.com/urfave/negroni"
	"log"
	"net/http"
)

type User struct {
	Name string `json:"name"`
	Email string `json:"email"`
	age int `json:"age"`
}

type Message struct {
	Type string `json:"type"`
	Data interface{} `json:"data"`
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}


func wsHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	for {
		m := &Message{}
		err := conn.ReadJSON(m)
		if err != nil {
			log.Println(err)
			return
		}

		err = conn.WriteJSON(m)
		if err != nil {
			log.Println(err)
			return
		}

	}
}

func main() {

	router := pat.New()
	router.Get("/ws", wsHandler)

	negroni := negroni.Classic()
	negroni.UseHandler(router)

	http.ListenAndServe(":3000", router)

}
