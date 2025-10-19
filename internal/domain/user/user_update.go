package user

import (
	"encoding/json"
	"log"
	"net/http"

	"catalyst.api/internal/utilities"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"github.com/google/uuid"
)

type UserUpdateCommand struct {
	ID           uuid.UUID
	Email        string
	FirstName    string
	LastName     string
	MobileNumber string
}

type UserUpdateApiDto struct {
	Email        string `json:"email" validate:"required"`
	FirstName    string `json:"firstName" validate:"required"`
	LastName     string `json:"lastName" validate:"required"`
	MobileNumber string `json:"mobileNumber" `
}

func (dto *UserUpdateApiDto) ValidateApiDto() error {
	validate := validator.New()
	return validate.Struct(dto)
}

type UserUpdateHandler struct {
	repository UserRepository
	logger     *log.Logger
}

func NewUserUpdateHandler(repository UserRepository, logger *log.Logger) *UserUpdateHandler {
	return &UserUpdateHandler{
		repository: repository,
		logger:     logger,
	}
}

// @Summary Update user information by ID
// @Description Updates the email, name, and mobile number for a given user.
// @Tags users
// @Param id path string true "User ID"
// @Accept json
// @Produce json
// @Param user body UserUpdateApiDto true "User update payload"
// @Success 200 {object} map[string]interface{} "Updated user object"
// @Failure 400 {object} map[string]string "Invalid input"
// @Failure 404 {object} map[string]string "User not found"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /users/{id} [put]
func (handler UserUpdateHandler) UpdateUser(ctx *gin.Context) {
	userID, err := utilities.ReadIDParam(ctx)
	if err != nil {
		handler.logger.Printf("ERROR: readIDParam: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Request Sent"})
		return
	}

	// build command
	var userUpdateApiDto UserUpdateApiDto
	err = json.NewDecoder(ctx.Request.Body).Decode(&userUpdateApiDto)
	if err != nil {
		handler.logger.Printf("Error: decodeUserUpdateApiDto: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Request Sent"})
		return
	}

	// validate
	err = userUpdateApiDto.ValidateApiDto()
	if err != nil {
		handler.logger.Printf("ERROR: validateUserUpdateApiDto: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Request Sent"})
		return
	}

	command := UserUpdateCommand{
		ID:           userID,
		Email:        userUpdateApiDto.Email,
		FirstName:    userUpdateApiDto.FirstName,
		LastName:     userUpdateApiDto.LastName,
		MobileNumber: userUpdateApiDto.MobileNumber,
	}

	user, err := handler.repository.FindUserByID(ctx.Request.Context(), command.ID)
	if err != nil {
		handler.logger.Printf("Error: repositoryGetUserByID: %v", err)
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Not Found"})
		return
	}

	user, err = user.Update(command.Email, command.FirstName, command.LastName, command.MobileNumber)
	if err != nil {
		handler.logger.Printf("Error: modelUserUpdate: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	user, err = handler.repository.UpdateUser(ctx.Request.Context(), user)
	if err != nil {
		handler.logger.Printf("Error: repositoryUpdateUser: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"User": user})
}
