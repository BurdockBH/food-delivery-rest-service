-- +goose Up
-- +goose StatementBegin
CREATE PROCEDURE `GetUsersByDetails`(
    IN inName VARCHAR(50),
    IN inEmail VARCHAR(50),
    IN inPhone VARCHAR(50)
)
BEGIN
    IF inName = '' AND inPhone = '' AND inEmail = '' THEN
        SELECT id, name, email, password, phone, created_at, updated_at FROM users;
    ELSE
        SELECT id, name, email, password, phone, created_at, updated_at
        FROM users
        WHERE (inName = '' OR name = inName)
          AND (inPhone = '' OR phone = inPhone)
          AND (inEmail = '' OR email = inEmail);
    END IF;
END
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP PROCEDURE IF EXISTS `GetUsersByDetails`;
-- +goose StatementEnd
