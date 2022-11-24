package interactor

import (
  "note-app/domain/dto/input/note"
  "note-app/domain/dto/output/note"
)

type updateNoteInteractor struct{}

func NewUpdateNoteInteractor() updateNoteInteractor {
  return updateNoteInteractor{}
}

func (i updateNoteInteractor) Handle(in input.UpdateNoteInput) output.UpdateNoteOutput {
  return output.UpdateNoteOutput{}
}