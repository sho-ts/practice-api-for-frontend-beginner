package interactor

import (
	"note-app/domain/dto/input/note"
	"note-app/domain/dto/output/note"
	"note-app/domain/repository"
)

type deleteNoteInteractor struct {
	noteRepository repository.INoteRepository
}

func NewDeleteNoteInteractor(
	noteRepository repository.INoteRepository,
) deleteNoteInteractor {
	return deleteNoteInteractor{
		noteRepository: noteRepository,
	}
}

func (i deleteNoteInteractor) Handle(in input.DeleteNoteInput) (output.DeleteNoteOutput, error) {
	err := i.noteRepository.DeleteNote(in.Id)

	return output.NewDeleteNoteOutput(in.Id), err
}
