package api

import "fmt"

type Article struct {
	ArticleId   *int    `db:"articleId" json:"articleId"`
	Title       string  `db:"title" json:"title"`
	Description string  `db:"description" json:"description"`
	Content     string  `db:"content" json:"content"`
	Author      string  `db:"author" json:"author"`
	Tags        *string `db:"tags" json:"tags"`
	Created     string  `db:"created" json:"created"`
	Updated     string  `db:"updated" json:"updated"`
}

func (article *Article) InsertString() string {
	return fmt.Sprintf("INSERT INTO Article (title, description, content, author, tags) VALUES (:title, :description, :content, :author, :tags)")
}

func (article *Article) UpdateString() string {
	return fmt.Sprintf("UPDATE Article SET title = :title, description = :description, content = :content, author = :author, tags = :tags WHERE articleId = :articleId")
}

/*
Inserts an article into the database. The article object does not need to have the
`articleId`, `created`, or `updated` field set. As a result of this function, the article
object will be updated internally to include these fields after a successful insert.
*/
func (article *Article) Insert() error {
	db, err := getDB()
	if err != nil {
		return fmt.Errorf("there was an issue getting the database: %s", err)
	}

	result, err := db.NamedExec(article.InsertString(), article)
	if err != nil {
		return fmt.Errorf("there was an issue execing the query: %s", err)
	}

	// get the id
	id, err := result.LastInsertId()
	if err != nil {
		return fmt.Errorf("there was an issue getting the new row id: %s", err)
	}

	// query for the full record
	err = db.Get(article, fmt.Sprintf("SELECT * FROM Article WHERE articleId=%d", id))
	if err != nil {
		return fmt.Errorf("there was an issue selecting the new row: %s", err)
	}

	return nil
}

func (article *Article) Update() error {
	db, err := getDB()
	if err != nil {
		return fmt.Errorf("there was an issue getting the database: %s", err)
	}

	_, err = db.NamedExec(article.UpdateString(), article)
	if err != nil {
		return fmt.Errorf("there was an issue execing the query: %s", err)
	}

	// query for the full record
	err = db.Get(article, fmt.Sprintf("SELECT * FROM Article WHERE articleId=%d", *article.ArticleId))
	if err != nil {
		return fmt.Errorf("there was an issue selecting the new row: %s", err)
	}

	return nil
}

/*
Gets all content from the articles table, sorted by created DESC
*/
func GetAllArticles() (*[]Article, error) {
	db, err := getDB()
	if err != nil {
		return nil, fmt.Errorf("there was an issue getting the database: %s", err)
	}
	articles := []Article{}
	err = db.Select(&articles, "SELECT * FROM Article ORDER BY created DESC")
	if err != nil {
		return nil, fmt.Errorf("there was an issue querying the database: %s", err)
	}
	return &articles, nil
}

/*
Gets all of the articles from the database without the `content` field, which should
be a much lighter payload, sorted by created DESC
*/
func GetAllArticlesLight() (*[]Article, error) {
	db, err := getDB()
	if err != nil {
		return nil, fmt.Errorf("there was an issue getting the database: %s", err)
	}
	articles := []Article{}
	err = db.Select(&articles, "SELECT articleId, title, description, author, tags, created, updated FROM Article ORDER BY created DESC")
	if err != nil {
		return nil, fmt.Errorf("there was an issue querying the database: %s", err)
	}
	return &articles, nil
}

/*
Gets a single article with an `articleId`
*/
func GetArticle(articleId int) (*Article, error) {
	db, err := getDB()
	if err != nil {
		return nil, fmt.Errorf("there was an issue getting the database: %s", err)
	}

	article := Article{}
	err = db.Get(&article, fmt.Sprintf("SELECT * FROM Article WHERE articleId = %d", articleId))
	if err != nil {
		return nil, fmt.Errorf("there was an issue querying the database: %s", err)
	}

	return &article, err
}
