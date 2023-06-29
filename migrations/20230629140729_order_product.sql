-- +goose Up
-- +goose StatementBegin
CREATE PROCEDURE `OrderProduct`(
    IN inProductId INT,
    IN inUserEmail VARCHAR(50),
    IN inOrderQuantity INT
)
BEGIN
    DECLARE totalPrice DECIMAL(10,2);
    DECLARE originalPrice DECIMAL(10,2);
    DECLARE userId INT;
    START TRANSACTION;

    IF (SELECT COUNT(*)FROM products WHERE id = inProductId) = 0 THEN
        ROLLBACK;
        SELECT -2;
    ELSE
        SELECT price INTO originalPrice FROM products WHERE id = inProductId;
        SET totalPrice = originalPrice * inOrderQuantity;

        SELECT id INTO userId FROM users WHERE email = inUserEmail;

        IF (SELECT quantity FROM products WHERE id = inProductId) >= inOrderQuantity THEN
            UPDATE products SET quantity = quantity - inOrderQuantity WHERE id = inProductId;
            INSERT INTO transactions (product_id, user_id, quantity, price, total_price) VALUES (inProductId, userId, inOrderQuantity, originalPrice, totalPrice);
            COMMIT;
            SELECT 1;
        ELSE
            ROLLBACK;
            SELECT -1;
        END IF;
    END IF;

END

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP PROCEDURE IF EXISTS `BuyProduct`;
-- +goose StatementEnd
