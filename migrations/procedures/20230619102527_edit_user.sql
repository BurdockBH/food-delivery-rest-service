-- +goose Up
-- +goose StatementBegin
CREATE PROCEDURE `EditUser`(
    IN inName VARCHAR(50),
    IN inEmail VARCHAR(140),
    IN inHashedPassword TEXT,
    IN inPhone VARCHAR(20),
    IN inUpdatedAt bigint
)
BEGIN
    DECLARE userId INT;
    SELECT id INTO userId FROM users WHERE email = inEmail;

    IF userId IS NULL THEN
        SELECT -1;
    ELSE
        IF (SELECT COUNT(*) FROM users WHERE phone = inPhone AND id <> userId) > 0 THEN
            SELECT -2;
        ELSE
            UPDATE users SET name = inName, password = inHashedPassword, phone = inPhone, updated_at = inUpdatedAt WHERE id = userId;
            SELECT 1;
        END IF;
    END IF;
END
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP PROCEDURE IF EXISTS EditUser;
-- +goose StatementEnd
