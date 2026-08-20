// @title Rosinas Food API
// @version 1.0
// @description API to magage incomes, materials, and expenses
// @host localhost:8080
// @BasePath /api/v1
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

	// material
	materialRepo := repository.NewGormMaterialRepository(db)
	materialService := service.NewMaterialService(materialRepo)
	materialHandler := handler.NewMaterialHandler(materialService)

	// income
	incomeRepo := repository.NewGormIncomeRepositoty(db)
	incomeService := service.NewIncomeService(incomeRepo, customerRepo)
	incomeHandler := handler.NewIncomeHandler(incomeService)

	router := router.Setup(router.Handlers{
		CustomerHanlder: customerHandler,
		MaterialHandler: materialHandler,
		IncomeHandler:   incomeHandler,
	})

	router.Run(":" + cfg.AppPort)
}
