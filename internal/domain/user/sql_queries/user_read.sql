-- name: GetUserDetailByID :one
SELECT id, email, first_name, last_name, mobile_number 
FROM users
WHERE id = $1;