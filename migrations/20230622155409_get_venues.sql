-- +goose Up
-- +goose StatementBegin
CREATE PROCEDURE `GetVenues`(
    IN inName VARCHAR(50),
    IN inAddress VARCHAR(50),
    IN inCreated_by_email VARCHAR(50)
)
BEGIN
    IF inName = '' AND inAddress = '' AND inCreated_by_email = '' THEN
        SELECT id, name, address, created_by, created_at, updated_at FROM food_venues;
    ELSE
        SELECT id, name, address, created_by, created_at, updated_at
        FROM food_venues
        WHERE (inName = '' OR name LIKE CONCAT('%', inName, '%'))
          AND (inAddress = '' OR address LIKE CONCAT('%', inAddress, '%'))
          AND (inCreated_by_email = '' OR created_by = inCreated_by_email);
    END IF;
END
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP PROCEDURE IF EXISTS `GetVenues`;
-- +goose StatementEnd
