package router

import (
	"github.com/RicardoEmm/rosinas-food/internal/handler"
	"github.com/gin-gonic/gin"
)

func registerCustomerRoutes(rg *gin.RouterGroup, h *handler.CustomerHandler) {
	customers := rg.Group("/customers")
	{
		customers.GET("/:id", h.FindById)
		customers.GET("", h.FindAll)
		customers.POST("", h.Create)
		customers.DELETE("/:id", h.DeleteById)
	}
}
