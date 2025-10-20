package handler

import (
	"errors"
	"net/http"
	"pankreatitmed/internal/app/authctx"
	"pankreatitmed/internal/app/dto/request"
	"pankreatitmed/internal/app/dto/response"
	"pankreatitmed/internal/app/services"

	"github.com/gin-gonic/gin"
)

func (h *Handler) MedUserRegistation(c *gin.Context) {
	var user request.MedUserRegistration
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	_, token, err := h.svcs.MedUsers.Register(user)
	switch {
	case errors.Is(err, services.ErrWeakPassword):
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	case errors.Is(err, services.ErrLoginTaken):
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	case err != nil:
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	exp := h.svcs.MedUsers.GetConfig().TTL.Hours()
	c.JSON(http.StatusCreated, response.AuthorizateUser{AccessToken: token, TokenType: "Bearer", ExpiresIn: int(exp)})
}

func (h *Handler) MedUserGetFields(c *gin.Context) {
	user, check := authctx.Get(c)
	if !check {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "problem with your token"})
		return
	}
	res, err := h.svcs.MedUsers.GetMyField(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

func (h *Handler) MedUserUpdateFields(c *gin.Context) {
	usr, check := authctx.Get(c)
	if !check {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "problem with your token"})
		return
	}
	var user request.UpdateMedUser
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.svcs.MedUsers.UpdateField(usr.ID, &user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusCreated)
}

func (h *Handler) MedUserLogIn(c *gin.Context) {
	var acces request.AuthenticateMedUser
	if err := c.ShouldBindJSON(&acces); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	token, err := h.svcs.MedUsers.Login(acces)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		return
	}
	exp := h.svcs.MedUsers.GetConfig().TTL.Hours()
	c.JSON(http.StatusOK, response.AuthorizateUser{AccessToken: token, TokenType: "Bearer", ExpiresIn: int(exp)})
}

func (h *Handler) MedUserLogOut(c *gin.Context) {
	token := c.Param("token")
	if err := h.svcs.MedUsers.Logout(token); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  "ok",
		"message": "logout success",
	})
}
