package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"yoyoserver/routers"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}
	endPoint := fmt.Sprintf(":%s", port)

	routersInit := routers.InitRouter()

	server := &http.Server{
		Addr:		endPoint,
		Handler:	routersInit,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	// Listen and serve on defined port
	log.Printf("Listening on port %s", port)
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}