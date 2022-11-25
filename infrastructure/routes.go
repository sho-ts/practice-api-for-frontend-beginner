package infrastructure

import (
	"github.com/gin-gonic/gin"
	a "note-app/adapter"
)

func GetRouter() *gin.Engine {
	r := gin.Default()

	p := r.Group("/v1")
	p.GET("/note", a.NoteController.FindAllNote)
	p.POST("/note", a.NoteController.CreateNote)
	p.GET("/note/:id", a.NoteController.FindByNoteId)

	return r
}
