BEGIN;

CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid()
);

CREATE TABLE boards (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL,
    name VARCHAR NOT NULL,
    access VARCHAR DEFAULT 'private',
    color VARCHAR NOT NULL
);

ALTER TABLE boards
    ADD CONSTRAINT fk_boards_user
    FOREIGN KEY (user_id)
    REFERENCES users (id)
    ON DELETE CASCADE;

CREATE INDEX idx_boards_user_id ON boards(user_id);

COMMIT;