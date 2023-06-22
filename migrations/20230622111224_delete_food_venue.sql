-- +goose Up
-- +goose StatementBegin
CREATE PROCEDURE `DeleteFoodVenue`(
    IN inName VARCHAR(50),
    IN inAddress VARCHAR(50)
)
BEGIN

    IF (SELECT COUNT(*) FROM food_venues WHERE name = inName AND address = inAddress) > 0 THEN
        DELETE FROM food_venues WHERE name = inName AND address = inAddress;
        SELECT 1;
    ELSE
        SELECT 0;
    END IF;
END

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP PROCEDURE IF EXISTS `RegisterUser`;
-- +goose StatementEnd
