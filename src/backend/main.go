package main

import (
	"net/http"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

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
	// Our API will consit of just two routes
	api.GET("/BFS", BFShandler)
	api.GET("IDS", IDShandler)

	// Start and run the server
	router.Run(":3000")
}

// BFShandler Receives
func BFShandler(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, "TES")
}

// LikeJoke increments the likes of a particular joke Item
func IDShandler(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, gin.H{
		"message": "LikeJoke handler not implemented yet",
	})
}
