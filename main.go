package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/",rootHandler)
	router.GET("/hello",helloHandler)

	router.Run()
}

func rootHandler(c *gin.Context){
	c.JSON(http.StatusOK, gin.H{
		"name":"Andri Maadsa",
		"bio": "seorang ustads",
	})
}

func helloHandler(c *gin.Context){
	c.JSON(http.StatusOK, gin.H{
		"title":"Hello World",
		"subtitle": "Belajar Golang aja",
	})
}