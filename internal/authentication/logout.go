package authentication

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/markbates/goth/gothic"
)

type LogoutHandler struct {
	logger *log.Logger
}

func NewLogoutHandler(logger *log.Logger) *LogoutHandler {
	return &LogoutHandler{
		logger: logger,
	}
}

// @Summary Logout user
// @Description Logs out the currently authenticated user via the configured provider session.
// @Tags auth
// @Produce json
// @Success 302 {string} string "Redirect to /"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /auth/logout/{provider} [get]
func (handler LogoutHandler) Logout(ctx *gin.Context) {
	err := gothic.Logout(ctx.Writer, ctx.Request)
	if err != nil {
		handler.logger.Printf("ERROR: gothicLogout: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}
	ctx.Redirect(http.StatusTemporaryRedirect, "/")
}
