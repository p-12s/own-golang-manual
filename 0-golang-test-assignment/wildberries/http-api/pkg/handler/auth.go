package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/p-12s/own-golang-manual/0-golang-test-assignment/wildberries/http-api"
	"net/http"
)

func (h *Handler) test(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]interface{}{
		"hello": 321,
	})
}

func (h *Handler) signUp(c *gin.Context) {
	// в задаче не требуется делать регистр.-авторизац.-аутентиф., но заготовка пусть будет
	var input common.User

	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.CreateUser(input)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) signIn(c *gin.Context) {

}
