package interactor_test

import (
	"github.com/stretchr/testify/assert"
	"note-app/adapter/interactor/note"
	"note-app/domain/dto/input/note"
	"note-app/repository/mock"
	"testing"
)

func TestFindByNoteId(t *testing.T) {
	interactor := interactor.NewFindByNoteIdInteractor(repository.NewMockNoteRepository())

	id := "test_id"

	in := input.NewFindByNoteIdInput(id)
	out, _ := interactor.Handle(in)

	assert.Equal(t, out.Note.Id, id)
}
