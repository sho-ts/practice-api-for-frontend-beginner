package useCase

import (
  "note-app/domain/dto/input/note"
  "note-app/domain/dto/output/note"
)

type ICreateNoteUseCase interface{
  Handle(input.CreateNoteInput) output.CreateNoteOutput
}