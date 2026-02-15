package controllers

import (
	"gin-tutorial/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthContrller struct {
	authService services.AuthService
}

func (a *AuthContrller) InitAuthController(authService services.AuthService) *AuthContrller {
	a.authService = authService
	return a
}
func (a *AuthContrller) InitRoutes(router *gin.Engine) {
	auth := router.Group("/auth")
	auth.POST("/login", a.Login())
	auth.POST("/register", a.Register())
	auth.POST("/forget-password")
}

func (a *AuthContrller) Register() gin.HandlerFunc {

	type RegisterBody struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	return func(c *gin.Context) {

		var registerBody RegisterBody
		if err := c.BindJSON(&registerBody); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		user, err := a.authService.Register(&registerBody.Email, &registerBody.Password)
		if err != nil {
			c.JSON(404, gin.H{
				"error": err,
			})
			return
		}
		c.JSON(200, gin.H{
			"user": user,
		})
		return
	}
}

func (a *AuthContrller) Login() gin.HandlerFunc {

	type LoginBody struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required,min=8,max=255"`
	}
	return func(c *gin.Context) {

		var loginBody LoginBody
		if err := c.BindJSON(&loginBody); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		user, err := a.authService.Login(&loginBody.Email, &loginBody.Password)
		if err != nil {
			c.JSON(404, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(200, gin.H{
			"user": user,
		})
		return
	}
}
