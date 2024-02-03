package routes

import (
	"encoding/xml"
	"fmt"
	"net/http"
	"time"

	"github.com/sapphirenw/jakeruns.com/src/api"
	"github.com/sapphirenw/jakeruns.com/src/logger"
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
	urls := []URL{
		{Loc: "http://jakeruns.com", LastMod: time.Now(), ChangeFreq: "daily", Priority: 1.0},
		{Loc: "http://jakeruns.com/why", LastMod: time.Now(), ChangeFreq: "daily", Priority: 1.0},
		{Loc: "http://jakeruns.com/schedule", LastMod: time.Now(), ChangeFreq: "daily", Priority: 1.0},
	}

	// fetch the articles and render a dynamic site map from them
	articles, err := api.GetAllArticlesLight()
	if err != nil {
		logger.Critical.Printf("There was an issue rendering the sitemap.xml: %s\n", err)
	} else {
		for _, item := range *articles {
			urls = append(urls, URL{
				Loc:        fmt.Sprintf("https://jakeruns.com/articles/%d", *item.ArticleId),
				LastMod:    time.Now(),
				ChangeFreq: "daily",
				Priority:   1.0,
			})
		}
	}

	set := URLSet{
		Xmlns: "http://www.sitemaps.org/schemas/sitemap/0.9",
		URLs:  urls,
	}

	w.Header().Set("Content-Type", "text/xml")
	enc := xml.NewEncoder(w)
	enc.Indent("", "  ")
	if err := enc.Encode(set); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
