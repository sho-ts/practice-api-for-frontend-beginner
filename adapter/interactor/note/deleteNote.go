package interactor

import (
  "note-app/domain/dto/input/note"
  "note-app/domain/dto/output/note"
)

type deleteNoteInteractor struct{}

func NewDeleteNoteInteractor() deleteNoteInteractor {
  return deleteNoteInteractor{}
}

func (i deleteNoteInteractor) Handle(in input.DeleteNoteInput) output.DeleteNoteOutput {
  return output.DeleteNoteOutput{}
}