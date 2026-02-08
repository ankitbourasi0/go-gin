package internal

type Notes struct {
	Id     int    `gorm:"primary_key;AUTO_INCREMENT"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}
