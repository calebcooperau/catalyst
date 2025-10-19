package user

import (
	"context"
	"database/sql"

	"catalyst.api/internal/domain/user/data"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "golang.org/x/crypto/bcrypt"
)

type UserRepository interface {
	FindUserByID(ctx context.Context, id uuid.UUID) (*User, error)
	RegisterUser(ctx context.Context, cmp *User) (uuid.UUID, error)
	UpdateUser(ctx context.Context, cmp *User) (*User, error)
	DeleteUser(ctx context.Context, cmp *User) error
}

type UserSqlRepository struct {
	queries *data.Queries
	db      *pgxpool.Pool
}

func NewUserSqlRepository(db *pgxpool.Pool) *UserSqlRepository {
	queries := data.New(db)
	return &UserSqlRepository{
		queries: queries,
		db:      db,
	}
}

func (repository *UserSqlRepository) FindUserByID(ctx context.Context, id uuid.UUID) (*User, error) {
	userData, err := repository.queries.FindUserByID(ctx, id)

	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	user := &User{
		ID:           userData.ID,
		Email:        userData.Email,
		FirstName:    userData.FirstName,
		LastName:     userData.LastName,
		MobileNumber: *userData.MobileNumber,
		CreatedAt:    userData.CreatedAt.Time,
		UpdatedAt:    userData.UpdatedAt.Time,
	}

	return user, nil
}

func (repository *UserSqlRepository) RegisterUser(ctx context.Context, user *User) (uuid.UUID, error) {
	addUserParams := data.AddUserParams{
		Email:        user.Email,
		FirstName:    user.FirstName,
		LastName:     user.LastName,
		MobileNumber: &user.MobileNumber,
	}
	userResult, err := repository.queries.AddUser(ctx, addUserParams)
	if err != nil {
		return uuid.Nil, err
	}
	return userResult.ID, err
}

func (repository *UserSqlRepository) UpdateUser(ctx context.Context, user *User) (*User, error) {
	updateUserParams := data.UpdateUserParams{
		Email:        user.Email,
		FirstName:    user.FirstName,
		LastName:     user.LastName,
		MobileNumber: &user.MobileNumber,
		ID:           user.ID,
	}

	result, err := repository.queries.UpdateUser(ctx, updateUserParams)
	if err != nil {
		return nil, err
	}

	rowsAffected := result.RowsAffected()
	if rowsAffected == 0 {
		return nil, sql.ErrNoRows
	}
	return user, nil
}

func (repository *UserSqlRepository) DeleteUser(ctx context.Context, user *User) error {
	result, err := repository.queries.DeleteUser(ctx, user.ID)
	if err != nil {
		return err
	}
	rowsAffected := result.RowsAffected()
	if rowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}
