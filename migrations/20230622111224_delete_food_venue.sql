-- +goose Up
-- +goose StatementBegin
CREATE PROCEDURE `DeleteFoodVenue`(
    IN inVenueId int
)
BEGIN

    IF (SELECT id FROM food_venues WHERE id = inVenueId) IS NOT NULL THEN
        DELETE FROM products WHERE food_venue_id = inVenueId;
        DELETE FROM food_venues WHERE id = inVenueId;
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
