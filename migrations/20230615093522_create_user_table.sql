-- +goose Up
-- +goose StatementBegin
CREATE TABLE users (
                       id INT PRIMARY KEY AUTO_INCREMENT,
                       name VARCHAR(255) NOT NULL,
                       email VARCHAR(255) NOT NULL,
                       password VARCHAR(255) NOT NULL,
                       phone VARCHAR(255) NOT NULL,
                       created_at BIGINT NOT NULL,
                       updated_at BIGINT NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
-- +goose StatementEnd
