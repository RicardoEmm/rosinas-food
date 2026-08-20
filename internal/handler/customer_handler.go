package handler

import (
	"net/http"

	"github.com/RicardoEmm/rosinas-food/internal/dto"
	"github.com/RicardoEmm/rosinas-food/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type CustomerHandler struct {
	customerService *service.CustomerService
}

func NewCustomerHandler(customerService *service.CustomerService) *CustomerHandler {
	return &CustomerHandler{customerService: customerService}
}

func (h *CustomerHandler) FindById(c *gin.Context) {
	parsedID, err := h.ParseUUID(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
		return
	}

	customer, err := h.customerService.FindById(c.Request.Context(), parsedID)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": customer})
}

func (h *CustomerHandler) FindAll(c *gin.Context) {
	customers, err := h.customerService.FindAll(c.Request.Context())

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": customers})
}

func (h *CustomerHandler) Create(c *gin.Context) {
	var req dto.CustomerDTO
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	parsedPrice, err := decimal.NewFromString(req.Price)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid price value"})
		return
	}

	input := dto.CreateCustomerInput{
		FullName: req.FullName,
		Phone:    req.Phone,
		Price:    parsedPrice,
	}

	if err := h.customerService.Create(c.Request.Context(), input); err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": "customer was created"})
}

func (h *CustomerHandler) DeleteById(c *gin.Context) {
	parsedID, err := h.ParseUUID(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
		return
	}

	if err := h.customerService.DeleteById(c.Request.Context(), parsedID); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": "customer was deleted"})
}

func (h *CustomerHandler) ParseUUID(id string) (uuid.UUID, error) {
	parsedID, err := uuid.Parse(id)

	if err != nil {
		return uuid.UUID{}, err
	}

	return parsedID, nil
}
