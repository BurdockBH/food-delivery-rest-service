-- +goose Up
-- +goose StatementBegin
DELIMITER //
CREATE DEFINER=`root`@`localhost` PROCEDURE `LoginUser`(
    IN inEmail VARCHAR(140),
    IN inPassword text
)
BEGIN

    IF EXISTS (SELECT id FROM users WHERE email = inEmail AND password = inPassword) THEN
        SELECT 'LOGGED IN';
    END IF;

END //
DELIMITER ;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELIMITER //
DROP PROCEDURE IF EXISTS `LoginUser`;
DELIMITER ;
-- +goose StatementEnd
