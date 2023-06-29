-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS transactions (
                                     id INT PRIMARY KEY AUTO_INCREMENT,
                                     product_id INT NOT NULL,
                                     user_id INT NOT NULL,
                                     quantity INT,
                                     price DECIMAL(10,2),
                                     total_price DECIMAL(10,2),
                                    FOREIGN KEY (product_id) REFERENCES products(id),
                                    FOREIGN KEY (user_id) REFERENCES users(id)
                                    );
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE transactions;
-- +goose StatementEnd
