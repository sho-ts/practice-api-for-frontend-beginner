package interactor_test

import (
	"github.com/stretchr/testify/assert"
	"note-app/adapter/interactor/note"
	"note-app/domain/dto/input/note"
	"note-app/repository/mock"
	"testing"
)

func TestDeleteNote(t *testing.T) {
	interactor := interactor.NewDeleteNoteInteractor(repository.NewMockNoteRepository())

	id := "test_id"

	in := input.NewDeleteNoteInput(id)
	out, _ := interactor.Handle(in)

	assert.Equal(t, out.Id, id)
}
