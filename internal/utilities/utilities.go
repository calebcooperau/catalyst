package utilities

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func ReadIDParam(ctx *gin.Context) (uuid.UUID, error) {
	idParam := ctx.Param("id")
	if idParam == "" {
		http.NotFound(ctx.Writer, ctx.Request)
		return uuid.Nil, errors.New("invalid id parameter")
	}

	id, err := uuid.Parse(idParam)
	if err != nil {
		http.NotFound(ctx.Writer, ctx.Request)
		return uuid.Nil, errors.New("invalid id parameter type")
	}
	return id, nil
}
