package controllers

//Instead of writing a lot of Routes we just categorize them into Controller
//This is an easy way to manage a lot of Routes in a Controller in future
import (
	"gin-tutorial/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type NotesController struct {
	notesService services.NotesService
}

func (n *NotesController) NewNotesController(router *gin.Engine, notesService services.NotesService) {
	//Group Create a new router group,
	//You should add all the routes that have common middlewares or the same path prefix.
	notes := router.Group("/notes")

	notes.GET("/", n.GetNumberOfNotes())
	notes.POST("/", n.CreateNotes())
	notes.GET("/getFromService", n.GetDataFromNotesService())
	n.notesService = notesService
}

func (n *NotesController) CreateNotes() gin.HandlerFunc {
	type NoteBody struct {
		Title  string `json:"title"`
		Status bool   `json:"status"`
	}

	return func(c *gin.Context) {
		var noteBody NoteBody
		//parse the request body in NoteBody
		if err := c.BindJSON(&noteBody); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		note, err := n.notesService.CreateNotes(noteBody.Title, noteBody.Status)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		c.JSON(200, gin.H{
			"message": "Note Created",
			"data":    note,
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

func (n *NotesController) GetDataFromNotesService() gin.HandlerFunc {
	return func(c *gin.Context) {
		notes, err := n.notesService.GetNotes()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		c.JSON(200, gin.H{
			"message": notes,
		})
	}
}
