package main

import (
  "code.google.com/p/go.net/websocket"
  "canary"
  "config"
  "encoding/json"
  "fmt"
)

func main() {
  ws, _ := websocket.Dial(config.WS_URL, "", config.HOST_URL)
  auth_set := canary.AuthenticationSet {
    DeviceId: config.DEVICE_ID,
    DevicePassword: config.DEVICE_PASSWORD,
  }
  json, _ := json.Marshal(auth_set)
  ws.Write([]byte(json))
  for ;; {
    var msg = make([]byte, 2048)
    n, err := ws.Read(msg)
    if err != nil {
      fmt.Println("error reading from web socket!")
    } else if n > 0 {
      fmt.Println(string(msg))
    }
  }
}
