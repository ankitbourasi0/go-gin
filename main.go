package main

import (
	"fmt"
	"gin-tutorial/controllers"
	"gin-tutorial/internal/database"
	"gin-tutorial/services"

	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// Create a Gin router with default middleware (logger and recovery)
	router := gin.Default()

	db := internal.InitDatabase()

	if db == nil {
		fmt.Println("Database initialization failed!!!")
	}

	fmt.Println("Database Initialization Successful", db)

	//Access notes service
	notesService := &services.NotesService{}
	//Calling DB Initialize METHOD
	notesService.InitService(db)

	//router.GET("/ping", func(c *gin.Context) {
	//	// Return JSON response
	//	c.JSON(http.StatusOK, gin.H{
	//		"message": "pong",
	//	})
	//})
	//
	//router.GET("/me/:id", func(c *gin.Context) {
	//	var id = c.Param("id")
	//
	//	c.JSON(http.StatusOK, gin.H{
	//		"user_id": id,
	//	})
	//})
	//
	//router.GET("/me/:id/:userId", func(c *gin.Context) {
	//	var id = c.Param("id")
	//	var userId = c.Param("userId")
	//
	//	c.JSON(http.StatusOK, gin.H{
	//		"id":      id,
	//		"user_id": userId,
	//	})
	//})
	//router.POST("/me", func(c *gin.Context) {
	//	//1. What you want?  - EMAIL, PASSWORD
	//
	//	//2. define the  type with
	//	type MeRequest struct { //json request format of client
	//		Email    string `json:"email"`
	//		Password string `json:"password"`
	//	}
	//	//3. create a var of your request type
	//	var meRequest MeRequest
	//	//4. bind request with variable pointer
	//	c.BindJSON(&meRequest)
	//
	//	//5.return the response,
	//	c.JSON(http.StatusOK, gin.H{
	//		//response
	//		"email":    meRequest.Email,
	//		"password": meRequest.Password,
	//	})
	//})
	//
	//router.POST("/me/id", func(c *gin.Context) {
	//
	//	type MeRequest struct { // 1. email is required
	//		Email    string `json:"email" binding:"required"`
	//		Password string `json:"password"`
	//	}
	//	var meRequest MeRequest
	//	//validate
	//	if err := c.BindJSON(&meRequest); err != nil {
	//		c.JSON(http.StatusBadRequest, gin.H{
	//			"error": err.Error(),
	//		})
	//		return
	//	}
	//
	//	c.JSON(http.StatusOK, gin.H{
	//		"email":    meRequest.Email,
	//		"password": meRequest.Password,
	//	})
	//})
	//
	////Replace entire object
	//router.PUT("/me", func(c *gin.Context) {
	//
	//	type MeRequest struct { // 1. email is required
	//		Email    string `json:"email" binding:"required"`
	//		Password string `json:"password"`
	//	}
	//	var meRequest MeRequest
	//	//validate
	//	if err := c.BindJSON(&meRequest); err != nil {
	//		c.JSON(http.StatusBadRequest, gin.H{
	//			"error": err.Error(),
	//		})
	//		return
	//	}
	//
	//	c.JSON(http.StatusOK, gin.H{
	//		"email":    meRequest.Email,
	//		"password": meRequest.Password,
	//	})
	//})
	//
	////Replace Sub Part not complete Object
	//router.PATCH("/me", func(c *gin.Context) {
	//
	//	type MeRequest struct { // 1. email is required
	//		Email    string `json:"email" binding:"required"`
	//		Password string `json:"password"`
	//	}
	//	var meRequest MeRequest
	//	//validate
	//	if err := c.BindJSON(&meRequest); err != nil {
	//		c.JSON(http.StatusBadRequest, gin.H{
	//			"error": err.Error(),
	//		})
	//		return
	//	}
	//
	//	c.JSON(http.StatusOK, gin.H{
	//		"email":    meRequest.Email,
	//		"password": meRequest.Password,
	//	})
	//})
	//
	//router.DELETE("/me/:key", func(c *gin.Context) {
	//
	//	var key = c.Param("key")
	//	fmt.Println(key)
	//
	//	c.JSON(http.StatusOK, gin.H{
	//		"message": "Deleted Successfully",
	//	})
	//})

	//1.You can assume you are accessing a Class(NotesController) via package(controllers)
	notesController := &controllers.NotesController{}
	//2.And Access each method of the class as shown below!
	notesController.NewNotesController(router, *notesService)

	if err := router.Run(); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
