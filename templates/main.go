package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

type NewsAggPage struct {
	Title string
	News  string
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Whoa, Go is neat!</h1>")
}

func newsAggHandler(w http.ResponseWriter, r *http.Request) {
	// Find the Working Directory
	wd, wdErr := os.Getwd()
	if wdErr != nil {
		log.Fatal(wdErr)
	}

	fmt.Println("WD: ", wd)

	p := NewsAggPage{Title: "Amazing New Aggregator", News: "Some News"}
	t, err := template.ParseFiles(wd + "/templates/basictemplating.html")
	if err != nil {
		fmt.Println("ERR: ", err)
	}

	fmt.Println(t.Execute(w, p))
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/foo/", newsAggHandler)
	http.ListenAndServe(":8000", nil)
}
