package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

// In computer science, marshalling or marshaling is the process of transforming the memory representation of an object to a data format suitable for storage or transmission, and it is typically used when data must be moved between different parts of a computer program or from one program to another.

type SitemapIndex struct {
	Locations []Location `xml:"sitemap"` // `` IS marshling and capital Locations means it will be exported
}

// [5]type == array => this is going to have a fixed size
// []type == slice

type Location struct {
	Loc string `xml:"loc"`
}

func (l Location) String() string { // NOTE: If you capitalize a method/receiver on a struct, it automatically gets run when user runs the struct itself
	return fmt.Sprintf(l.Loc) // converts an interface in to a string
}

func main() {
	resp, _ := http.Get("https://www.washingtonpost.com/news-sitemaps/index.xml") // _ in GO means to define any var that you don't intend to use because unused vars will through an error
	bytes, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	var s SitemapIndex
	xml.Unmarshal(bytes, &s)
	// fmt.Print(s.Locations)

	// fmt.Printf("Here %s some %s", "are", "variables")

	for _, Location := range s.Locations {
		fmt.Printf("\n%s", Location)
	}
}
