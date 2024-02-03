package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/sapphirenw/jakeruns.com/src/api"
	"github.com/sapphirenw/jakeruns.com/src/api/response"
	"github.com/sapphirenw/jakeruns.com/src/logger"
	"github.com/sapphirenw/jakeruns.com/src/routes"
	"github.com/sapphirenw/jakeruns.com/src/xtempl"
)

func main() {
	xtempl.XT = xtempl.New()
	if err := xtempl.XT.ParseDir("templates/", []string{".html"}); err != nil {
		fmt.Printf("There was an issue parsing the templates: %s\n", err)
		os.Exit(1)
	}

	r := chi.NewRouter()

	// middleware
	r.Use(logger.LoggerMiddleware)
	r.Use(middleware.Compress(5, "text/html", "text/css", "text/javascript"))
	r.Use(middleware.RedirectSlashes)

	// handle static files
	fsPublic := http.FileServer(http.Dir("./public"))
	r.Handle("/public/*", http.StripPrefix("/public", fsPublic))

	// static assets
	r.Get("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		file, err := os.ReadFile("./public/images/favicon.ico")
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(404)
			w.Write([]byte("Could not find"))
			return
		}
		w.Header().Set("Content-Type", "image/x-icon")
		w.Write(file)
	})

	// sitemap
	r.Get("/sitemap.xml", routes.Sitemap)

	// routes
	r.NotFound(routes.NotFound)
	r.Get("/error", routes.Error)
	r.Get("/", routes.Index)
	r.Get("/why", routes.Why)
	r.Get("/schedule", routes.Schedule)

	// articles handler
	r.Get("/articles/{articleId}", routes.Article)

	// api router
	sr := chi.NewRouter()
	sr.Use(api.AuthMiddleware)

	// api routes
	sr.Get("/test", func(w http.ResponseWriter, r *http.Request) {
		response.WriteStr(w, 200, "Configuration successful")
	})
	sr.Get("/articles", api.RouteGetAllArticles)
	sr.Post("/articles", api.RouteCreateArticle)

	// mount on main router
	r.Mount("/api", sr)

	fmt.Println(http.ListenAndServe(":3000", r))
}
