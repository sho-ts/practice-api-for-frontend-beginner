package interactor

import (
	"note-app/domain/dto/input/note"
	"note-app/domain/dto/output/note"
	"note-app/domain/repository"
)

type findAllNoteInteractor struct {
	noteRepository repository.INoteRepository
}

func NewFindAllNoteInteractor(
	noteRepository repository.INoteRepository,
) findAllNoteInteractor {
	return findAllNoteInteractor{
		noteRepository: noteRepository,
	}
}

func (i findAllNoteInteractor) Handle(in input.FindAllNoteInput) (output.FindAllNoteOutput, error) {
	notes, err := i.noteRepository.FindAllNote(in.Limit.Value, in.Offset.Value)

	return output.NewFindAllNoteOutput(notes), err
}
