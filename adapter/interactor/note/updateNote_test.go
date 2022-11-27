package interactor_test

import (
	"github.com/stretchr/testify/assert"
	"note-app/adapter/interactor/note"
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
