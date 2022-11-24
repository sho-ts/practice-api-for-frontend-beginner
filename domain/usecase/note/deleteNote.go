package useCase

import (
  "note-app/domain/dto/input/note"
  "note-app/domain/dto/output/note"
)

type IDeleteNoteUseCase interface{
  Handle(input.DeleteNoteInput) output.DeleteNoteOutput
}