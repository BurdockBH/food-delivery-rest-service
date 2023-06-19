-- +goose Up
-- +goose StatementBegin
DELIMITER //
CREATE DEFINER=`root`@`localhost` PROCEDURE `RegisterUser`(
    IN inName VARCHAR(50),
    IN inEmail VARCHAR(140),
    IN inHashedPassword TEXT,
    IN inPhone VARCHAR(20),
    IN inCreatedAt bigint,
    IN inUpdatedAt bigint
)
BEGIN

    IF (SELECT COUNT(*) FROM users WHERE email = inEmail) > 0 THEN
        SELECT 0;
    ELSE
        INSERT INTO users (name, email, password, phone, created_at, updated_at)
        VALUES (inName, inEmail, inHashedPassword, inPhone, inCreatedAt, inUpdatedAt);

        SELECT 1;
    END IF;
END //

DELIMITER ;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELIMITER //
DROP PROCEDURE IF EXISTS `RegisterUser`;
DELIMITER ;
-- +goose StatementEnd
