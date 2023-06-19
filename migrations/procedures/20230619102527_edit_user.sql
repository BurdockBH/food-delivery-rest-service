-- +goose Up
-- +goose StatementBegin
DELIMITER //
CREATE DEFINER=`root`@`localhost` PROCEDURE `EditUser`(
    IN inName VARCHAR(50),
    IN oldEmail VARCHAR(140),
    IN inEmail VARCHAR(140),
    IN inHashedPassword TEXT,
    IN inPhone VARCHAR(20),
    IN inUpdatedAt BIGINT
)
BEGIN
    DECLARE userId INT;

    SELECT id INTO userId FROM users WHERE email = oldEmail;

    IF userId IS NOT NULL THEN
        UPDATE users
        SET name = inName,
            email = inEmail,
            password = inHashedPassword,
            phone = inPhone,
            updated_at = inUpdatedAt
        WHERE id = userId;

        SELECT 'EDITED' AS Message;
    END IF;
END //
DELIMITER ;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELIMITER //
DROP PROCEDURE IF EXISTS EditUser;
DELIMITER ;
-- +goose StatementEnd
