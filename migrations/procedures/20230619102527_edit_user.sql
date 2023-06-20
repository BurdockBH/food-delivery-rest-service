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
    DECLARE existingUserPhoneCount INT;

    -- Check if the user with the specified email exists
    SELECT id INTO userId FROM users WHERE email = inEmail;

    IF userId IS NULL THEN
        -- The user with the specified email does not exist, return an error
        SELECT -1;
    ELSE
        -- Check if the phone number already exists for another user
        SELECT COUNT(*) INTO existingUserPhoneCount FROM users WHERE phone = inPhone;

        IF existingUserPhoneCount > 0 THEN
            -- Another account with the same phone number already exists, return an error
            SELECT -2;
        ELSE
            -- Update the user
            UPDATE users SET name = inName, password = inHashedPassword, phone = inPhone, updated_at = inUpdatedAt WHERE id = userId AND email = inEmail;
            SELECT 1;
        END IF;
    END IF;
END//
DELIMITER ;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELIMITER //
DROP PROCEDURE IF EXISTS EditUser;
DELIMITER ;
-- +goose StatementEnd
