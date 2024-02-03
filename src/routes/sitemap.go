package routes

import (
	"encoding/xml"
	"net/http"
	"time"
)

// URLSet is a container for the set of URLs.
type URLSet struct {
	XMLName xml.Name `xml:"urlset"`
	Xmlns   string   `xml:"xmlns,attr"`
	URLs    []URL    `xml:"url"`
}

// URL is the data structure for each URL entry.
type URL struct {
	Loc        string    `xml:"loc"`
	LastMod    time.Time `xml:"lastmod"`
	ChangeFreq string    `xml:"changefreq"`
	Priority   float64   `xml:"priority"`
}

func Sitemap(w http.ResponseWriter, r *http.Request) {
	// urls := []URL{
	// 	{Loc: "http://talk.portlandai.io/", LastMod: time.Now(), ChangeFreq: "daily", Priority: 1.0},
	// 	{Loc: "http://talk.portlandai.io/search", LastMod: time.Now(), ChangeFreq: "daily", Priority: 1.0},
	// }

	// // fetch the articles and render a dynamic site map from them
	// urlStr := fmt.Sprintf("%s/vendors/%s/articles", lib.API_BASE, lib.VENDOR_ID)
	// articles, err := article.GetArticles(urlStr)
	// if err == nil {
	// 	for _, item := range articles {
	// 		urls = append(urls, URL{
	// 			Loc:        fmt.Sprintf("%s/articles/%d", lib.BLOG_BASE, item.Id),
	// 			LastMod:    time.Now(),
	// 			ChangeFreq: "daily",
	// 			Priority:   1.0,
	// 		})
	// 	}
	// }

	// set := URLSet{
	// 	Xmlns: "http://www.sitemaps.org/schemas/sitemap/0.9",
	// 	URLs:  urls,
	// }

	// w.Header().Set("Content-Type", "text/xml")
	// enc := xml.NewEncoder(w)
	// enc.Indent("", "  ")
	// if err := enc.Encode(set); err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// }
}
