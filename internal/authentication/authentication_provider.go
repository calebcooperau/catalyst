package authentication

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/markbates/goth/gothic"
)

type ProviderHandler struct {
	logger *log.Logger
}

func NewProviderHandler(logger *log.Logger) *ProviderHandler {
	return &ProviderHandler{
		logger: logger,
	}
}

// @Summary Start OAuth login
// @Description Begins the OAuth flow with the selected provider or redirects if already authenticated.
// @Tags auth
// @Param provider path string true "OAuth Provider"
// @Produce json
// @Success 302 {string} string "Redirect to provider login page or front-end"
// @Router /auth/{provider} [get]
func (handler ProviderHandler) GetProvider(ctx *gin.Context) {
	provider := ctx.Param("provider")

	newCtx := context.WithValue(ctx.Request.Context(), "provider", provider)
	ctx.Request = ctx.Request.WithContext(newCtx)
	_, err := gothic.CompleteUserAuth(ctx.Writer, ctx.Request)
	if err != nil {
		gothic.BeginAuthHandler(ctx.Writer, ctx.Request)
	} else {
		ctx.Redirect(http.StatusTemporaryRedirect, "http://localhost:4200")
	}
}
