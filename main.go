package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// Create a Gin router with default middleware (logger and recovery)
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		// Return JSON response
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	router.GET("/me/:id", func(c *gin.Context) {
		var id = c.Param("id")

		c.JSON(http.StatusOK, gin.H{
			"user_id": id,
		})
	})

	router.GET("/me/:id/:userId", func(c *gin.Context) {
		var id = c.Param("id")
		var userId = c.Param("userId")

		c.JSON(http.StatusOK, gin.H{
			"id":      id,
			"user_id": userId,
		})
	})

	router.POST("/me", func(c *gin.Context) {
		//1. What you want?  - EMAIL, PASSWORD

		//2. define the  type with
		type MeRequest struct { //json request format of client
			Email    string `json:"email"`
			Password string `json:"password"`
		}
		//3. create a var of your request type
		var meRequest MeRequest
		//4. bind request with variable pointer
		c.BindJSON(&meRequest)

		//5.return the response,
		c.JSON(http.StatusOK, gin.H{
			//response
			"email":    meRequest.Email,
			"password": meRequest.Password,
		})
	})

	router.POST("/me/id", func(c *gin.Context) {

		type MeRequest struct { // 1. email is required
			Email    string `json:"email" binding:"required"`
			Password string `json:"password"`
		}
		var meRequest MeRequest
		//validate
		if err := c.BindJSON(&meRequest); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"email":    meRequest.Email,
			"password": meRequest.Password,
		})
	})

	if err := router.Run(); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
