package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/msteinert/pam"
)

// Handler que devuelve siempre la contraseña
type passwordHandler struct {
	password string
}

func (h *passwordHandler) RespondPAM(style pam.Style, msg string) (string, error) {
	switch style {
	case pam.PromptEchoOff, pam.PromptEchoOn:
		// Solo enviamos la contraseña si PAM realmente la está pidiendo
		return h.password, nil
	default:
		// Para mensajes informativos o de error, devolvemos vacío sin romper la sesión
		return "", nil
	}
}
func CheckLinuxUser(username, password string) bool {
	handler := &passwordHandler{password: password}
	t, err := pam.Start("login", username, handler)
	if err != nil {
		return false
	}

	err = t.Authenticate(0)
	return err == nil
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		user, password, hasAuth := c.Request.BasicAuth()
		if !hasAuth || !CheckLinuxUser(user, password) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "usuario o contraseña de linux invalido"})
			c.Abort()
			return
		}
		c.Next()
	}
}
