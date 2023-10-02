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
    updated_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP NULL
);

CREATE TABLE user_social_profiles
(
    id                      BIGSERIAL       NOT NULL PRIMARY KEY,
    user_id                 BIGINT          NOT NULL REFERENCES users (id),
    social_provider         SOCIAL_PROVIDER NOT NULL,
    social_provider_user_id VARCHAR         NOT NULL, -- erase on soft delete
    email                   VARCHAR         NOT NULL, -- erase on soft delete
    username                VARCHAR         NOT NULL, -- erase on soft delete
    name                    VARCHAR         NOT NULL, -- erase on soft delete
    created_at              TIMESTAMP       NOT NULL,
    updated_at              TIMESTAMP       NOT NULL,
    deleted_at              TIMESTAMP       NULL
);

CREATE UNIQUE INDEX
    ON user_social_profiles (social_provider, social_provider_user_id)
    WHERE deleted_at IS NULL;

CREATE INDEX
    ON user_social_profiles (email)
    WHERE deleted_at IS NULL;

CREATE TABLE user_social_profile_change_history
(
    id                     BIGSERIAL NOT NULL PRIMARY KEY,
    user_id                BIGINT    NOT NULL REFERENCES users (id),
    user_social_profile_id BIGINT    NOT NULL REFERENCES user_social_profiles (id),
    email                  VARCHAR   NOT NULL, -- erase on soft delete
    username               VARCHAR   NOT NULL, -- erase on soft delete
    name                   VARCHAR   NOT NULL, -- erase on soft delete
    created_at             TIMESTAMP NOT NULL,
    deleted_at             TIMESTAMP NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE user_social_profile_change_history;
DROP TABLE user_social_profiles;
DROP TABLE users;

DROP TYPE SOCIAL_PROVIDER;
-- +goose StatementEnd
