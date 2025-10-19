-- name: FindUserIDByProvider :one
SELECT user_id 
FROM auth_user_providers 
WHERE provider_user_id = $1;

-- name: FindAuthUserByID :one
SELECT id, email, first_name, last_name, mobile_number, created_at, updated_at 
FROM users
WHERE id = $1;

-- name: CreateAuthUser :one
WITH new_auth_user AS (
    INSERT INTO auth_users (email, first_name, last_name)
    VALUES ($1, $2, $3)
    RETURNING id
),
new_provider AS (
    INSERT INTO auth_user_providers (user_id, provider, provider_user_id)
    SELECT id, $4, $5 FROM new_auth_user
),
new_user AS (
    INSERT INTO users (id, email, first_name, last_name)
    SELECT id, $1, $2, $3 FROM new_auth_user
)
SELECT id FROM new_auth_user;