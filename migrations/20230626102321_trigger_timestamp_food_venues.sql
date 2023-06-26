-- +goose Up
-- +goose StatementBegin
CREATE TRIGGER IF NOT EXISTS setTimestampsFoodVenues BEFORE INSERT ON food_venues
    FOR EACH ROW
SET NEW.created_at = UNIX_TIMESTAMP(NOW()),
    NEW.updated_at = UNIX_TIMESTAMP(NOW());
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TRIGGER setTimestamps;
-- +goose StatementEnd
