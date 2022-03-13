package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// SendError godoc
func SendError(c *gin.Context, code int, message string) {
	ResponseJSONMessage(c, code, message, false, nil)
}

// SendErrValidator godoc
func SendErrValidator(c *gin.Context, message string) {
	ResponseJSONMessage(c, http.StatusCreated, message, false, nil)
}

// SendBadRequets godoc
func SendBadRequets(c *gin.Context, message string) {
	ResponseJSONMessage(c, http.StatusBadRequest, message, false, nil)
}

// SendNotFound godoc
func SendNotFound(c *gin.Context, message string) {
	ResponseJSONMessage(c, http.StatusNotFound, message, false, nil)
}

// SendInternalServerError godoc
func SendInternalServerError(c *gin.Context, message string) {
	ResponseJSONMessage(c, http.StatusInternalServerError, message, false, nil)
}

// SendSuccess godoc
func SendSuccess(c *gin.Context, data interface{}) {
	ResponseJSONMessage(c, http.StatusOK, "success", true, data)
}

// ResponseJSONMessage godoc
func ResponseJSONMessage(c *gin.Context, code int, message string, status bool, payload interface{}) {
	c.JSON(code, map[string]interface{}{
		"status":  status,
		"message": message,
		"data":    payload,
	})
	return
}
