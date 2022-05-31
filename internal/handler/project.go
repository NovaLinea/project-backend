package handler

import (
	"net/http"

	"github.com/ProjectUnion/project-backend.git/internal/domain"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (h *Handler) GetPopular(c *gin.Context) {
	projects, err := h.services.GetPopular(c)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, projects)
}

func (h *Handler) GetDataProject(c *gin.Context) {
	projectID, err := primitive.ObjectIDFromHex(c.Param("ID"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	data, err := h.services.GetDataProject(c, projectID)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, data)
}

func (h *Handler) Create(c *gin.Context) {
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
	inp.CreatorID = userID

	userData, err := h.services.GetDataProfile(c, userID)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	inp.CreatorName = userData.Name
	inp.CreatorPhoto = userData.Photo

	if err := h.services.Create(c, inp); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
}

func (h *Handler) GetProjectsUser(c *gin.Context) {
	userID, err := primitive.ObjectIDFromHex(c.Param("ID"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	projects, err := h.services.GetProjectsUser(c, userID)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, projects)
}

func (h *Handler) GetHome(c *gin.Context) {
	userID, err := primitive.ObjectIDFromHex(c.Param("ID"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	projects, err := h.services.GetHome(c, userID)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, projects)
}

func (h *Handler) GetFavorites(c *gin.Context) {
	userID, err := primitive.ObjectIDFromHex(c.Param("ID"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	projects, err := h.services.GetFavorites(c, userID)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, projects)
}

func (h *Handler) Like(c *gin.Context) {
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

	if err := h.services.Like(c, projectID, userID); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
}

func (h *Handler) Dislike(c *gin.Context) {
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

	if err := h.services.Dislike(c, projectID, userID); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
}

func (h *Handler) Favorite(c *gin.Context) {
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

	if err := h.services.Favorite(c, projectID, userID); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
}

func (h *Handler) RemoveFavorite(c *gin.Context) {
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

	if err := h.services.RemoveFavorite(c, projectID, userID); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
}

func (h *Handler) DeleteProject(c *gin.Context) {
	projectID, err := primitive.ObjectIDFromHex(c.Param("ID"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.DeleteProject(c, projectID); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
}

func (h *Handler) Update(c *gin.Context) {
	projectID, err := primitive.ObjectIDFromHex(c.Param("ID"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	var inp domain.ProjectEdit
	if err := c.BindJSON(&inp); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid input body")
		return
	}
	inp.ID = projectID

	if err := h.services.Update(c, inp); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
}
