-- name: UserCompanyVisibilityOverrideUpsert :execrows
INSERT INTO user_company_visibility_overrides AS t (user_id, company_id, visibility, created_at, updated_at)
VALUES (@user_id, @company_id, @visibility, @created_at, @created_at)
ON CONFLICT (user_id, company_id) DO UPDATE
    SET visibility = excluded.visibility,
        updated_at = excluded.updated_at
WHERE t.visibility <> excluded.visibility;

-- name: UserCompanyVisibilityOverrideHistoryNew :exec
INSERT INTO user_company_visibility_override_history (user_id, company_id, visibility, created_at)
VALUES (@user_id, @company_id, @visibility, @created_at);

-- name: UserIndustryVisibilityOverrideUpsert :execrows
INSERT INTO user_industry_visibility_overrides AS t (user_id, industry_id, visibility, created_at, updated_at)
VALUES (@user_id, @industry_id, @visibility, @created_at, @created_at)
ON CONFLICT (user_id, industry_id) DO UPDATE
    SET visibility = excluded.visibility,
        updated_at = excluded.updated_at
WHERE t.visibility <> excluded.visibility;

-- name: UserIndustryVisibilityOverrideHistoryNew :exec
INSERT INTO user_industry_visibility_override_history (user_id, industry_id, visibility, created_at)
VALUES (@user_id, @industry_id, @visibility, @created_at);
