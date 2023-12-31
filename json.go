package main

import (
  "net/http"
  "encoding/json"
  "log"
)

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
  data, err := json.Marshal(payload)
  if err != nil {
    log.Println("Failed to marshal JSON response: %v", payload)
    w.WriteHeader(500)
    return
  }

  w.Header().Add("Content-Type", "application/json")
  w.WriteHeader(200)
  w.Write(data)
}
