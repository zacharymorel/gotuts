package main

import (
	"encoding/xml"
	"io/ioutil"
	"net/http"
)

// In computer science, marshalling or marshaling is the process of transforming the memory representation of an object to a data format suitable for storage or transmission, and it is typically used when data must be moved between different parts of a computer program or from one program to another.

type SitemapIndex struct {
	Locations []string `xml:"sitemap>loc"` // `` IS marshling and capital Locations means it will be exported
}

type News struct {
	Titles    []string `xml:"url>news>title"`
	Keywords  []string `xml:"url>news>keywords"`
	Locations []string `xml:"url>location"`
}

// [5]type == array => this is going to have a fixed size
// []type == slice

func main() {
	var s SitemapIndex
	var n News

	resp, _ := http.Get("https://www.washingtonpost.com/news-sitemaps/index.xml") // _ in GO means to define any var that you don't intend to use because unused vars will through an error
	bytes, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	xml.Unmarshal(bytes, &s)
	// fmt.Print(s.Locations)

	for _, Location := range s.Locations {
		// fmt.Printf("\n%s", Location)
		resp, _ := http.Get(Location) // _ in GO means to define any var that you don't intend to use because unused vars will through an error
		bytes, _ := ioutil.ReadAll(resp.Body)
		// resp.Body.Close()
		xml.Unmarshal(bytes, &n)
	}
}
