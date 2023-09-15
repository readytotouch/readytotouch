-- +goose Up
-- +goose StatementBegin
CREATE TYPE SOCIAL_PROVIDER AS ENUM (
    'github',-----    -- developers       -- use      -- email = true
    'gitlab',-----    -- developers       -- use      -- email = true
    'bitbucket',--    -- developers       -- use      -- email = true
    'figma',------    -- designers        -- @TODO
    'dribbble',---    -- designers        -- use      -- email = false
    'behance',----    -- designers        -- none provider
    'linkedin',---    -- business         -- use      -- email = true
    'xing',-------    -- business         -- @TODO
    'google'------    -- global           -- use      -- email = true
    );
-- stackoverflow PROVIDER https://api.stackexchange.com/docs/authentication

CREATE TABLE users
(
    id         BIGSERIAL NOT NULL PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);

CREATE TABLE user_social_profiles
(
    id         BIGSERIAL       NOT NULL PRIMARY KEY,
    user_id    BIGINT          NOT NULL REFERENCES users (id),
    provider   SOCIAL_PROVIDER NOT NULL,
    social_id  VARCHAR         NOT NULL,
    email      VARCHAR         NOT NULL,
    username   VARCHAR         NOT NULL,
    created_at TIMESTAMP       NOT NULL,
    updated_at TIMESTAMP       NOT NULL,
    UNIQUE (provider, social_id)
);

CREATE INDEX ON user_social_profiles (email);

CREATE TABLE user_social_profile_change_history
(
    id                     BIGSERIAL NOT NULL PRIMARY KEY,
    user_id                BIGINT    NOT NULL REFERENCES users (id),
    user_social_profile_id BIGINT    NOT NULL REFERENCES user_social_profiles (id),
    email                  VARCHAR   NOT NULL,
    username               VARCHAR   NOT NULL,
    created_at             TIMESTAMP NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE user_social_profile_change_history;
DROP TABLE user_social_profiles;
DROP TABLE users;

DROP TYPE SOCIAL_PROVIDER;
-- +goose StatementEnd
