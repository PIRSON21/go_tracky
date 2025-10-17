-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd
ALTER TABLE boards
    ADD CONSTRAINT unique_board_user
        UNIQUE (user_id, name);


-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
ALTER TABLE boards
    DROP CONSTRAINT unique_board_user;
-- +goose StatementEnd
