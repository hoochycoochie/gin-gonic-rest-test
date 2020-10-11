package main

import "errors"

type article struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

var articleList = []article{
	article{ID: 1, Title: "Article 1", Content: "Article 1 body"},
	article{ID: 2, Title: "Article 2", Content: "Article 2 body"},
}

func GetAllArticles() []article {
	return articleList
}

func getArticleByID(id int) (*article, error) {
	for _, v := range articleList {
		if v.ID == id {
			return &v, nil
		}
	}

	return nil, errors.New("recod not found")
}
