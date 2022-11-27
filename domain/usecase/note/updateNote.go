package usecase

import (
  "note-app/domain/dto/input/note"
  "note-app/domain/dto/output/note"
)

type IUpdateNoteUseCase interface{
  Handle(input.UpdateNoteInput) (output.UpdateNoteOutput, error)
}