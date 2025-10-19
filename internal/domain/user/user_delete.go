package user

import (
	"database/sql"
	"log"
	"net/http"

	"catalyst.api/internal/utilities"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UserDeleteCommand struct {
	ID uuid.UUID
}

type UserDeleteHandler struct {
	repository UserRepository
	logger     *log.Logger
}

func NewUserDeleteHandler(repository UserRepository, logger *log.Logger) *UserDeleteHandler {
	return &UserDeleteHandler{
		repository: repository,
		logger:     logger,
	}
}

// @Summary Delete a user by ID
// @Description Deletes the specified user if they exist and can be deleted.
// @Tags users
// @Param id path string true "User ID"
// @Success 204 {string} string "No Content"
// @Failure 400 {object} map[string]string "Invalid ID"
// @Failure 404 {object} map[string]string "User not found"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /users/{id} [delete]
func (handler UserDeleteHandler) DeleteUser(ctx *gin.Context) {
	userID, err := utilities.ReadIDParam(ctx)
	if err != nil {
		handler.logger.Printf("ERROR: readIDParam: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid User ID"})
		return
	}

	UserDeleteCommand := UserDeleteCommand{
		ID: userID,
	}
	user, err := handler.repository.FindUserByID(ctx.Request.Context(), UserDeleteCommand.ID)
	if err != nil {
		handler.logger.Printf("ERROR: repositoryGetUserByID: %v", err)
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Not Found"})
		return
	}

	if user == nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Not Found"})
		return
	}
	err = user.CanDelete()
	if err != nil {
		handler.logger.Printf("ERROR: userCanDelete: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	err = handler.repository.DeleteUser(ctx.Request.Context(), user)
	if err == sql.ErrNoRows {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Not Found"})
		return
	}
	if err != nil {
		handler.logger.Printf("ERROR repositoryDeleteUser: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}
	ctx.Writer.WriteHeader(http.StatusNoContent)
}
