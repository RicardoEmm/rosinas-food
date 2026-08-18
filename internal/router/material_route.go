package router

import (
	"github.com/RicardoEmm/rosinas-food/internal/handler"
	"github.com/gin-gonic/gin"
)

func registerMaterialRoutes(rg *gin.RouterGroup, h *handler.MaterialHandler) {
	materials := rg.Group("/materials")
	{
		materials.GET("/:id", h.FindById)
		materials.GET("", h.FindAll)
		materials.POST("", h.Create)
		materials.DELETE("/:id", h.DeleteById)
	}
}
