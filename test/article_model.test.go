package test

import (
	"testing"

	"github.com/lucasBiazon/microservices/models"
)

func TestGetAllArticlesModels(t *testing.T) {
	alist := models.GetAllArticles()

	articleList := models.ArticleList
	if len(alist) != len(articleList) {
		t.Fail()
	}

	for i, v := range alist {
		if v.Content != articleList[i].Content ||
			v.ID != articleList[i].ID ||
			v.Title != articleList[i].Title {

			t.Fail()
			break
		}
	}
}
