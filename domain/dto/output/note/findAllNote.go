package output

import (
	"note-app/domain/entity"
)

type FindAllNoteOutput struct {
	Items []entity.Note `json:"items"`
}

func NewFindAllNoteOutput(
	items []entity.Note,
) FindAllNoteOutput {
	return FindAllNoteOutput{
		Items: items,
	}
}
