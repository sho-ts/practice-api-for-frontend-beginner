package interactor

import (
  "note-app/domain/dto/input/note"
  "note-app/domain/dto/output/note"
)

type findByNoteIdInteractor struct{}

func NewFindByNoteIdInteractor() findByNoteIdInteractor {
  return findByNoteIdInteractor{}
}

func (i findByNoteIdInteractor) Handle(in input.FindByNoteIdInput) output.FindByNoteIdOutput {
  
}