-- +goose Up
-- +goose StatementBegin
-- https://survey.stackoverflow.co/2024/technology
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

CREATE TABLE user_feature_waitlist
(
    user_id    BIGINT       NOT NULL REFERENCES users (id),
    feature    FEATURE_WAIT NOT NULL,
    active     BOOLEAN      NOT NULL,
    created_at TIMESTAMP    NOT NULL,
    updated_at TIMESTAMP    NOT NULL,
    PRIMARY KEY (user_id, feature)
);

CREATE TABLE user_feature_waitlist_stats
(
    feature    FEATURE_WAIT NOT NULL PRIMARY KEY,
    user_count BIGINT       NOT NULL
);

CREATE TABLE feature_view_daily_stats
(
    feature    FEATURE_WAIT NOT NULL,
    created_at DATE         NOT NULL,
    view_count BIGINT       NOT NULL,
    PRIMARY KEY (feature, created_at)
);

CREATE TABLE feature_view_stats
(
    feature    FEATURE_WAIT NOT NULL,
    view_count BIGINT       NOT NULL,
    PRIMARY KEY (feature)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE feature_view_stats;
DROP TABLE feature_view_daily_stats;

DROP TABLE user_feature_waitlist_stats;
DROP TABLE user_feature_waitlist;

DROP TYPE FEATURE_WAIT;
-- +goose StatementEnd
