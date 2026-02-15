-- +goose Up
-- +goose StatementBegin
CREATE TABLE user_company_visibility_overrides
(
    user_id    BIGINT                   NOT NULL REFERENCES users (id),
    company_id BIGINT                   NOT NULL REFERENCES companies (id),
    visibility BOOLEAN                  NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL,
    PRIMARY KEY (user_id, company_id)
);

CREATE TABLE user_company_visibility_override_history
(
    id         BIGSERIAL                NOT NULL PRIMARY KEY,
    user_id    BIGINT                   NOT NULL REFERENCES users (id),
    company_id BIGINT                   NOT NULL REFERENCES companies (id),
    visibility BOOLEAN                  NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL
);


CREATE TABLE user_industry_visibility_overrides
(
    user_id     BIGINT                   NOT NULL REFERENCES users (id),
    industry_id BIGINT                   NOT NULL REFERENCES industries (id),
    visibility  BOOLEAN                  NOT NULL,
    created_at  TIMESTAMP WITH TIME ZONE NOT NULL,
    updated_at  TIMESTAMP WITH TIME ZONE NOT NULL,
    PRIMARY KEY (user_id, industry_id)
);

CREATE TABLE user_industry_visibility_override_history
(
    id          BIGSERIAL                NOT NULL PRIMARY KEY,
    user_id     BIGINT                   NOT NULL REFERENCES users (id),
    industry_id BIGINT                   NOT NULL REFERENCES industries (id),
    visibility  BOOLEAN                  NOT NULL,
    created_at  TIMESTAMP WITH TIME ZONE NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE user_industry_visibility_override_history;
DROP TABLE user_industry_visibility_overrides;
DROP TABLE user_company_visibility_override_history;
DROP TABLE user_company_visibility_overrides;
-- +goose StatementEnd
