package main

import (
	"fmt"
	"net/http"
)

func index_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w,
		"<h1>Hey bois</h1>")
	fmt.Fprintf(w,
		"<h2>SUP %s and %s </h2>", "alos", "cinam")

	fmt.Fprintf(w, `
		<span>Walla</span>
		<input>Enter</input>
	
	`)

}

func main() {
	http.HandleFunc("/", index_handler)
	http.ListenAndServe(":8000", nil)

}
