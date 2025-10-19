-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS auth_user_providers (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(), 
  user_id UUID NOT NULL REFERENCES auth_users(id),
  provider VARCHAR(255) NOT NULL,
  provider_user_id VARCHAR(255) NOT NULL,
  created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE auth_user_providers;
-- +goose StatementEnd