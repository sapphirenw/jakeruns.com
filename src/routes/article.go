package routes

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/sapphirenw/jakeruns.com/src/api"
	"github.com/sapphirenw/jakeruns.com/src/logger"
	"github.com/sapphirenw/jakeruns.com/src/markdown"
	"github.com/sapphirenw/jakeruns.com/src/xtempl"
)

func Article(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	// parse the articleId
	articleId := chi.URLParam(r, "articleId")
	if articleId == "" {
		logger.Critical.Printf("The articleId was invalid: %s\n", articleId)
		NotFound(w, r)
		return
	}

	// convert to an int to ensure it is correct
	i, err := strconv.Atoi(articleId)
	if err != nil {
		logger.Critical.Printf("The articleId was invalid: %s\n", articleId)
		NotFound(w, r)
		return
	}

	// get the article
	article, err := api.GetArticle(i)
	if err != nil {
		logger.Critical.Printf("There was an issue getting the article: %s\n", err)
		NotFound(w, r)
		return
	}

	// render the markdown into html
	content, err := markdown.RenderMarkdown([]byte(article.Content))
	if err != nil {
		logger.Critical.Printf("There was an issue rendering the markdown: %s\n", err)
		Error(w, r)
		return
	}

	body := map[string]any{"Article": article, "Content": content, "Link": fmt.Sprintf("https://jakeruns.com/articles/%d", i)}
	if err := xtempl.XT.ExecuteTemplate(w, "article.html", &body); err != nil {
		logger.Critical.Printf("There was an issue rendering the article: %s", err)
		Error(w, r)
	}
}
