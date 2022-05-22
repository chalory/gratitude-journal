package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
)

type entry struct {
	ID        int    `json:"id"`
	Entry     string `json:"entry"`
	Date      string `json:"date"`
	User_name string `json:"user_name"`
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	// Set the router as the default one shipped with Gin
	router := gin.Default()

	// Serve frontend static files
	router.Use(static.Serve("/", static.LocalFile("./views/js", true)))

	// Setup route group for the API
	{
		router.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		})
	}

	router.GET("/getEntries", GetEntry)
	router.Use(cors.Default())

	// JokeHandler retrieves a list of available jokes

	// Start and run the server
	router.Run(":8080")
}
