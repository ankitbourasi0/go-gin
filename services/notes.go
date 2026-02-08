package services

import (
	"fmt"
	"gin-tutorial/internal/models"
	"gorm.io/gorm"
)

// Services is used to Talk with Database
// We Write Queries Here or Business Logic
type NotesService struct {
	db *gorm.DB
}

func (n *NotesService) InitService(database *gorm.DB) {
	//it same like this.db = db
	n.db = database //initialize database
	//register model in the database
	n.db.AutoMigrate(&internal.Notes{})
}

type Note struct {
	Id   int
	Name string
}

func (n *NotesService) GetNotes() []Note {
	data := []Note{
		{Id: 1, Name: "test"},
		{Id: 2, Name: "test2"},
	}
	return data
}

// create notes in database
func (n *NotesService) CreateNotes() string {
	err := n.db.Create(&internal.Notes{Id: 1, Title: "test", Status: true}).Error
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return "Notes Create Successfully"
	
}
