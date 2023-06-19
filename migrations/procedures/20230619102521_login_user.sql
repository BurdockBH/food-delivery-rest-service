-- +goose Up
-- +goose StatementBegin
CREATE DEFINER=`root`@`localhost` PROCEDURE `LoginUser`(
    IN inEmail VARCHAR(140),
    IN inPassword text
)
BEGIN

    IF EXISTS (SELECT id FROM users WHERE email = inEmail AND password = inPassword) THEN
        SELECT 'LOGGED IN';
    END IF;

END
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP PROCEDURE IF EXISTS `LoginUser`;
-- +goose StatementEnd
