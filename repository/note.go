package repository

import (
	"note-app/domain/entity"
	"note-app/infrastructure/db/schema"
)

type NoteRepository struct {
	table string
}

func NewNoteRepository() NoteRepository {
	return NoteRepository{
		table: "notes",
	}
}

func (n NoteRepository) CreateNote(id string, title string, content string) (entity.Note, error) {
	note := schema.Note{
		Id:      id,
		Title:   title,
		Content: content,
	}

	r := Repository.Create(&note)

	return entity.Note{
		Id:        note.Id,
		Title:     note.Title,
		Content:   note.Content,
		CreatedAt: note.CreatedAt,
		UpdatedAt: note.UpdatedAt,
	}, r.Error
}

func (n NoteRepository) FindByNoteId(id string) (entity.Note, error) {
	var note = entity.Note{}

	r := Repository.
		Debug().
		Table(n.table).
		Select([]string{
			"id as Id",
			"title as Title",
			"content as Content",
			"created_at as CreatedAt",
			"updated_at as UpdatedAt",
		}).
		Where("id = ?", id).
		Where("deleted_at is null").
		First(&note)

	return note, r.Error
}

func (n NoteRepository) FindAllNote(limit int, offset int) ([]entity.Note, int64, error) {
	var notes = []entity.Note{}
	var total int64

	q := Repository.
		Table(n.table).
		Select([]string{
			"id as Id",
			"title as Title",
			"content as Content",
			"created_at as CreatedAt",
			"updated_at as UpdatedAt",
		}).
		Where("deleted_at is null")

	q.Count(&total)
	r := q.
		Limit(limit).
		Offset(offset).
		Scan(&notes)

	return notes, total, r.Error
}

func (n NoteRepository) UpdateNote(id string, title string, content string) error {
	var note schema.Note

	r := Repository.Model(&note).
		Where("id = ?", id).
		Where("deleted_at is null").
		Updates(schema.Note{
			Title:   title,
			Content: content,
		})

	return r.Error
}

func (n NoteRepository) DeleteNote(id string) error {
	r := Repository.
		Where("id = ?", id).
		Delete(&schema.Note{})

	return r.Error
}
