package usecase

import (
  "note-app/domain/dto/input/note"
  "note-app/domain/dto/output/note"
)

type IFindByNoteIdUseCase interface{
  Handle(input.FindByNoteIdInput) output.FindByNoteIdOutput
}