package routes

import (
	"net/http"

	"github.com/sapphirenw/jakeruns.com/src/logger"
	"github.com/sapphirenw/jakeruns.com/src/xtempl"
)

func Why(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	// render the template
	if err := xtempl.XT.ExecuteTemplate(w, "why.html", nil); err != nil {
		logger.Critical.Printf("There was an issue rendering the index page: %s", err)
		Error(w, r)
	}
}
