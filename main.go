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
	if port == "" {
		port = "80"
	}
	log.Println("port is ", port)
	err := http.ListenAndServe(fmt.Sprintf(":%v", 80), nil)
	if err != nil {
		log.Fatalln(err)
	}
}
