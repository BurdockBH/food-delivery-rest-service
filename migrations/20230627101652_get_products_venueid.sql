-- +goose Up
-- +goose StatementBegin
CREATE PROCEDURE `GetProducts`(
    IN inVenueId int
)
BEGIN

    SELECT p.id, p.name, p.description, p.price, p.food_venue_id, p.created_by, p.created_at, p.updated_at, v.name AS food_venue_name, v.address AS food_venue_address
    FROM products p
    LEFT JOIN food_venues v ON p.food_venue_id = v.id
    WHERE food_venue_id = inVenueId;
END
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP PROCEDURE IF EXISTS `GetProducts`;
-- +goose StatementEnd
