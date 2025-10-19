package user

import (
	"log"
	"net/http"

	"catalyst.api/internal/domain/user/data"
	"catalyst.api/internal/utilities"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UserDetailQuery struct {
	ID uuid.UUID
}

type UserDetailApiDto struct {
	ID           uuid.UUID
	Email        string
	FirstName    string
	LastName     string
	MobileNumber *string
}

type UserDetailHandler struct {
	queries *data.Queries
	logger  *log.Logger
}

func NewUserDetailHandler(queries *data.Queries, logger *log.Logger) *UserDetailHandler {
	return &UserDetailHandler{
		queries: queries,
		logger:  logger,
	}
}

// @Summary Get user details by ID
// @Description Retrieves a user's detailed profile by their unique ID.
// @Tags users
// @Param id path string true "User ID"
// @Produce json
// @Success 200 {object} map[string]interface{} "User object"
// @Failure 400 {object} map[string]string "Invalid ID"
// @Failure 404 {object} map[string]string "User not found"
// @Router /users/{id} [get]
func (handler UserDetailHandler) GetUserByID(ctx *gin.Context) {
	userID, err := utilities.ReadIDParam(ctx)
	if err != nil {
		handler.logger.Printf("ERROR: ReadIDParam: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid User ID"})
		return
	}

	query := UserDetailQuery{
		ID: userID,
	}

	user, err := handler.queries.GetUserDetailByID(ctx.Request.Context(), query.ID)
	if err != nil {
		handler.logger.Printf("ERROR: repositoryGetUserByID: %v", err)
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Not Found"})
		return
	}

	userApiDto := UserDetailApiDto{
		ID:           user.ID,
		Email:        user.Email,
		FirstName:    user.FirstName,
		LastName:     user.LastName,
		MobileNumber: user.MobileNumber,
	}

	ctx.JSON(http.StatusOK, gin.H{"User": userApiDto})
}
