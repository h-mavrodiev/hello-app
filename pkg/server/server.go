package server

import (
	"fmt"
	"net/http"
	"os"
	"sync"

	"github.com/gin-gonic/gin"
)

// SafeCounter is safe to use concurrently.
type SafeBreak struct {
	mu        sync.Mutex
	breakBool bool
}

// getEnv reads and returns environment variables
func GetEnv(envVar, fallback string) string {
	if value, present := os.LookupEnv(envVar); present {
		return value
	}
	return fallback

}

// HelloAppRouter sets router
func HelloAppRouter(m *sync.Mutex) *gin.Engine {

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
		m.Lock()
		breakBool = true
		m.Unlock()
		c.JSON(http.StatusOK, gin.H{
			"message": "Break request was successful!"})
	})

	return router
}
