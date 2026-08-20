package handler

import (
	"net/http"

	"github.com/RicardoEmm/rosinas-food/internal/dto"
	"github.com/RicardoEmm/rosinas-food/internal/service"
	"github.com/RicardoEmm/rosinas-food/pkg"
	"github.com/gin-gonic/gin"
)

type IncomeHandler struct {
	incomeService *service.IncomeService
}

func NewIncomeHandler(incomeService *service.IncomeService) *IncomeHandler {
	return &IncomeHandler{incomeService: incomeService}
}

func (h *IncomeHandler) FindById(c *gin.Context) {
	parsedID, err := pkg.ParseUUID(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	income, err := h.incomeService.FindById(c.Request.Context(), parsedID)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": income})
}

func (h *IncomeHandler) FindAll(c *gin.Context) {
	incomes, err := h.incomeService.FindAll(c.Request.Context())

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": incomes})
}

func (h *IncomeHandler) FindAllByCustomerId(c *gin.Context) {
	parsedID, err := pkg.ParseUUID(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	incomes, err := h.incomeService.FindAllByCustomerId(c.Request.Context(), parsedID)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": incomes})
}

func (h *IncomeHandler) Create(c *gin.Context) {
	var req dto.IncomeRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.incomeService.Create(c.Request.Context(), dto.CreateIncomeInput{
		Description: req.Description,
		ProductType: req.ProductType,
		UnitPrice:   req.UnitPrice,
		Quantity:    req.Quantity,
		CustomerID:  req.CustomerID,
		Status:      req.Status,
	}); err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": "income was created"})
}

func (h *IncomeHandler) ChangeToPaid(c *gin.Context) {
	parsedID, err := pkg.ParseUUID(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.incomeService.ChangeToPaid(c.Request.Context(), parsedID); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": "income was masked as paid"})
}

func (h *IncomeHandler) DeleteById(c *gin.Context) {
	parsedID, err := pkg.ParseUUID(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.incomeService.DeleteById(c.Request.Context(), parsedID); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": "income was deleted"})
}

func (h *IncomeHandler) DeleteAllByCustomerId(c *gin.Context) {
	parsedID, err := pkg.ParseUUID(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.incomeService.DeleteAllByCustomerId(c.Request.Context(), parsedID); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": "incomes was deleted"})
}
