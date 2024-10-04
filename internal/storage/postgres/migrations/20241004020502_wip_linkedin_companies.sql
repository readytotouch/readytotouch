-- +goose Up
-- +goose StatementBegin
CREATE TABLE wip_linkedin_companies
(
    id          BIGINT                   NOT NULL PRIMARY KEY,
    vanity_name VARCHAR                  NOT NULL UNIQUE,
    name        VARCHAR                  NOT NULL,
    created_at  TIMESTAMP WITH TIME ZONE NOT NULL,
    created_by  BIGINT                   NOT NULL REFERENCES users (id)
);

CREATE TABLE wip_linkedin_company_request_history
(
    id                  BIGSERIAL                NOT NULL PRIMARY KEY,
    vanity_name         VARCHAR                  NOT NULL UNIQUE,
    linkedin_company_id BIGINT                   NULL REFERENCES wip_linkedin_companies (id),
    response_payload    JSONB                    NOT NULL,
    created_at          TIMESTAMP WITH TIME ZONE NOT NULL,
    created_by          BIGINT                   NOT NULL REFERENCES users (id)
);

CREATE TABLE wip_user_to_linkedin_companies
(
    user_id             BIGINT                   NOT NULL REFERENCES users (id),
    linkedin_company_id BIGINT                   NOT NULL REFERENCES wip_linkedin_companies (id),
    active              BOOLEAN                  NOT NULL,
    created_at          TIMESTAMP WITH TIME ZONE NOT NULL,
    created_by          BIGINT                   NOT NULL REFERENCES users (id),
    updated_at          TIMESTAMP WITH TIME ZONE NOT NULL,
    updated_by          BIGINT                   NOT NULL REFERENCES users (id),
    PRIMARY KEY (user_id, linkedin_company_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE wip_user_to_linkedin_companies;
DROP TABLE wip_linkedin_company_request_history;
DROP TABLE wip_linkedin_companies;
-- +goose StatementEnd
