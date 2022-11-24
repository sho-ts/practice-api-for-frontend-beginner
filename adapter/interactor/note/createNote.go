package interactor

import (
  "note-app/domain/dto/input/note"
  "note-app/domain/dto/output/note"
)

type createNoteInteractor struct{}

func NewCreateNoteInteractor() createNoteInteractor {
  return createNoteInteractor{}
}

func (i createNoteInteractor) Handle(in input.CreateNoteInput) output.CreateNoteOutput {
    return output.CreateNoteOutput{}
}