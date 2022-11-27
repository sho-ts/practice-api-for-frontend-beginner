package entity

import (
	"note-app/domain/object"
	"time"
)

type Note struct {
	Id        string    `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func NewNote(title object.NoteTitle, content object.NoteContent) (Note, error) {
	var note Note

	note.Id = object.NewNoteId().Value
	note.Title = title.Value
	note.Content = content.Value

	return note, nil
}
