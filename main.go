package main

import (
	"fmt"
	"gin-tutorial/controllers"
	"gin-tutorial/internal/database"
	"gin-tutorial/services"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func main() {

	db := DatabaseInit()                     //Database initialize
	router := gin.Default()                  // Router Create
	notesService := &services.NotesService{} //Get Service
	notesService.InitService(db)             //Service Initialize

	notesController := &controllers.NotesController{} //Access CONTROLLER
	notesController.InitController(*notesService)     //Controller Initialize
	notesController.InitRoutes(router)                // Routes Initialize

	//Auth Service
	authService := &services.AuthService{}
	authService.InitAuthService(db)
	//Auth Controller
	authController := &controllers.AuthContrller{}
	authController = authController.InitAuthController(*authService)
	authController.InitRoutes(router)

	if err := router.Run(); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}

}

func DatabaseInit() *gorm.DB {
	db := internal.InitDatabase()

	if db == nil {
		fmt.Println("Database initialization failed!!!")
	}

	fmt.Println("Database Initialization Successful", db)
	return db
}
