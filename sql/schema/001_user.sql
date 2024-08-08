-- +goose Up
CREATE TABLE 
"user" (
  id VARCHAR(64) PRIMARY KEY,
  created_at TIMESTAMP NOT NULL,
  name VARCHAR(128) NOT NULL,
  email VARCHAR(256) NOT NULL,
  password TEXT NOT NULL
);
-- +goose Down
DROP TABLE "user";

