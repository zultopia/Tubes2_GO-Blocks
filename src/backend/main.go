package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

type Request struct {
	StartTitle  string `json:"startTitle" binding:"required"`
	StartURL    string `json:"startURL" binding:"required"`
	TargetTitle string `json:"targetTitle" binding:"required"`
	TargetURL   string `json:"targetURL" binding:"required"`
}

func main() {
	// Set the router as the default one shipped with Gin
	router := gin.Default()

	// Serve frontend static files
	router.Use(static.Serve("/", static.LocalFile("../wikirace/build", true)))
	router.Use(static.Serve("/home", static.LocalFile("../wikirace/build", true)))
	router.Use(static.Serve("/about", static.LocalFile("../wikirace/build", true)))
	router.Use(static.Serve("/howtouse", static.LocalFile("../wikirace/build", true)))

	// Setup route group for the API
	api := router.Group("/api")
	{
		api.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		})
	}
	// Our API will consist of just two routes
	api.POST("/BFS", BFShandler)
	api.POST("/IDS", IDShandler)

	// Start and run the server
	router.Run(":3000")
}

// BFShandler Receives
func BFShandler(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	request := Request{}
	c.BindJSON(&request)
	startPage := WikiPage{Title: request.StartTitle, URL: request.StartURL}
	targetPage := WikiPage{Title: request.TargetTitle, URL: request.TargetURL}
	startTime := time.Now()
	fmt.Println(request)
	path, articlesVisited := BFSGo(startPage, targetPage, true)
	fmt.Println(path)
	endTime := time.Now()
	elapsedTime := endTime.Sub(startTime)

	c.JSON(http.StatusOK, gin.H{
		"paths":           path,
		"articlesVisited": articlesVisited,
		"executionTime":   elapsedTime.Milliseconds(),
	})
}

// LikeJoke increments the likes of a particular joke Item
func IDShandler(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	request := Request{}
	c.BindJSON(&request)
	startPage := WikiPage{Title: request.StartTitle, URL: request.StartURL}
	targetPage := WikiPage{Title: request.TargetTitle, URL: request.TargetURL}
	startTime := time.Now()
	fmt.Println(request)
	path, articlesVisited := IDS(startPage, targetPage, 10, true)
	fmt.Println(path)
	endTime := time.Now()
	elapsedTime := endTime.Sub(startTime)

	c.JSON(http.StatusOK, gin.H{
		"paths":           path,
		"articlesVisited": articlesVisited,
		"executionTime":   elapsedTime.Milliseconds(),
	})
}
