package interactor

import (
	"note-app/domain/dto/input/note"
	"note-app/domain/dto/output/note"
	"note-app/domain/object"
	"note-app/domain/repository"
)

type updateNoteInteractor struct {
	noteRepository repository.INoteRepository
}

func NewUpdateNoteInteractor(
	noteRepository repository.INoteRepository,
) updateNoteInteractor {
	return updateNoteInteractor{
		noteRepository: noteRepository,
	}
}

func (i updateNoteInteractor) Handle(in input.UpdateNoteInput) (output.UpdateNoteOutput, error) {
	title, t_err := object.NewNoteTitle(in.Title)

	if t_err != nil {
		return output.UpdateNoteOutput{}, t_err
	}

	content, c_err := object.NewNoteContent(in.Content)

	if c_err != nil {
		return output.UpdateNoteOutput{}, c_err
	}

	err := i.noteRepository.UpdateNote(in.Id, title.Value, content.Value)

	return output.NewUpdateNoteOutput(in.Id), err
}
