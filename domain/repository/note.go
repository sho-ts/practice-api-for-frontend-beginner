package repository

import (
	"note-app/domain/entity"
)

type INoteRepository interface {
	FindByNoteId(id string) (entity.Note, error)
	FindAllNote(limit int, offset int) ([]entity.Note, error)
}
