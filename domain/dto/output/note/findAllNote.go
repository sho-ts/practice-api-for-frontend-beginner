package output

import (
	"note-app/domain/entity"
)

type FindAllNoteOutput struct {
	Items []entity.Note `json:"items"`
	Total int64         `json:"total"`
}

func NewFindAllNoteOutput(
	items []entity.Note,
	total int64,
) FindAllNoteOutput {
	return FindAllNoteOutput{
		Items: items,
		Total: total,
	}
}
