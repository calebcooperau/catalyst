-- name: FindUserByID :one
SELECT id, email, first_name, last_name, mobile_number, created_at, updated_at
FROM users
WHERE id = $1;

-- name: AddUser :one
INSERT INTO users (email, first_name, last_name, mobile_number)	
VALUES ($1, $2, $3, $4)
RETURNING id, created_at, updated_at;

-- name: UpdateUser :execresult
UPDATE users 
SET email = $1, first_name = $2, last_name = $3, mobile_number = $4, updated_at = CURRENT_TIMESTAMP
WHERE id = $5
RETURNING updated_at;

-- name: DeleteUser :execresult
DELETE from users WHERE id = $1;
