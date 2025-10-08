package handler

import (
	"net/http"
	"pankreatitmed/internal/app/dto/request"

	"github.com/gin-gonic/gin"
)

func (h *Handler) MedUserRegistation(c *gin.Context) {
	var user request.MedUserRegistration
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.svcs.MedUsers.Registrate(user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusCreated)
}

func (h *Handler) MedUserGetFields(c *gin.Context) {
	res, err := h.svcs.MedUsers.GetMyField()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, res)
}

func (h *Handler) MedUserUpdateFields(c *gin.Context) {
	var user request.UpdateMedUser
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.svcs.MedUsers.UpdateMyField(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusCreated)
}

func (h *Handler) MedUserLogIn(c *gin.Context) {
	var acces request.AuthenticateUser
	if err := c.ShouldBindJSON(&acces); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	c.Status(http.StatusOK)
	c.JSON(http.StatusOK, gin.H{
		"status":  "ok",
		"message": "login success",
		"user_id": 1,
		"token":   "fake-token-123",
	})
}

func (h *Handler) MedUserLogOut(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "ok",
		"message": "logout success",
	})
}
