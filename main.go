package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	log.Println("Time trigger function starting.")

	executeTimmeFunction()

	http.HandleFunc("/", timeTriggerHandler)
	port := os.Getenv("FUNCTIONS_CUSTOMHANDLER_PORT")
	log.Printf("Listening on port %s...", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}

func timeTriggerHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Timer trigger function executed at:", time.Now())
	executeTimmeFunction()
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Time trigger function executed successfully!"))
}

func executeTimmeFunction() {
	log.Println("Hello, this is your scheduled function running!")
}
