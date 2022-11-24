package useCase

import (
  "note-app/domain/dto/input/note"
  "note-app/domain/dto/output/note"
)

type IFindAllNoteUseCase interface{
  Handle(input.FindAllNoteInput) output.FindAllNoteOutput
}