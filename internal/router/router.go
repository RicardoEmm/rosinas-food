package router

import (
	"github.com/RicardoEmm/rosinas-food/internal/handler"
	"github.com/gin-gonic/gin"
)

type Handlers struct {
	CustomerHanlder *handler.CustomerHandler
}

func Setup(h Handlers) *gin.Engine {
	router := gin.Default()

	api := router.Group("/api/v1")
	{
		registerCustomerRoutes(api, h.CustomerHanlder)
	}
	return router
}
