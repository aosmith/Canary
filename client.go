package main

import (
  "code.google.com/p/go.net/websocket"
)

func main() {
  _, _ = websocket.Dial("ws://localhost:8080/ws", "", "http://localhost/")
  for ;; {

  }
}
