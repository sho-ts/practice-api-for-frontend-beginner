package object

import (
	"errors"
	"unicode/utf8"
)

type NoteTitle struct {
	Value string
}

func NewNoteTitle(title string) (NoteTitle, error) {
	var noteTitle NoteTitle

	if title == "" {
		return noteTitle, errors.New("titleは必須です")
	}
	if utf8.RuneCountInString(title) > 120 {
		return noteTitle, errors.New("titleは120文字以内である必要があります")
	}

	noteTitle.Value = title

	return noteTitle, nil
}
