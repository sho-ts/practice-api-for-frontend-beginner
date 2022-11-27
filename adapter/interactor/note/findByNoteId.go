package interactor

import (
	"note-app/domain/dto/input/note"
	"note-app/domain/dto/output/note"
	"note-app/domain/repository"
)

type findByNoteIdInteractor struct {
	noteRepository repository.INoteRepository
}

func NewFindByNoteIdInteractor(
	noteRepository repository.INoteRepository,
) findByNoteIdInteractor {
	return findByNoteIdInteractor{
		noteRepository: noteRepository,
	}
}

func (i findByNoteIdInteractor) Handle(in input.FindByNoteIdInput) (output.FindByNoteIdOutput, error) {
	note, err := i.noteRepository.FindByNoteId(in.Id)

	return output.NewFindByNoteIdOutput(note), err
}
