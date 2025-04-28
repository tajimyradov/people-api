package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"people-api/internal/domain"
	"people-api/internal/service"
)

type PersonHandler struct {
	service *service.PersonService
}

func NewPersonHandler(service *service.PersonService) *PersonHandler {
	return &PersonHandler{service: service}
}

// @Summary Create Person
// @Tags People
// @Accept json
// @Produce json
// @Param person body domain.Person true "Person Info"
// @Success 200 {object} map[string]int
// @Router /people [post]
func (h *PersonHandler) CreatePerson(c *gin.Context) {
	var input domain.Person
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := h.service.Create(&input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": id})
}

// ListPeople - получить список с фильтрацией и пагинацией
// @Summary List People
// @Tags People
// @Produce json
// @Param name query string false "Name"
// @Param surname query string false "Surname"
// @Param nationality query string false "Nationality"
// @Param page query int false "Page number" default(1)
// @Param limit query int false "Items per page" default(10)
// @Success 200 {array} domain.Person
// @Router /people [get]
func (h *PersonHandler) ListPeople(c *gin.Context) {
	name := c.Query("name")
	surname := c.Query("surname")
	nationality := c.Query("nationality")

	page := 1
	limit := 10

	if p := c.Query("page"); p != "" {
		fmt.Sscanf(p, "%d", &page)
	}
	if l := c.Query("limit"); l != "" {
		fmt.Sscanf(l, "%d", &limit)
	}

	people, err := h.service.List(name, surname, nationality, page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, people)
}

// GetPerson - получить по ID
// @Summary Get Person by ID
// @Tags People
// @Produce json
// @Param id path int true "Person ID"
// @Success 200 {object} domain.Person
// @Router /people/{id} [get]
func (h *PersonHandler) GetPerson(c *gin.Context) {
	id := c.Param("id")

	person, err := h.service.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Person not found"})
		return
	}

	c.JSON(http.StatusOK, person)
}

// UpdatePerson - обновить человека
// @Summary Update Person
// @Tags People
// @Accept json
// @Produce json
// @Param id path int true "Person ID"
// @Param person body domain.Person true "Person Info"
// @Success 200 {object} map[string]string
// @Router /people/{id} [put]
func (h *PersonHandler) UpdatePerson(c *gin.Context) {
	id := c.Param("id")
	var input domain.Person

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.Update(id, &input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "updated"})
}

// DeletePerson - удалить человека
// @Summary Delete Person
// @Tags People
// @Param id path int true "Person ID"
// @Success 200 {object} map[string]string
// @Router /people/{id} [delete]
func (h *PersonHandler) DeletePerson(c *gin.Context) {
	id := c.Param("id")

	if err := h.service.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "deleted"})
}
