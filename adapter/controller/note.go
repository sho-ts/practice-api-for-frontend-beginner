package controller

import (
	"github.com/gin-gonic/gin"
	input "note-app/domain/dto/input/note"
	"note-app/domain/object"
	"note-app/domain/usecase/note"
)

type noteController struct {
	findByNoteIdUseCase usecase.IFindByNoteIdUseCase
	findAllNoteUseCase  usecase.IFindAllNoteUseCase
	createNoteUseCase   usecase.ICreateNoteUseCase
	deleteNoteUseCase   usecase.IDeleteNoteUseCase
}

func NewNoteController(
	findByNoteIdUseCase usecase.IFindByNoteIdUseCase,
	findAllNoteUseCase usecase.IFindAllNoteUseCase,
	createNoteUseCase usecase.ICreateNoteUseCase,
	deleteNoteUseCase usecase.IDeleteNoteUseCase,
) noteController {
	return noteController{
		findByNoteIdUseCase: findByNoteIdUseCase,
		findAllNoteUseCase:  findAllNoteUseCase,
		createNoteUseCase:   createNoteUseCase,
		deleteNoteUseCase:   deleteNoteUseCase,
	}
}

func (nc noteController) CreateNote(c *gin.Context) {
	var rb struct {
		Title   string `json:"title"`
		Content string `json:"content"`
	}

	c.ShouldBindJSON(&rb)

	i := input.NewCreateNoteInput(rb.Title, rb.Content)
	o, o_err := nc.createNoteUseCase.Handle(i)

	if o_err != nil {
		c.JSON(405, gin.H{
			"messege": o_err.Error(),
		})
		return
	}

	c.JSON(200, o.Note)
}

func (nc noteController) FindByNoteId(c *gin.Context) {
	i := input.NewFindByNoteIdInput(c.Param("id"))
	o, err := nc.findByNoteIdUseCase.Handle(i)

	if err != nil {
		c.JSON(404, gin.H{
			"messege": "コンテンツが見つかりませんでした",
		})
		return
	}

	c.JSON(200, o.Note)
}

func (nc noteController) FindAllNote(c *gin.Context) {
	limit, l_err := object.NewLimit(c.Query("limit"))
	offset, o_err := object.NewOffset(c.Query("offset"))

	if l_err != nil || o_err != nil {
		c.JSON(405, gin.H{
			"messege": "クエリパラメータに問題があります",
		})
		return
	}

	i := input.NewFindAllNoteInput(limit, offset)
	o, i_err := nc.findAllNoteUseCase.Handle(i)

	if i_err != nil {
		c.JSON(500, gin.H{
			"messege": "エラーが発生しました",
		})
		return
	}

	c.JSON(200, o)
}

func (nc noteController) DeleteNote(c *gin.Context) {
  i := input.NewDeleteNoteInput(c.Param("id"))
  o, err := nc.deleteNoteUseCase.Handle(i)

  if (err != nil) {
    c.JSON(500, gin.H{
      "message": "削除に失敗しました",
    })
  }

  c.JSON(200, o)
}