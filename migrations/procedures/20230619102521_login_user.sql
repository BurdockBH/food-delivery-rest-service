-- +goose Up
-- +goose StatementBegin
DELIMITER //
CREATE DEFINER=`root`@`localhost` PROCEDURE `LoginUser`(
    IN inEmail VARCHAR(140)
)
BEGIN
    SELECT password FROM users WHERE email = inEmail;
END //
DELIMITER ;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELIMITER //
DROP PROCEDURE IF EXISTS `LoginUser`;
DELIMITER ;
-- +goose StatementEnd
