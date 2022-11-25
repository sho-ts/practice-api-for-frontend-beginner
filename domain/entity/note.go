package entity

import (
	"errors"
	"time"
	"unicode/utf8"
  "note-app/domain/object"
)

type Note struct {
	Id        string    `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"createdAt"`
  UpdatedAt time.Time `json:"updatedAt"`
}

func NewNote(title string, content string) (Note, error) {
	var note Note

	if title == "" {
		return note, errors.New("titleは必須です")
	}
	if utf8.RuneCountInString(title) > 120 {
		return note, errors.New("titleは120文字以内である必要があります")
	}
	if content == "" {
		return note, errors.New("contentは必須です")
	}
  if utf8.RuneCountInString(content) > 100000 {
		return note, errors.New("contentは100000文字以内である必要があります")
	}

  note.Id = object.NewNoteId().Value
	note.Title = title
	note.Content = content

	return note, nil
}
