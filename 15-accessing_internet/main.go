package main

import (
	"encoding/xml"
	"fmt"
)

var washPostXML = []byte(`
<sitemapindex>
   <sitemap>
      <loc>http://www.washingtonpost.com/news-politics-sitemap.xml</loc>
   </sitemap>
   <sitemap>
      <loc>http://www.washingtonpost.com/news-blogs-politics-sitemap.xml</loc>
   </sitemap>
   <sitemap>
      <loc>http://www.washingtonpost.com/news-opinions-sitemap.xml</loc>
   </sitemap>
</sitemapindex>`)

type SitemapIndex struct {
	Locations []Location `xml:"sitemap"`
}

type Location struct {
	Loc string `xml:"loc"`
}

func (l Location) String() string {
	return fmt.Sprintf(l.Loc)
}

func main() {

	fmt.Println("\nWelcome to go")

	// resp, _ := http.Get("https://www.washingtonpost.com/news-sitemap-index.xml")
	// bytes, _ := ioutil.ReadAll(resp.Body)
	// resp.Body.Close()

	bytes := washPostXML

	var s SitemapIndex
	xml.Unmarshal(bytes, &s)

	// fmt.Println(s.Locations)
	for _, Location := range s.Locations {
		fmt.Printf("\n%s", Location)
	}

}
