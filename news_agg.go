package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

type SitemapIndex struct {
	Locations []Location `xml:"sitemap"`
}

type Location struct {
	Loc string `xml:"loc"`
}

/*
[5 5]type == array
[5 5]int is an array
[]type == slice

*/
func (l Location) String() string {
	return fmt.Sprintf(l.Loc)
}

func main() {
	resp, _ := http.Get("https://www.washingtonpost.com/sitemaps/index.xml")
	bytes, _ := ioutil.ReadAll(resp.Body)
	//fmt.Println(bytes)
	resp.Body.Close()

	var s SitemapIndex
	xml.Unmarshal(bytes, &s)
	//fmt.Println(s.Locations)

	//Range returns index and value
	for _, Location := range s.Locations {
		fmt.Printf("%s", Location)
	}
}
