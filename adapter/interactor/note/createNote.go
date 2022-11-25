package interactor

import (
	"note-app/domain/dto/input/note"
	"note-app/domain/dto/output/note"
	"note-app/domain/entity"
	"note-app/domain/repository"
)

type createNoteInteractor struct {
	noteRepository repository.INoteRepository
}

func NewCreateNoteInteractor(
	noteRepository repository.INoteRepository,
) createNoteInteractor {
	return createNoteInteractor{
		noteRepository: noteRepository,
	}
}

func (i createNoteInteractor) Handle(in input.CreateNoteInput) (output.CreateNoteOutput, error) {
	note, n_err := entity.NewNote(in.Title, in.Content)

	if n_err != nil {
		return output.CreateNoteOutput{}, n_err
	}

	note, r_err := i.noteRepository.CreateNote(note.Id, note.Title, note.Content)

	return output.NewCreateNoteOutput(note), r_err
}
