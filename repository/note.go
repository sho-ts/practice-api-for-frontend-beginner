package repository

import (
	"note-app/domain/entity"
)

type NoteRepository struct {
	table string
}

func NewNoteRepository() NoteRepository {
	return NoteRepository{
		table: "notes",
	}
}

func (n NoteRepository) FindByNoteId(id string) (entity.Note, error) {
	var note = entity.Note{}

	r := Repository.
		Table(n.table).
		Select([]string{
			"id as Id",
			"title as Title",
			"content as Content",
			"created_at as CreatedAt",
		}).
		Where("id = ?", id).
		First(&note)

	return note, r.Error
}

func (n NoteRepository) FindAllNote(limit int, offset int) ([]entity.Note, error) {
	var notes = []entity.Note{}

	r := Repository.
		Debug().
		Table(n.table).
		Select([]string{
			"id as Id",
			"title as Title",
			"content as Content",
			"created_at as CreatedAt",
		}).
    Limit(limit).
    Offset(offset).
		Scan(&notes)

	return notes, r.Error
}
