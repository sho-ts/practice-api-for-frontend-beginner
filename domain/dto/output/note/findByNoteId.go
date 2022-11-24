package output

import (
	"note-app/domain/entity"
	"time"
)

type FindByNoteIdOutput struct {
	Id        string    `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"createdAt"`
}

func NewFindByNoteIdOutput(note entity.Note) FindByNoteIdOutput {
	return FindByNoteIdOutput{
		Id:        note.Id,
		Title:     note.Title,
		Content:   note.Content,
		CreatedAt: note.CreatedAt,
	}
}
