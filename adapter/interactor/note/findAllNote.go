package interactor

import (
  "note-app/domain/dto/input/note"
  "note-app/domain/dto/output/note"
)

type findAllNoteInteractor struct{}

func NewFindAllNoteInteractor() findAllNoteInteractor {
  return findAllNoteInteractor{}
}

func (i findAllNoteInteractor) Handle(in input.FindAllNoteInput) output.FindAllNoteOutput {
  
}