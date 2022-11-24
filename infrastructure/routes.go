package infrastructure

import (
	"github.com/gin-gonic/gin"
	// a "note-app/adapter"
)

func GetRouter() *gin.Engine {
	r := gin.Default()
	// p := r.Group("/v1")

	return r
}
