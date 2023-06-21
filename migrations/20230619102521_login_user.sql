-- +goose Up
-- +goose StatementBegin
CREATE PROCEDURE `LoginUser`(
    IN inEmail VARCHAR(140)
)
BEGIN
    SELECT password FROM users WHERE email = inEmail;
END

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP PROCEDURE IF EXISTS `LoginUser`;

-- +goose StatementEnd
