package object

import (
	"errors"
	"unicode/utf8"
)

type NoteContent struct {
	Value string
}

func NewNoteContent(content string) (NoteContent, error) {
	var noteContent NoteContent

	if content == "" {
		return noteContent, errors.New("contentは必須です")
	}
	if utf8.RuneCountInString(content) > 100000 {
		return noteContent, errors.New("contentは100000文字以内である必要があります")
	}

	noteContent.Value = content

	return noteContent, nil
}
