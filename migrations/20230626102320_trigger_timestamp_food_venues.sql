-- +goose Up
-- +goose StatementBegin
CREATE TRIGGER IF NOT EXISTS setTimestampsFoodVenues BEFORE INSERT ON food_venues
    FOR EACH ROW
SET NEW.created_at = CURRENT_TIMESTAMP(),
    NEW.updated_at = CURRENT_TIMESTAMP();


-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TRIGGER setTimestamps;
-- +goose StatementEnd
