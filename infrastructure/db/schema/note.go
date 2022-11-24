package schema

type Note struct {
	Id      string `gorm:"size:255;primary_key;unique"`
	Title   string `gorm:"size: 128;"`
	Content string
	Schema
}
