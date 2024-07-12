package main

import (
	"github.com/gin-gonic/gin"
	"github.com/lucasBiazon/microservices/handlers"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")

	router.GET("/", handlers.ShowIndexPage)
	router.GET("/article/view/:article_id", handlers.GetArticle)

	router.Run()
}
