package main

import (
	"fmt"
	"groupie-tracker/groupie"
	"net/http"
)

func main() {
	groupie.UnmarshalData()
	http.HandleFunc("/", groupie.Handler)
	fmt.Println("Listen And Serve Port 8080")
	http.ListenAndServe(":8080", nil)
}
