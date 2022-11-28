package interactor_test

import (
	interactor "note-app/adapter/interactor/note"
	"note-app/application/constant"
	input "note-app/domain/dto/input/note"
	repository "note-app/repository/mock"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateNoteInteractor(t *testing.T) {
	interactor := interactor.NewCreateNoteInteractor(repository.NewMockNoteRepository())

	title := "123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890"
	content := "test_content"

	in := input.NewCreateNoteInput(title, content)
	out, err := interactor.Handle(in)

	assert.NotEmpty(t, out.Note.Id)
	assert.Equal(t, out.Note.Title, title)
	assert.Equal(t, out.Note.Content, content)
	assert.Nil(t, err)
}

func TestCreateNoteInteractorRequiredTitleValue(t *testing.T) {
	interactor := interactor.NewCreateNoteInteractor(repository.NewMockNoteRepository())

	content := "test_content"

	in := input.NewCreateNoteInput("", content)
	_, err := interactor.Handle(in)

	assert.EqualError(t, err, constant.GetRequiredValidateErrorMessage("title"))
}

func TestCreateNoteInteractorRequiredContentValue(t *testing.T) {
	interactor := interactor.NewCreateNoteInteractor(repository.NewMockNoteRepository())

	title := "test_title"

	in := input.NewCreateNoteInput(title, "")
	_, err := interactor.Handle(in)

	assert.EqualError(t, err, constant.GetRequiredValidateErrorMessage("content"))
}

func TestCraeteNoteInteractorMaximumTitleValue(t *testing.T) {
	interactor := interactor.NewCreateNoteInteractor(repository.NewMockNoteRepository())

	title := "1234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901"
	content := "test_content"

	in := input.NewCreateNoteInput(title, content)
	_, err := interactor.Handle(in)

	assert.EqualError(t, err, constant.GetMaximumCharValidateErrorMessage("title", 120))
}
