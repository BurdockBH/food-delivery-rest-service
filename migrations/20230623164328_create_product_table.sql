-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS products (
                          id INT PRIMARY KEY AUTO_INCREMENT,
                          name VARCHAR(255) NOT NULL,
                          description VARCHAR(255) NOT NULL,
                          price DECIMAL(10, 2) NOT NULL,
                          food_venue_id INT,
                          created_at BIGINT NOT NULL,
                          updated_at BIGINT NOT NULL,
                          FOREIGN KEY (food_venue_id) REFERENCES food_venues(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS products;
-- +goose StatementEnd
