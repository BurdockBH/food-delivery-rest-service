-- +goose Up
-- +goose StatementBegin
CREATE DEFINER=`root`@`localhost` PROCEDURE `RegisterUser`(
    IN inName VARCHAR(50),
    IN inEmail VARCHAR(140),
    IN inHashedPassword TEXT,
    IN inPhone VARCHAR(20),
    IN inCreatedAt bigint,
    IN inUpdatedAt bigint
)
BEGIN
    DECLARE existingCount INT;

    SELECT COUNT(*) INTO existingCount FROM users WHERE email = inEmail;

    IF existingCount > 0 THEN
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
