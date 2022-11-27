package controller

import (
	"github.com/gin-gonic/gin"
	"note-app/application/constant"
	input "note-app/domain/dto/input/note"
	"note-app/domain/object"
	"note-app/domain/usecase/note"
)

type noteController struct {
	createNoteUseCase   usecase.ICreateNoteUseCase
	findByNoteIdUseCase usecase.IFindByNoteIdUseCase
	findAllNoteUseCase  usecase.IFindAllNoteUseCase
	updateNoteUseCase   usecase.IUpdateNoteUseCase
	deleteNoteUseCase   usecase.IDeleteNoteUseCase
}

func NewNoteController(
	createNoteUseCase usecase.ICreateNoteUseCase,
	findByNoteIdUseCase usecase.IFindByNoteIdUseCase,
	findAllNoteUseCase usecase.IFindAllNoteUseCase,
	updateNoteUseCase usecase.IUpdateNoteUseCase,
	deleteNoteUseCase usecase.IDeleteNoteUseCase,
) noteController {
	return noteController{
		createNoteUseCase:   createNoteUseCase,
		findByNoteIdUseCase: findByNoteIdUseCase,
		findAllNoteUseCase:  findAllNoteUseCase,
		updateNoteUseCase:   updateNoteUseCase,
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
	o, err := nc.createNoteUseCase.Handle(i)

	if err != nil {
		c.JSON(405, gin.H{
			"messege": err.Error(),
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
			"messege": constant.MESSAGE_NOT_FOUND,
		})
		return
	}

	c.JSON(200, o.Note)
}

func (nc noteController) FindAllNote(c *gin.Context) {
	limit, l_err := object.NewLimit(c.Query("limit"))
	offset, o_err := object.NewOffset(c.Query("offset"), limit)

	if l_err != nil || o_err != nil {
		c.JSON(405, gin.H{
			"messege": constant.MESSAGE_HAS_QUERY_VALUE_PROBLEM,
		})
		return
	}

	i := input.NewFindAllNoteInput(limit, offset)
	o, i_err := nc.findAllNoteUseCase.Handle(i)

	if i_err != nil {
		c.JSON(500, gin.H{
			"messege": constant.MESSAGE_UNKNOWN_ERROR,
		})
		return
	}

	c.JSON(200, o)
}

func (nc noteController) UpdateNote(c *gin.Context) {
	var rb struct {
		Title   string `json:"title"`
		Content string `json:"content"`
	}

	c.ShouldBindJSON(&rb)

	i := input.NewUpdateNoteInput(
		c.Param("id"),
		rb.Title,
		rb.Content,
	)
	o, err := nc.updateNoteUseCase.Handle(i)

	if err != nil {
		c.JSON(405, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(200, o)
}

func (nc noteController) DeleteNote(c *gin.Context) {
	i := input.NewDeleteNoteInput(c.Param("id"))
	o, err := nc.deleteNoteUseCase.Handle(i)

	if err != nil {
		c.JSON(500, gin.H{
			"message": constant.MESSAGE_UNKNOWN_ERROR,
		})
		return
	}

	c.JSON(200, o)
}
