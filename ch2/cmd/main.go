package main

import (
	"github.com/Vadakan/holmes89/handlers"
	"log"
	"net/http"


)

func main() {

	addr := ":8080"

	mux := http.NewServeMux()

	mux.HandleFunc("/hello", handlers.TranslateHandler)

	log.Printf("listening on %s\n", addr)

	log.Fatal(http.ListenAndServe(addr, mux))
}
