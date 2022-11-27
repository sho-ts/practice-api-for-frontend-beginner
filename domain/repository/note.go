package repository

import (
	"note-app/domain/entity"
)

type INoteRepository interface {
	CreateNote(id string, title string, content string) (entity.Note, error)
	FindByNoteId(id string) (entity.Note, error)
	FindAllNote(limit int, offset int) ([]entity.Note, error)
	UpdateNote(id string, title string, content string) error
	DeleteNote(id string) error
}
