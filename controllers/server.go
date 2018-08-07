package controllers

import (
	"fmt"
	"github.com/gorilla/websocket"
)

var queue []*websocket.Conn
var mess = make(chan string, 10)

func init() {
	go func() {
		for {
			str, ok := <-mess
			if ok {
				go func(t string) {
					for k, v := range queue {
						if v != nil {
							err := v.WriteMessage(1, []byte(t))
							if err != nil {
								fmt.Println("Current Client has Closed ", k)
								v.Close()
								queue[k] = nil
							}
						}
					}
				}(str)
			} else {

			}
		}
	}()
}

func Add(conn *websocket.Conn) {
	fmt.Println("Add a Client")
	sig := false
	for k, v := range queue {
		if v == nil {
			fmt.Println(k)
			queue[k] = conn
			sig = true
		}
	}
	if !sig {
		queue = append(queue, conn)
	}
	go Handle(conn)
}

func Handle(conn *websocket.Conn) {
	for {
		_, p, err := conn.ReadMessage()
		if err == nil {
			mess <- string(p)
		} else {
			conn.Close()
			break
		}
	}
}
