package main

import (
	"fmt"
	"log"
	"net/http"

	api "API/handler"
)

func main() {
	http.HandleFunc("/", api.Homehandler)
	http.HandleFunc("/artist/", api.ArtistHandler)
	log.Println("Starting server on http://127.0.0.1:8070")
	if err := http.ListenAndServe(":8070", nil); err != nil {
		fmt.Println(err.Error())
		return
	}
}
