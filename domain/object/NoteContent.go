package object

import (
	"errors"
	"note-app/application/constant"
	"unicode/utf8"
)

type NoteContent struct {
	Value string
}

func NewNoteContent(content string) (NoteContent, error) {
	var noteContent NoteContent

	if content == "" {
		return noteContent, errors.New(constant.GetRequiredValidateErrorMessage("content"))
	}
	if utf8.RuneCountInString(content) > 100000 {
		return noteContent, errors.New(constant.GetMaximumCharVaridateErrorMessage("content", 100000))
	}

	noteContent.Value = content

	return noteContent, nil
}
