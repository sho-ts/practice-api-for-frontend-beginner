package entity

import "time"

type Note struct {
	Id        string
	Title     string
	Content   string
	CreatedAt time.Time
}
