package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shorthack_backend/internal/entities"
	"shorthack_backend/internal/service"
)

type PublicHandler struct {
	service service.StudentServ
}

func InitPublicHandler(service service.StudentServ) PublicHandler {
	return PublicHandler{
		service: service,
	}
}

// @Summary Create user
// @Tags public
// @Accept  json
// @Produce  json
// @Param data body entities.CreateStudent true "User create"
// @Success 200 {object} int "Successfully created user, returning JWT and Session"
// @Failure 400 {object} map[string]string "Invalid input"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /student/create [post]
func (p PublicHandler) CreateUser(c *gin.Context) {
	var userCreate entities.CreateStudent

	if err := c.ShouldBindJSON(&userCreate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx := c.Request.Context()

	id, err := p.service.Create(ctx, userCreate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": id})
}
