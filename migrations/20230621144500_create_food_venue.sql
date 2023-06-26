-- +goose Up
-- +goose StatementBegin
CREATE PROCEDURE `CreateFoodVenue`(
    IN inName VARCHAR(50),
    IN inAddress VARCHAR(50),
    IN inCreated_by VARCHAR(255)
)
BEGIN
    INSERT INTO food_venues (name, address, created_by)
    VALUES (inName, inAddress, inCreated_by);
END

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP PROCEDURE IF EXISTS `RegisterUser`;
-- +goose StatementEnd
a