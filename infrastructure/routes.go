package infrastructure

import (
	"github.com/gin-gonic/gin"
	a "note-app/adapter"
)

func GetRouter() *gin.Engine {
	r := gin.Default()
  r.Use(GetCorsConfig())
	v1 := r.Group("/v1")

	v1.GET("/note", a.NoteController.FindAllNote)
	v1.POST("/note", a.NoteController.CreateNote)
	v1.GET("/note/:id", a.NoteController.FindByNoteId)
	v1.PUT("/note/:id", a.NoteController.UpdateNote)
	v1.DELETE("/note/:id", a.NoteController.DeleteNote)

	return r
}
