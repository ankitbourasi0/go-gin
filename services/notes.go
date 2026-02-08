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

func (n *NotesService) GetNotes() ([]*internal.Notes, error) {
	var notes []*internal.Notes
	if err := n.db.Find(&notes).Error; err != nil {
		fmt.Println(err)
		return nil, err
	}
	return notes, nil
}

// create notes in db, data coming from frontend
func (n *NotesService) CreateNotes(title string, status bool) (*internal.Notes, error) {
	note := &internal.Notes{
		Title:  title,
		Status: status,
	}
	err := n.db.Create(note).Error
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return note, nil

}
