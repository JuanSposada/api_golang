package main

import (
	"github.com/JuanSposada/api_golang/cmd/internal/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Grupo de rutas protegidas
	systemGroup := r.Group("/system")

	// Aplicamos el middleware de usuario Linux
	systemGroup.Use(handlers.AuthMiddleware())
	{
		// Aquí usamos la función que definimos en system.go
		systemGroup.POST("/execute", handlers.ExecuteCommandHandler)
	}

	r.Run(":8080")
}
