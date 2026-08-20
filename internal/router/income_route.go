package router

import (
	"github.com/RicardoEmm/rosinas-food/internal/handler"
	"github.com/gin-gonic/gin"
)

func registerIncomeRoutes(rg *gin.RouterGroup, h *handler.IncomeHandler) {
	incomes := rg.Group("/incomes")
	{
		incomes.GET("/:id", h.FindById)
		incomes.GET("", h.FindAll)
		incomes.GET("/customer/:id", h.FindAllByCustomerId)
		incomes.POST("", h.Create)
		incomes.PATCH("/:id", h.ChangeToPaid)
		incomes.DELETE("/:id", h.DeleteById)
		incomes.DELETE("/customer/:id", h.DeleteAllByCustomerId)
	}
}
