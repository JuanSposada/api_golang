package handlers

import (
	"net/http"

	"github.com/JuanSposada/api_golang/cmd/internal/executor" // Importamos tu ejecutor
	"github.com/gin-gonic/gin"
)

// Estructura para el body de la petición
type CommandRequest struct {
	Command string   `json:"command" binding:"required"`
	Params  []string `json:"params"`
}

// ExecuteCommandHandler es la función que "maneja" la ruta POST /system/execute
func ExecuteCommandHandler(c *gin.Context) {
	var req CommandRequest

	// 1. Validar el JSON de entrada
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "JSON inválido o campos faltantes"})
		return
	}

	// 2. Llamar a la lógica de ejecución segura (lo que hicimos en executor.go)
	output, err := executor.Execute(req.Command, req.Params)
	if err != nil {
		// Si el comando no está en la whitelist o falla
		c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		return
	}

	// 3. Responder con éxito
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"output": output,
	})
}
