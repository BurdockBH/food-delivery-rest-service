-- +goose Up
-- +goose StatementBegin
CREATE TRIGGER updateTimestamps BEFORE UPDATE ON users
    FOR EACH ROW
    SET NEW.updated_at = UNIX_TIMESTAMP(NOW());
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TRIGGER updateTimestamps;
-- +goose StatementEnd
