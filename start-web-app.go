package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

// getEnv reads and returns environment variables
func getEnv(envVar, fallback string) string {
	if value, present := os.LookupEnv(envVar); present {
		return value
	}
	return fallback

}

// helloAppRouter sets router
func helloAppRouter() *gin.Engine {

	var breakBool bool //defaults to false
	router := gin.Default()

	router.GET("/hello", func(c *gin.Context) {
		userInput := c.DefaultQuery("name", "World")
		message := fmt.Sprintf("Hello %s!", userInput)
		c.JSON(http.StatusOK, gin.H{
			"message": message})
	})

	router.GET("/healthz", func(c *gin.Context) {
		if breakBool {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "HTTP status 500"})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"message": "HTTP status 200"})
		}
	})

	router.POST("/break", func(c *gin.Context) {
		breakBool = true
		c.JSON(http.StatusOK, gin.H{
			"message": "Break request was successful!"})
	})

	return router
}

func main() {
	webAppPort := getEnv("WEBAPP_PORT", "8080")
	r := helloAppRouter()
	r.Run(":" + webAppPort)
}
