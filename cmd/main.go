package main

import (
	"github.com/RicardoEmm/rosinas-food/internal/config"
	"github.com/RicardoEmm/rosinas-food/internal/database"
	"github.com/RicardoEmm/rosinas-food/internal/handler"
	"github.com/RicardoEmm/rosinas-food/internal/repository"
	"github.com/RicardoEmm/rosinas-food/internal/router"
	"github.com/RicardoEmm/rosinas-food/internal/service"
)

func main() {
	cfg := config.Load()
	db := database.Connect(cfg)

	// customer
	customerRepo := repository.NewGormCustomerRepository(db)
	customerService := service.NewCustomerService(customerRepo)
	customerHandler := handler.NewCustomerHandler(customerService)

	router := router.Setup(router.Handlers{
		CustomerHanlder: customerHandler,
	})

	router.Run(":" + cfg.AppPort)
}
