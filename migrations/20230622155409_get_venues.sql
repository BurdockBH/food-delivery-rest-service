-- +goose Up
-- +goose StatementBegin
CREATE PROCEDURE `GetVenues`(
    IN inName VARCHAR(50),
    IN inAddress VARCHAR(50),
    IN inCreated_by_email VARCHAR(50)
)
BEGIN
    SELECT id, name, address, created_by, created_at, updated_at
    FROM food_venues
    WHERE (LENGTH(inName) = 0 OR name LIKE CONCAT('%', inName, '%'))
        AND (LENGTH(inAddress) = 0 OR address LIKE CONCAT('%', inAddress, '%'))
        AND (LENGTH(inCreated_by_email) = 0 OR created_by = inCreated_by_email)
       OR (LENGTH(inName) = 0 AND LENGTH(inAddress) = 0 AND LENGTH(inCreated_by_email) = 0);
END
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP PROCEDURE IF EXISTS `GetVenues`;
-- +goose StatementEnd
