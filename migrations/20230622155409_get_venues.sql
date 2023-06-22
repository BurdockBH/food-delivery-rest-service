-- +goose Up
-- +goose StatementBegin
CREATE PROCEDURE `GetVenues`(
    IN inName VARCHAR(50),
    IN inAddress VARCHAR(50)
)
BEGIN
    IF inName = '' AND inAddress = '' THEN
        SELECT id, name, address, created_at, updated_at FROM food_venues;
    ELSE
        SELECT id, name, address, created_at, updated_at
        FROM food_venues
        WHERE (inName = '' OR name = inName)
          AND (inAddress = '' OR address = inAddress);
    END IF;
END
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP PROCEDURE IF EXISTS `GetVenues`;
-- +goose StatementEnd
