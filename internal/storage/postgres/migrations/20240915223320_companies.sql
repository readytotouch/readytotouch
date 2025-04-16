-- +goose Up
-- +goose StatementBegin
CREATE TABLE companies
(
    id         BIGSERIAL NOT NULL PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    created_by BIGINT    NOT NULL REFERENCES users (id)
);

CREATE TABLE user_favorite_companies
(
    user_id    BIGINT                   NOT NULL REFERENCES users (id),
    company_id BIGINT                   NOT NULL REFERENCES companies (id),
    favorite   BOOLEAN                  NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL,
    PRIMARY KEY (user_id, company_id)
);

CREATE TABLE company_view_daily_stats
(
    company_id BIGINT NOT NULL REFERENCES companies (id),
    created_at DATE   NOT NULL,
    view_count BIGINT NOT NULL,
    PRIMARY KEY (company_id, created_at)
);

INSERT INTO companies (id, created_at, created_by)
SELECT company_id,
       '2024-09-17 12:00:00',
       1
FROM GENERATE_SERIES(1, 2048) AS company_id
ON CONFLICT DO NOTHING;

SELECT SETVAL('companies_id_seq', (SELECT COALESCE(MAX(id), 1) FROM companies));
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE company_view_daily_stats;
DROP TABLE user_favorite_companies;
DROP TABLE companies;
-- +goose StatementEnd
