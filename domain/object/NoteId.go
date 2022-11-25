package object

import (
	"github.com/oklog/ulid"
	"math/rand"
	"time"
)

type NoteId struct {
	Value string
}

func NewNoteId() NoteId {
	entropy := rand.New(rand.NewSource(time.Now().UnixNano()))
	ms := ulid.Timestamp(time.Now())
	id, _ := ulid.New(ms, entropy)

	return NoteId{Value: id.String()}
}
