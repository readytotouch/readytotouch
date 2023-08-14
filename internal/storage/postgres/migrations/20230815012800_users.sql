-- +goose Up
-- +goose StatementBegin
CREATE TABLE users
(
    id         BIGSERIAL                NOT NULL PRIMARY KEY,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL
);

CREATE TABLE user_online
(
    user_id BIGINT PRIMARY KEY,
    online  TIMESTAMP NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE user_online;

DROP TABLE users;
-- +goose StatementEnd
