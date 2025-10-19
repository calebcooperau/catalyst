package authentication

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
)

type SignInHandler struct {
	repository AuthenticationRepository
	logger     *log.Logger
}

func NewSignInHandler(authenticationRepo AuthenticationRepository, logger *log.Logger) *SignInHandler {
	return &SignInHandler{
		repository: authenticationRepo,
		logger:     logger,
	}
}

// @Summary OAuth callback
// @Description Handle the callback from the OAuth provider and generate a JWT token.
// @Tags auth
// @Accept json
// @Produce json
// @Param provider path string true "OAuth Provider"
// @Success 302 {string} string "Redirect"
// @Failure 500 {object} map[string]string
// @Router /auth/{provider}/callback [get]
func (handler *SignInHandler) SignInCallback(ctx *gin.Context) {
	provider := ctx.Param("provider")
	newCtx := context.WithValue(ctx.Request.Context(), "provider", provider)
	ctx.Request = ctx.Request.WithContext(newCtx)
	gothUser, err := gothic.CompleteUserAuth(ctx.Writer, ctx.Request)
	if err != nil {
		fmt.Fprintln(ctx.Writer, err)
		return
	}

	authUserID, err := handler.repository.FindUserIDByProvider(ctx.Request.Context(), gothUser.UserID)
	if err == sql.ErrNoRows {
		// create new user
		authUserID, err = handler.registerAuthUser(ctx, gothUser)
		if err != nil {
			handler.logger.Printf("ERROR: handlerRegisterAuthUser: %v", err)
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
			return
		}
	}
	if err != nil {
		handler.logger.Printf("ERROR: repositoryFindUserIDByProvider: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	// found user, generate token etc
	token, err := GenerateJWT(authUserID, gothUser.Email, AuthScope, 2*time.Hour)
	if err != nil {
		handler.logger.Printf("ERROR: tokenGenerateJWT: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	redirectUrl := fmt.Sprintf("http://localhost:4200/callback?token=%s", token)
	ctx.Redirect(http.StatusTemporaryRedirect, redirectUrl)
}

func (handler *SignInHandler) registerAuthUser(ctx *gin.Context, gothUser goth.User) (uuid.UUID, error) {
	userID, err := handler.repository.RegisterAuthUser(ctx.Request.Context(), gothUser)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return uuid.Nil, err
	}
	return userID, nil
}
