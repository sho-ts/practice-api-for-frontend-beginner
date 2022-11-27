package input

import (
	"note-app/domain/object"
)

type FindAllNoteInput struct {
	Limit  object.Limit
	Offset object.Offset
}

func NewFindAllNoteInput(
	limit object.Limit,
	offset object.Offset,
) FindAllNoteInput {
	return FindAllNoteInput{
		Limit:  limit,
		Offset: offset,
	}
}
