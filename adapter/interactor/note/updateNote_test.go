package interactor_test

import (
	"github.com/stretchr/testify/assert"
	"note-app/adapter/interactor/note"
	"note-app/application/constant"
	"note-app/domain/dto/input/note"
	"note-app/repository/mock"
	"testing"
)

func TestUpdateNote(t *testing.T) {
	interactor := interactor.NewUpdateNoteInteractor(repository.NewMockNoteRepository())

	id := "test_id"
	title := "test_title"
	content := "test_content"

	in := input.NewUpdateNoteInput(id, title, content)
	out, _ := interactor.Handle(in)

	assert.Equal(t, out.Id, id)
}

func TestUpdateNoteInteractorRequiredTitleValue(t *testing.T) {
	interactor := interactor.NewUpdateNoteInteractor(repository.NewMockNoteRepository())

	id := "test_id"
	content := "test_content"

	in := input.NewUpdateNoteInput(id, "", content)
	_, err := interactor.Handle(in)

	assert.EqualError(t, err, constant.GetRequiredValidateErrorMessage("title"))
}

func TestUpdateNoteInteractorRequiredContentValue(t *testing.T) {
	interactor := interactor.NewUpdateNoteInteractor(repository.NewMockNoteRepository())

	id := "test_id"
	title := "test_title"

	in := input.NewUpdateNoteInput(id, title, "")
	_, err := interactor.Handle(in)

	assert.EqualError(t, err, constant.GetRequiredValidateErrorMessage("content"))
}

func TestUpdateNoteInteractorMaximumTitleValue(t *testing.T) {
	interactor := interactor.NewUpdateNoteInteractor(repository.NewMockNoteRepository())

	id := "test_id"
	title := "1234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901"
	content := "test_content"

	in := input.NewUpdateNoteInput(id, title, content)
	_, err := interactor.Handle(in)

	assert.EqualError(t, err, constant.GetMaximumCharVaridateErrorMessage("title", 120))
}
