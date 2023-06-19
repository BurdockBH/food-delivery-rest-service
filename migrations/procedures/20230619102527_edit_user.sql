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
    IF EXISTS (SELECT id FROM users WHERE email = email) THEN
        UPDATE users SET name = inName, email = inEmail, password = inHashedPassword, phone = inPhone, updated_at = inUpdatedAt WHERE email = email;
        SELECT 'UPDATED';
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
