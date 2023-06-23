-- +goose Up
-- +goose StatementBegin
CREATE PROCEDURE `RegisterUser`(
    IN inName VARCHAR(50),
    IN inEmail VARCHAR(140),
    IN inHashedPassword TEXT,
    IN inPhone VARCHAR(20),
    IN inCreatedAt bigint,
    IN inUpdatedAt bigint
)
BEGIN
    IF (SELECT COUNT(*) FROM users WHERE email = inEmail) OR (SELECT COUNT(*) FROM users WHERE phone = inPhone) > 0 THEN
        SELECT 0;
    ELSE
        INSERT INTO users (name, email, password, phone, created_at, updated_at)
        VALUES (inName, inEmail, inHashedPassword, inPhone, inCreatedAt, inUpdatedAt);
        SELECT 1;
    END IF;
END


-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP PROCEDURE IF EXISTS `RegisterUser`;
-- +goose StatementEnd
