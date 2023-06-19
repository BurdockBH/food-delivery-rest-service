-- +goose Up
-- +goose StatementBegin
DELIMITER //
CREATE DEFINER=`root`@`localhost` PROCEDURE `DeleteUser`(
    IN inEmail VARCHAR(140),
    IN inHashedPassword TEXT
)
BEGIN

    IF EXISTS (SELECT id FROM users WHERE Email = inEmail) THEN
        DELETE FROM users WHERE email = inEmail AND password = inHashedPassword;
        SELECT 'DELETED';
    END IF;

END //
DELIMITER ;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELIMITER //
DROP PROCEDURE IF EXISTS `DeleteUser`;
DELIMITER ;
-- +goose StatementEnd
