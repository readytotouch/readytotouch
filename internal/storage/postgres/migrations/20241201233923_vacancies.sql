-- +goose Up
-- +goose StatementBegin
CREATE TABLE vacancies
(
    id         BIGSERIAL NOT NULL PRIMARY KEY,
    company_id BIGINT    NOT NULL REFERENCES companies (id),
    created_at TIMESTAMP NOT NULL,
    created_by BIGINT    NOT NULL REFERENCES users (id)
);

CREATE TABLE user_favorite_vacancies
(
    user_id    BIGINT                   NOT NULL REFERENCES users (id),
    vacancy_id BIGINT                   NOT NULL REFERENCES vacancies (id),
    favorite   BOOLEAN                  NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL,
    PRIMARY KEY (user_id, vacancy_id)
);

CREATE TABLE user_vacancy_view_hourly_stats
(
    vacancy_id BIGINT    NOT NULL REFERENCES vacancies (id),
    created_at TIMESTAMP NOT NULL,
    user_id    BIGINT    NOT NULL,
    PRIMARY KEY (vacancy_id, created_at, user_id)
);

CREATE TABLE user_vacancy_view_daily_stats
(
    vacancy_id BIGINT NOT NULL REFERENCES vacancies (id),
    created_at DATE   NOT NULL,
    user_id    BIGINT NOT NULL,
    PRIMARY KEY (vacancy_id, created_at, user_id)
);

CREATE TABLE vacancy_view_daily_stats
(
    vacancy_id BIGINT NOT NULL REFERENCES vacancies (id),
    created_at DATE   NOT NULL,
    user_count BIGINT NOT NULL,
    PRIMARY KEY (vacancy_id, created_at)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE vacancy_view_daily_stats;
DROP TABLE user_vacancy_view_daily_stats;
DROP TABLE user_vacancy_view_hourly_stats;
DROP TABLE user_favorite_vacancies;
DROP TABLE vacancies;
-- +goose StatementEnd
