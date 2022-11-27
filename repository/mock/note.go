package repository

import (
	"note-app/domain/entity"
	"strconv"
)

type MockNoteRepository struct{}

func NewMockNoteRepository() MockNoteRepository {
  return MockNoteRepository{}
}

func (r MockNoteRepository) CreateNote(
	id string, title string, content string,
) (entity.Note, error) {
	return entity.Note{
		Id:      id,
		Title:   title,
		Content: content,
	}, nil
}

func (r MockNoteRepository) FindByNoteId(id string) (entity.Note, error) {
	return entity.Note{
		Id:      id,
		Title:   "mock-note-title",
		Content: "mock-note-content",
	}, nil
}

func (r MockNoteRepository) FindAllNote(
	limit int, offset int,
) ([]entity.Note, error) {
  var mockNotes = make([]entity.Note, limit)

  for i := 0; i < limit; i++ {
    mockNotes[i] = entity.Note{
      Id:      "mock-note-id" + strconv.Itoa(i),
      Title:   "mock-note-title",
      Content: "mock-note-content",
    }
  }

  return mockNotes, nil
}

func (r MockNoteRepository) UpdateNote(
  id string,
  title string,
  content string,
) error {
  return nil
}

func (r MockNoteRepository) DeleteNote(
  id string,
) error {
  return nil
}