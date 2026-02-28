package handlers

import (
	"net/http"

	"github.com/JuanSposada/api_golang/cmd/internal/executor"
	"github.com/gin-gonic/gin"
)

type CommandRequest struct {
	Command string   `json:"command" binding:"required"`
	Params  []string `json:"params"`
}

func ExecuteCommandHandler(c *gin.Context) {
	var req CommandRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "JSON inválido"})
		return
	}

	// Obtenemos el usuario que se autenticó
	username, _, _ := c.Request.BasicAuth()

	// LLAMADA A LA FUNCIÓN (con el nombre que elegiste)
	output, err := executor.ExecuteAsUser(req.Command, req.Params, username)

	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{
			"status":  "error",
			"message": err.Error(),
			"output":  output,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"output": output,
	})
}
