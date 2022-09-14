package models

type Track struct {
	ID     int    `json:"id" gorm:"primary_key"`
	Artist string `json:"artist"`
	Title  string `json:"title"`
}
