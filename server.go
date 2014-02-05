package main

import (
  "code.google.com/p/go.net/websocket"
//  "config"
  "canary"
  "fmt"
  "net/http"
)

var MessageChannel = make(chan canary.Message)
var LogChannel = make(chan string)


func MessageServer(ws *websocket.Conn) {
  var msg = make([]byte, 2048)
  for ;; {
    fmt.Println("New websocket connected")
    n := 0
    n, err := ws.Read(msg)
    if err != nil {
      fmt.Println("Error reading from websocket, this means the client probably disconnected.  Killing thread!")
      break;
    } else if n > 0 {
      clientId := string(msg)
      message := <- MessageChannel
      if message.DestinationId == clientId {
        //ws.Write(byte(message))
      }
    }
  }
}

func ApiServer(w http.ResponseWriter, r *http.Request) {
  fmt.Println("New API connection!")
  // MessageChannel <- r.Body
}

func main() {
  go func() {
    LogChannel <- "Booting websocket server..."
    // fmt.Println("Booting websocket server...")
    http.Handle("/ws", websocket.Handler(MessageServer))
    err := http.ListenAndServe(":8080", nil)
    if err != nil {
      LogChannel <- err.Error()
      panic(err.Error())
    }
  }()
  go func() {
    LogChannel <- "Booting API server..."
    // fmt.Println("Booting API server...")
    http.HandleFunc("/api", ApiServer)
    err := http.ListenAndServe(":8082", nil)
    if err != nil {
      LogChannel <- err.Error()
      panic(err.Error())
    }
  }()
  for ;; {
    log_message := <- LogChannel
    fmt.Println(log_message)
  }
}
