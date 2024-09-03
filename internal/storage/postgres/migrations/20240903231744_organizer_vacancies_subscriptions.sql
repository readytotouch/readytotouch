-- +goose Up
-- +goose StatementBegin
CREATE TYPE FEATURE_WAIT AS ENUM (
    'organizer_golang_companies', -- available from start
    'organizer_golang_vacancies',
    'organizer_rust_companies',
    'organizer_rust_vacancies',
    'organizer_zig_companies',
    'organizer_zig_vacancies',
    'organizer_scala_companies',
    'organizer_scala_vacancies',
    'organizer_elixir_companies',
    'organizer_elixir_vacancies',
    'organizer_clojure_companies',
    'organizer_clojure_vacancies'
    );

CREATE TABLE user_feature_waitlist_subscriptions
(
    id         BIGSERIAL    NOT NULL PRIMARY KEY,
    user_id    BIGINT       NOT NULL REFERENCES users (id),
    feature    FEATURE_WAIT NOT NULL,
    active     BOOLEAN      NOT NULL,
    created_at TIMESTAMP    NOT NULL,
    updated_at TIMESTAMP    NOT NULL,
    UNIQUE (user_id, feature)
);

CREATE TABLE user_feature_waitlist_subscription_history
(
    id         BIGSERIAL    NOT NULL PRIMARY KEY,
    user_id    BIGINT       NOT NULL REFERENCES users (id),
    feature    FEATURE_WAIT NOT NULL,
    active     BOOLEAN      NOT NULL,
    created_at TIMESTAMP    NOT NULL,
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE user_feature_waitlist_subscription_history;

DROP TABLE user_feature_waitlist_subscriptions;
-- +goose StatementEnd
