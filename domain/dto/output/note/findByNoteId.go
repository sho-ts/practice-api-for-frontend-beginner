package output

import (
	"note-app/domain/entity"
)

type FindByNoteIdOutput struct {
	Note entity.Note
}

func NewFindByNoteIdOutput(note entity.Note) FindByNoteIdOutput {
	return FindByNoteIdOutput{
		Note: note,
	}
}
