package controllers

//Instead of writing a lot of Routes we just categorize them into Controller
//This is an easy way to manage a lot of Routes in a Controller in future
import "github.com/gin-gonic/gin"

type NotesController struct {
}

func (n *NotesController) NewNotesController(router *gin.Engine) {
	//Group Create a new router group,
	//You should add all the routes that have common middlewares or the same path prefix.
	notes := router.Group("/notes")

	notes.GET("/", n.GetNumberOfNotes())
	notes.POST("/", n.CreateNotes())
}

func (n *NotesController) CreateNotes() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Notes Created",
		})
	}
}

func (n *NotesController) GetNumberOfNotes() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "5 Notes in the DB",
		})
	}
}
