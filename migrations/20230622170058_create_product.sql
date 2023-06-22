-- +goose Up
-- +goose StatementBegin
CREATE PROCEDURE `CreateProduct`(
    IN inProductName VARCHAR(50),
    IN inProductDescription VARCHAR(255),
    IN inProductPrice FLOAT,
    IN inFoodVenueName VARCHAR(50),
    IN inFoodVenueAddress VARCHAR(50),
    IN createdAt bigint,
    IN updatedAt bigint
)
BEGIN
    DECLARE venueId int;

    SELECT id INTO venueId FROM food_venues WHERE name = inFoodVenueName AND address = inFoodVenueAddress;

    IF venueId IS NULL THEN
        SELECT -1;
    ELSE
        IF (SELECT id FROM products WHERE Name = name AND food_venue_id = venueId) IS NULL THEN
            INSERT INTO products (name, description, price, food_venue_id, created_at, updated_at)
            VALUES (inProductName, inProductDescription, inProductPrice, venueId, createdAt, updatedAt);

            SELECT 1;
        ELSE
            SELECT 0;
        END IF;
    END IF;
END
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP PROCEDURE IF EXISTS `CreateProduct`;
-- +goose StatementEnd
