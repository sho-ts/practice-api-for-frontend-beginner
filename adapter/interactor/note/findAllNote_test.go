package interactor_test

import (
	"github.com/stretchr/testify/assert"
	"note-app/adapter/interactor/note"
	"note-app/domain/dto/input/note"
	"note-app/domain/object"
	"note-app/repository/mock"
	"testing"
)

func TestFindAllNote(t *testing.T) {
	interactor := interactor.NewFindAllNoteInteractor(repository.NewMockNoteRepository())

	limit, _ := object.NewLimit("10")
	offset, _ := object.NewOffset("1")
	var total int64 = 100

	in := input.NewFindAllNoteInput(limit, offset)
	out, _ := interactor.Handle(in)

	assert.Equal(t, out.Items[0].Id, "mock-note-id-0")
	assert.Equal(t, out.Items[0].Title, "mock-note-title")
	assert.Equal(t, out.Items[0].Content, "mock-note-content")
	assert.Equal(t, out.Total, total)
	assert.Equal(t, len(out.Items), 10)
}
