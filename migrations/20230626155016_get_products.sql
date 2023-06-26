-- +goose Up
-- +goose StatementBegin
CREATE PROCEDURE `GetProducts`(
    IN product_name VARCHAR(255),
    IN venue_name VARCHAR(255),
    IN venue_address VARCHAR(255)
)
BEGIN
    DECLARE venue_id int;

    IF venue_name <> '' AND venue_address <> '' THEN
        SELECT id INTO venue_id
        FROM food_venues
        WHERE name = venue_name AND address = venue_address;
    ELSEIF venue_name <> '' THEN
        SELECT id INTO venue_id
        FROM food_venues
        WHERE name = venue_name;
    END IF;

    IF product_name <> '' THEN
        -- Case 1: Product name is provided, venue name and address are empty
        IF venue_name = '' AND venue_address = '' THEN
            SELECT p.id, p.name, p.description, p.price, p.food_venue_id, p.created_by, p.created_at, p.updated_at, v.name AS food_venue_name, v.address AS food_venue_address
            FROM products p
                     LEFT JOIN food_venues v ON p.food_venue_id = v.id
            WHERE p.name LIKE CONCAT('%', product_name, '%');

            -- Case 2: Venue name and address are provided, product name is empty
        ELSEIF venue_name <> '' AND venue_address <> '' THEN
            SELECT p.id, p.name, p.description, p.price, p.food_venue_id, p.created_by, p.created_at, p.updated_at,
                   v.name AS food_venue_name, v.address AS food_venue_address
            FROM products p
                     LEFT JOIN food_venues v ON p.food_venue_id = v.id
            WHERE p.food_venue_id = venue_id;

            -- Case 3: Only venue name is provided, product name and address are empty
        ELSE
            SELECT p.id, p.name, p.description, p.price, p.food_venue_id, p.created_by, p.created_at, p.updated_at, v.name AS food_venue_name, v.address AS food_venue_address
            FROM products p
                     LEFT JOIN food_venues v ON p.food_venue_id = v.id
            WHERE p.food_venue_id IN (SELECT id FROM food_venues WHERE name = venue_name);
        END IF;

        -- Case 4: All parameters are empty
    ELSE
        SELECT p.id, p.name, p.description, p.price, p.food_venue_id, p.created_by, p.created_at, p.updated_at, v.name AS food_venue_name, v.address AS food_venue_address
        FROM products p
                 LEFT JOIN food_venues v ON p.food_venue_id = v.id;
    END IF;
END
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP PROCEDURE IF EXISTS `GetProducts`;
-- +goose StatementEnd
