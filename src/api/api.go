package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/sapphirenw/jakeruns.com/src/api/response"
	"github.com/sapphirenw/jakeruns.com/src/logger"
)

var db *sqlx.DB

func getDB() (*sqlx.DB, error) {
	if db != nil {
		return db, nil
	}

	host := os.Getenv("MYSQL_HOST")
	user := os.Getenv("MYSQL_USER")
	pass := os.Getenv("MYSQL_PASS")

	if user == "" || pass == "" || host == "" {
		return nil, fmt.Errorf("missing required environment variables")
	}

	connStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, pass, host, "3306", "jakeruns")

	maxRetries := 5
	for i := 0; i < maxRetries; i++ {
		db, err := sqlx.Connect("mysql", connStr)
		if err == nil {
			return db, nil
		}
		backoff := time.Duration(1<<i) * time.Second
		log.Printf("Failed to connect to database: %v, retrying in %v", err, backoff)
		time.Sleep(backoff)
	}

	if db == nil {
		return nil, fmt.Errorf("database connection not established")
	}

	return db, nil
}

func RouteCreateArticle(w http.ResponseWriter, r *http.Request) {
	// de-serialize the json
	var article Article
	err := json.NewDecoder(r.Body).Decode(&article)
	if err != nil {
		response.WriteStr(w, 400, "There was an issue de-serializing the body: %s", err)
		return
	}

	// insert or update the article into the database based on if a valid articleId was passed
	if article.ArticleId == nil {
		logger.Debug.Println("Inserting a new article into the database")
		err = article.Insert()
	} else {
		logger.Debug.Printf("Updating article with articleId: %d\n", *article.ArticleId)
		err = article.Update()
	}
	if err != nil {
		response.WriteStr(w, 500, "There was an issue creating the article: %s", err)
		return
	}

	logger.Debug.Printf("Successful action on article: %d\n", article.ArticleId)

	// return the request
	response.WriteObj(w, 200, article)
}

func RouteGetAllArticles(w http.ResponseWriter, r *http.Request) {
	articles, err := GetAllArticles()
	if err != nil {
		response.WriteStr(w, 500, "There was an issue getting the articles: %s", err)
		return
	}

	response.WriteObj(w, 200, articles)
}

func RouteGetArticle(w http.ResponseWriter, r *http.Request) {
	// read the params
	articleId := chi.URLParam(r, "articleId")
	if articleId == "" {
		response.WriteStr(w, 400, "Invalid articleId: %s", articleId)
		return
	}

	// convert to an int to ensure type checking
	i, err := strconv.Atoi(articleId)
	if err != nil {
		response.WriteStr(w, 400, "Invalid articleId: %s", err)
		return
	}

	article, err := GetArticle(i)
	if err != nil {
		response.WriteStr(w, 500, "There was an issue getting the article: %s", err)
	}

	response.WriteObj(w, 200, article)
}
