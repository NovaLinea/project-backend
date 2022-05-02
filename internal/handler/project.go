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

func (h *Handler) GetProjects(c *gin.Context) {
	userID, err := primitive.ObjectIDFromHex(c.Param("ID"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	projects, err := h.services.GetProjects(c, userID)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, projects)
}

func (h *Handler) GetFavoritesProjects(c *gin.Context) {
	userID, err := primitive.ObjectIDFromHex(c.Param("ID"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	projects, err := h.services.GetFavoritesProjects(c, userID)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, projects)
}

func (h *Handler) LikeProject(c *gin.Context) {
	projectID, err := primitive.ObjectIDFromHex(c.Param("ID"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	userID, err := primitive.ObjectIDFromHex(c.Param("userID"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.LikeProject(c, projectID, userID); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
}

func (h *Handler) DislikeProject(c *gin.Context) {
	projectID, err := primitive.ObjectIDFromHex(c.Param("ID"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	userID, err := primitive.ObjectIDFromHex(c.Param("userID"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.DislikeProject(c, projectID, userID); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
}

func (h *Handler) FavoriteProject(c *gin.Context) {
	projectID, err := primitive.ObjectIDFromHex(c.Param("ID"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	userID, err := primitive.ObjectIDFromHex(c.Param("userID"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.FavoriteProject(c, projectID, userID); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
}

func (h *Handler) RemoveFavoriteProject(c *gin.Context) {
	projectID, err := primitive.ObjectIDFromHex(c.Param("ID"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	userID, err := primitive.ObjectIDFromHex(c.Param("userID"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.RemoveFavoriteProject(c, projectID, userID); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
}
