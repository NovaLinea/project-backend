package handler

import (
	"net/http"

	"github.com/ProjectUnion/project-backend.git/internal/domain"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (h *Handler) CreateProject(c *gin.Context) {
	userID, err := primitive.ObjectIDFromHex(c.Param("ID"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	var inp domain.ProjectData
	if err := c.BindJSON(&inp); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid input body")
		return
	}
	inp.UserID = userID

	if err := h.services.CreateProject(c, inp); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
}
