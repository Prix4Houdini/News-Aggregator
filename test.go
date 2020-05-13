package main

import (
	"fmt"
	"net/http"
)

/*
type SitemapIndex struct {
	Locations []string `xml:"sitemap>loc"`
}
*/
/*
type News struct {
	Titles []string `xml:"url>loc>title"`
}
*/
/*
[5 5]type == array
[5 5]int is an array
[]type == slice

*/

func main() {
	resp, _ := http.Get("https://www.washingtonpost.com/sitemaps/index.xml")
	//bytes, _ := ioutil.ReadAll(resp.Body)
	//fmt.Println(bytes)
	resp.Body.Close()
	fmt.Println(resp)
	fmt.Println("\n")
	fmt.Println(resp.Body)
	fmt.Println("\n")
	//fmt.Println(bytes)
	//var s SitemapIndex
	//xml.Unmarshal(bytes, &s)
	//fmt.Println(s.Locations)

	//Range returns index and value
	//for _, Location := range s.Locations {
	//fmt.Printf("%s", Location)
	//}
}
