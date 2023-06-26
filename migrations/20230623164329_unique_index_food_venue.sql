-- +goose Up
-- +goose StatementBegin
CREATE UNIQUE INDEX venue_index ON food_venues(name, address);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE food_venues DROP INDEX venue_index;
-- +goose StatementEnd
