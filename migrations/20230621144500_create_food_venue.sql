-- +goose Up
-- +goose StatementBegin
CREATE PROCEDURE `CreateFoodVenue`(
    IN inName VARCHAR(50),
    IN inAddress VARCHAR(50),
    IN inCreatedAt bigint,
    IN inUpdatedAt bigint
)
BEGIN

    IF (SELECT COUNT(*) FROM food_venues WHERE name = inName AND address = inAddress) > 0 THEN
        SELECT 0;
    ELSE
        INSERT INTO food_venues (name, address, created_at, updated_at)
        VALUES (inName, inAddress, inCreatedAt, inUpdatedAt);
        SELECT 1;
    END IF;
END

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP PROCEDURE IF EXISTS `RegisterUser`;
-- +goose StatementEnd
