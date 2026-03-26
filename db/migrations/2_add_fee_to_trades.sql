-- +goose Up
-- +goose StatementBegin
ALTER TABLE trades 
ADD COLUMN fee DECIMAL(18, 8) NOT NULL DEFAULT 0;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE trades 
DROP COLUMN fee;
-- +goose StatementEnd