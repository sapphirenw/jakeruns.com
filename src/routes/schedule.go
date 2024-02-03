package routes

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/sapphirenw/jakeruns.com/src/api"
	"github.com/sapphirenw/jakeruns.com/src/logger"
	"github.com/sapphirenw/jakeruns.com/src/xtempl"
)

func Schedule(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	file, err := os.Open("./public/static/schedule.json")
	if err != nil {
		logger.Critical.Printf("There was an issue getting the schedule json: %s", err)
		Error(w, r)
		return
	}
	defer file.Close()

	b, _ := io.ReadAll(file)

	// Declare map for JSON data
	var schedule map[string]any

	// Unmarshal JSON to map
	json.Unmarshal([]byte(b), &schedule)

	// check if an article exists for each day in the schedule with cursed unsafe code
	articles, err := api.GetAllArticlesLight()
	if err != nil {
		logger.Critical.Printf("There was an issue getting the articles: %s", err)
		Error(w, r)
		return
	}

	// this is the most disgusting code I have ever written, and it would be
	// easily solved by creating a couple structs but I am just too lazy its 1am
	// and this works so voila
	currentIndex := 0
	articleDates := map[string]any{}
	for _, article := range *articles {
		articleDates[article.Created[0:10]] = article.ArticleId
	}
	for idx, week := range schedule["weeks"].([]any) {
		for _, day := range week.(map[string]any)["days"].([]any) {
			// parse the date into the wanted format
			parsedDate, _ := time.Parse("1/2/06", day.(map[string]any)["date"].(string))
			if parsedDate.Compare(time.Now()) == -1 {
				currentIndex = idx
			}
			sqlDate := parsedDate.Format("2006-01-02")
			if value, exists := articleDates[sqlDate]; exists {
				day.(map[string]any)["articleId"] = value
			}
		}
	}
	schedule["currentIndex"] = currentIndex

	// render the template
	if err := xtempl.XT.ExecuteTemplate(w, "schedule.html", schedule); err != nil {
		logger.Critical.Printf("There was an issue rendering the index page: %s", err)
		Error(w, r)
	}
}
