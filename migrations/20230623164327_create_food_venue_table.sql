-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS food_venues (
                                           id INT PRIMARY KEY AUTO_INCREMENT,
                                           name VARCHAR(255) NOT NULL,
                                           address VARCHAR(255) NOT NULL,
                                           created_by VARCHAR(255) NOT NULL,
                                           created_at BIGINT NOT NULL,
                                           updated_at BIGINT NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS food_venues;
-- +goose StatementEnd
