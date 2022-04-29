package handler

import (
	"fmt"
	"net/http"

	"github.com/ProjectUnion/project-backend.git/internal/domain"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (h *Handler) GetDataUser(c *gin.Context) {
	userID, err := primitive.ObjectIDFromHex(c.Param("userID"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	data, err := h.services.GetData(c, userID)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	fmt.Println(data)

	c.JSON(http.StatusOK, data)
}

func (h *Handler) SaveDataUser(c *gin.Context) {
	userID, err := primitive.ObjectIDFromHex(c.Param("userID"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	var inp domain.UserData
	if err := c.BindJSON(&inp); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid input body")
		return
	}
	fmt.Println(inp)

	if err := h.services.SaveData(c, userID, inp); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
}
