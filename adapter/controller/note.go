package controller

import (
	"github.com/gin-gonic/gin"
	input "note-app/domain/dto/input/note"
	"note-app/domain/usecase/note"
)

type noteController struct {
	findByNoteIdUseCase usecase.IFindByNoteIdUseCase
}

func NewNoteController(
	findByNoteIdUseCase usecase.IFindByNoteIdUseCase,
) noteController {
	return noteController{
		findByNoteIdUseCase: findByNoteIdUseCase,
	}
}

func (uc noteController) FindByNoteId(c *gin.Context) {
	i := input.NewFindByNoteIdInput(c.Param("id"))
	o, err := uc.findByNoteIdUseCase.Handle(i)

	if err != nil {
		c.JSON(404, gin.H{
			"messege": "コンテンツが見つかりませんでした",
		})
		return
	}

	c.JSON(200, o)
}
