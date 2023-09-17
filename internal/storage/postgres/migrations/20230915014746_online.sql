-- +goose Up
-- +goose StatementBegin
CREATE TABLE user_online_hourly_stats
(
    user_id    BIGINT    NOT NULL,
    created_at TIMESTAMP NOT NULL,
    PRIMARY KEY (user_id, created_at)
);

CREATE TABLE user_online_daily_stats
(
    created_at DATE   NOT NULL,
    user_id    BIGINT NOT NULL,
    PRIMARY KEY (created_at, user_id)
);

CREATE TABLE user_online_daily_count_stats
(
    created_at DATE   NOT NULL PRIMARY KEY,
    user_count BIGINT NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE user_online_daily_count_stats;
DROP TABLE user_online_daily_stats;
DROP TABLE user_online_hourly_stats;
-- +goose StatementEnd
