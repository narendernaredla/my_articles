package db

import "errors"

type Article struct {
	Id      string
	Title   string
	Content string
	Author  string
}

func (article *Article) Create() (string, error) {
	return "123456", nil
}

func (article *Article) GetById() (Article, error) {
	return Article{Id: "123456"}, errors.New("Failed")
}

func (article *Article) GetAll() ([]Article, error) {
	var allArticles []Article
	return allArticles, errors.New("Failed")
}
