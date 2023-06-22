-- +goose Up
-- +goose StatementBegin
CREATE PROCEDURE `DeleteUser`(
    IN inEmail VARCHAR(140)
)
BEGIN
    DECLARE userId INT;
    SELECT id INTO userId FROM users WHERE email = inEmail;

    IF userId IS NOT NULL THEN
        DELETE FROM users WHERE id = userId AND email = inEmail;
        SELECT 1;
    ELSE SELECT 0;
    END IF;
END
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP PROCEDURE IF EXISTS `DeleteUser`;
-- +goose StatementEnd
