package output

import "note-app/domain/entity"

type CreateNoteOutput struct {
	Note entity.Note
}

func NewCreateNoteOutput(note entity.Note) CreateNoteOutput {
	return CreateNoteOutput{
		Note: note,
	}
}
