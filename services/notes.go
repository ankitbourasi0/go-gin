package services

// Services is used to Talk with Database
// We Write Queries Here or Business Logic
type NotesService struct {
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
