package server

import (
	"github.com/gin-gonic/gin"

	"people-api/internal/handler"

	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	_ "people-api/docs"
)

func NewRouter(h *handler.PersonHandler) *gin.Engine {
	r := gin.Default()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	people := r.Group("/people")
	{
		people.POST("", h.CreatePerson)
		people.GET("", h.ListPeople)
		people.GET("/:id", h.GetPerson)
		people.PUT("/:id", h.UpdatePerson)
		people.DELETE("/:id", h.DeletePerson)
	}

	return r
}
