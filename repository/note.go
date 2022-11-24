package repository

type noteRepository struct{}

func NewNoteRepository() noteRepository {
  return noteRepository{}
}