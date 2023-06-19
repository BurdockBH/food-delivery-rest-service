-- +goose Up
-- +goose StatementBegin
DELIMITER //
CREATE DEFINER=`root`@`localhost` PROCEDURE `DeleteUser`(
    IN inEmail VARCHAR(140)
)
BEGIN
    DECLARE userId INT;
    SELECT id INTO userId FROM users WHERE email = inEmail;

    IF userId IS NOT NULL THEN
        DELETE FROM users WHERE id = userId AND email = inEmail;
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
