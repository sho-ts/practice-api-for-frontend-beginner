package object

import (
	"errors"
	"note-app/application/constant"
	"unicode/utf8"
)

type NoteTitle struct {
	Value string
}

func NewNoteTitle(title string) (NoteTitle, error) {
	var noteTitle NoteTitle

	if title == "" {
		return noteTitle, errors.New(constant.GetRequiredValidateErrorMessage("title"))
	}
	if utf8.RuneCountInString(title) > 120 {
		return noteTitle, errors.New(constant.GetMaximumCharValidateErrorMessage("title", 120))
	}

	noteTitle.Value = title

	return noteTitle, nil
}
