package routes

import (
	"net/http"

	"github.com/sapphirenw/jakeruns.com/src/api"
	"github.com/sapphirenw/jakeruns.com/src/logger"
	"github.com/sapphirenw/jakeruns.com/src/xtempl"
)

func Index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	// get all articles
	articles, err := api.GetAllArticlesLight()
	if err != nil {
		logger.Critical.Printf("There was an issue getting the index page: %s", err)
		Error(w, r)
		return
	}

	// render the template
	body := map[string]any{"Articles": articles}
	if err := xtempl.XT.ExecuteTemplate(w, "index.html", &body); err != nil {
		logger.Critical.Printf("There was an issue rendering the index page: %s", err)
		Error(w, r)
	}
}
