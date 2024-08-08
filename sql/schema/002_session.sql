-- +goose Up
CREATE TABLE IF NOT EXISTS session(
  id VARCHAR(64) PRIMARY KEY,
  expires_at TIMESTAMP,
  user_id VARCHAR(64) REFERENCES "user"(id) NOT NULL 
);

-- +goose Down
DROP TABLE session;
