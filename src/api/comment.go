package api

import (
	"fmt"
)

type Comment struct {
	ArticleId int    `db:"articleId" json:"articleId"`
	CommentId int    `db:"commentId" json:"commentId"`
	Name      string `db:"name" json:"name"`
	Email     string `db:"email" json:"email"`
	Content   string `db:"content" json:"content"`
	Created   string `db:"created" json:"created"`
	Updated   string `db:"updated" json:"updated"`
}

func (comment *Comment) InsertString() string {
	return fmt.Sprintf("INSERT INTO Comment (articleId, name, email, content) VALUES (:articleId, :name, :email, :content)")
}

func (comment *Comment) UpdateString() string {
	return fmt.Sprintf("UPDATE Comment SET name = :name, email = :email, content = :content WHERE commentId = :commentId")
}
