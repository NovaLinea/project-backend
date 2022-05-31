package handler

import (
	"errors"
	"net/http"

	"github.com/ProjectUnion/project-backend.git/internal/domain"
	"github.com/gin-gonic/gin"
)

func (h *Handler) Register(c *gin.Context) {
	var inp domain.UserAuth
	if err := c.BindJSON(&inp); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid input body")
		return
	}

	res, err := h.services.Authorization.Register(c, inp)
	if err != nil {
		if errors.Is(err, domain.ErrUserNotFound) {
			newErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		}
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	h.logger.Infof("Register user %s", res.UserID)

	c.JSON(http.StatusOK, map[string]interface{}{
		"accessToken": res.AccessToken,
		"id":          res.UserID,
	})
}

func (h *Handler) Login(c *gin.Context) {
	var inp domain.UserAuth
	if err := c.BindJSON(&inp); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid input body")
		return
	}

	res, err := h.services.Authorization.Login(c, inp.Email, inp.Password)
	if err != nil {
		if errors.Is(err, domain.ErrUserNotFound) {
			newErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		}
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	h.logger.Infof("Login user %s", res.UserID)

	c.JSON(http.StatusOK, map[string]interface{}{
		"accessToken": res.AccessToken,
		"id":          res.UserID,
	})
}
