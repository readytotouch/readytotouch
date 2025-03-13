-- +goose Up
-- +goose StatementBegin
CREATE TABLE vacancies
(
    id         BIGSERIAL NOT NULL PRIMARY KEY,
    company_id BIGINT    NOT NULL REFERENCES companies (id),
    created_at TIMESTAMP NOT NULL,
    created_by BIGINT    NOT NULL REFERENCES users (id)
);

-- INSERT INTO vacancies (id, company_id, created_at, created_by) VALUES (1, 1, '2024-12-03 12:00:00', 1);

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

INSERT INTO vacancies (id, company_id, created_at, created_by)
SELECT vacancy_id,
       1 AS company_id, -- hardcode company_id for now
       '2024-09-17 12:00:00',
       1
FROM GENERATE_SERIES(1, 2048) AS vacancy_id
ON CONFLICT DO NOTHING;

SELECT SETVAL('vacancies_id_seq', (SELECT COALESCE(MAX(id), 1) FROM vacancies));
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE vacancy_view_daily_stats;
DROP TABLE user_vacancy_view_daily_stats;
DROP TABLE user_vacancy_view_hourly_stats;
DROP TABLE user_favorite_vacancies;
DROP TABLE vacancies;
-- +goose StatementEnd
