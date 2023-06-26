-- +goose Up
-- +goose StatementBegin
CREATE PROCEDURE `EditProduct`(
    IN inProductId int,
    IN inProductName VARCHAR(50),
    IN inProductDescription VARCHAR(255),
    IN inProductPrice FLOAT
)
BEGIN

    IF (SELECT id FROM products WHERE id = inProductId) IS NULL THEN
        SELECT 0;
    ELSE
        UPDATE products
        SET name = inProductName, description = inProductDescription, price = inProductPrice
        WHERE id = inProductId;
        SELECT 1;
    END IF;
END
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP PROCEDURE IF EXISTS `EditProduct`;
-- +goose StatementEnd
