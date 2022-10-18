package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "hello world")
	})

	port := os.Getenv("PORT")
	log.Printf("get port from env => (%v)\n", port)
	if port == "" {
		port = "8080"
	}
	log.Println("now listening on port:", port)
	err := http.ListenAndServe(fmt.Sprintf("0.0.0.0:%v", port), nil)
	if err != nil {
		log.Fatalln(err)
	}
}
