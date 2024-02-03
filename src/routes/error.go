package routes

import (
	"fmt"
	"net/http"

	"github.com/sapphirenw/jakeruns.com/src/xtempl"
)

func Error(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if err := xtempl.XT.ExecuteTemplate(w, "error.html", nil); err != nil {
		fmt.Println(err)
		http.Error(w, "<html><p>There was a critical error</p></html>", http.StatusInternalServerError)
	}
}
