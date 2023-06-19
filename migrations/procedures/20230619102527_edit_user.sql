-- +goose Up
-- +goose StatementBegin
DELIMITER //
CREATE DEFINER=`root`@`localhost` PROCEDURE `EditUser`(
    IN inName VARCHAR(50),
    IN inEmail VARCHAR(140),
    IN inHashedPassword TEXT,
    IN inPhone VARCHAR(20),
    IN inUpdatedAt bigint
)
BEGIN
    DECLARE userId INT;
    SELECT id INTO userId FROM users WHERE email = inEmail;
    IF userId IS NOT NULL THEN
        UPDATE users SET name = inName, password = inHashedPassword, phone = inPhone, updated_at = inUpdatedAt WHERE id = userID AND email = inEmail;
        SELECT 1;
    ELSE SELECT 0;
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
