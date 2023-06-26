-- +goose Up
-- +goose StatementBegin
CREATE PROCEDURE `DeleteFoodVenue`(
    IN inId INT,
    IN inName VARCHAR(50),
    IN inAddress VARCHAR(50)
)
BEGIN
    IF (inName IS NULL OR inName = '') AND (inAddress IS NULL OR inAddress = '') THEN
        DELETE FROM food_venues WHERE id = inId;
        SELECT 1;
    ELSE
        IF (SELECT COUNT(*) FROM food_venues WHERE name = inName AND address = inAddress) > 0 THEN
            DELETE FROM food_venues WHERE name = inName AND address = inAddress;
            SELECT 1;
        ELSE
            SELECT 0;
        END IF;
    END IF;
END

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP PROCEDURE IF EXISTS `RegisterUser`;
-- +goose StatementEnd
