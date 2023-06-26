-- +goose Up
-- +goose StatementBegin
CREATE PROCEDURE `DeleteProduct`(
    IN inProductId int
)
BEGIN

    IF (SELECT id FROM products WHERE id = inProductId) IS NOT NULL THEN
        DELETE FROM products WHERE id = inProductId;
        SELECT 1;
    ELSE
        SELECT 0;
    END IF;
END
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP PROCEDURE IF EXISTS `DeleteProduct`;
-- +goose StatementEnd
