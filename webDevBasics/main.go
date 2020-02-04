package main

import (
	"fmt"
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `<h1>HEY THERE</h1>
		<p>Go is fast</p>
	`)

	fmt.Fprintf(w, "<h1>HEY THERE</h1>")
	fmt.Fprintf(w, "<p>Go is fast</p>")
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.ListenAndServe(":3000", nil)
}
