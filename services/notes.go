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

func (n *NotesService) GetNotes(status bool) ([]*internal.Notes, error) {
	var notes []*internal.Notes
	if err := n.db.Where("status = ?", status).Find(&notes).Error; err != nil {
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

func (n *NotesService) UpdateNotes(title string, status bool, id int) (*internal.Notes, error) {
	var note internal.Notes

	if err := n.db.Where("id = ?", id).First(&note).Error; err != nil {

		fmt.Println("Update note :Error in fetching ", err)

		return nil, err
	}
	note.Title = title
	note.Status = status
	if err := n.db.Save(&note).Error; err != nil {
		fmt.Println("Update note :Error in Saving ", err)
		return nil, err
	}
	return &note, nil
}

func (n *NotesService) DeleteNotes(id int64) error {
	var note internal.Notes
	if err := n.db.Where("id = ?", id).First(&note).Error; err != nil {

		fmt.Println("Delete note :Error in fetching ", err)

		return err
	}
	if err := n.db.Where("id = ?", id).Delete(note).Error; err != nil {
		fmt.Println("Delete Notes : Error in Deleting", err)
		return err
	}
	return nil
}
