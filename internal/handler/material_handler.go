package handler

import (
	"net/http"
	"strconv"

	"github.com/RicardoEmm/rosinas-food/internal/service"
	"github.com/gin-gonic/gin"
)

type MaterialHandler struct {
	materialService *service.MaterialService
}

type CreateMaterialDTO struct {
	Name string `json:"name" binding:"required,min=1,max=50"`
}

func NewMaterialHandler(materialService *service.MaterialService) *MaterialHandler {
	return &MaterialHandler{materialService: materialService}
}

func (h *MaterialHandler) FindById(c *gin.Context) {
	parsedID, err := strconv.Atoi("id")

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
		return
	}

	material, err := h.materialService.FindById(c.Request.Context(), uint(parsedID))

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": material})
}

func (h *MaterialHandler) FindAll(c *gin.Context) {
	materials, err := h.materialService.FindAll(c.Request.Context())

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": materials})
}

func (h *MaterialHandler) Create(c *gin.Context) {
	var req CreateMaterialDTO

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	input := &service.CreateMaterialInput{Name: req.Name}

	if err := h.materialService.Create(c.Request.Context(), *input); err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": "material was created"})
}

func (h *MaterialHandler) DeleteById(c *gin.Context) {
	parsedID, err := strconv.Atoi("id")

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
		return
	}

	if err := h.materialService.DeleteById(c.Request.Context(), uint(parsedID)); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": "material was deleted"})
}
