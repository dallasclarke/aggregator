package main

import (
	"fmt"
	"log"
	"os"
  "net/http"

  "github.com/joho/godotenv"
  "github.com/go-chi/chi"
  "github.com/go-chi/cors"
)

func main() {
  godotenv.Load(".env")

  portString := os.Getenv("PORT")
  if portString == "" {
    log.Fatal("PORT is not found in .env file")
  }

  router := chi.NewRouter()

  router.Use(cors.Handler(cors.Options{
    AllowedOrigins: []string{"https://*", "http://*"},
    AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
    AllowedHeaders: []string{"*"},
    ExposedHeaders: []string{"Link"},
    MaxAge: 300,
  }))

  srv := &http.Server{
    Handler: router,
    Addr: ":" + portString,
  }

  log.Printf("Server starting on port %v", portString)
  err := srv.ListenAndServe()
  if err != nil {
    log.Fatal(err)
  }

  fmt.Println("PORT: ", portString)
}
