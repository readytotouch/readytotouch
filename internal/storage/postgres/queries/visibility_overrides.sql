-- name: UserCompanyVisibilityOverrideUpsert :exec
WITH changed AS (
    INSERT INTO user_company_visibility_overrides AS t (user_id, company_id, visibility, created_at, updated_at)
    VALUES (@user_id, @company_id, @visibility, @created_at, @created_at)
    ON CONFLICT (user_id, company_id) DO UPDATE
        SET visibility = excluded.visibility,
            updated_at = excluded.updated_at
    WHERE t.visibility <> excluded.visibility
    RETURNING user_id, company_id, visibility
)
INSERT INTO user_company_visibility_override_history (user_id, company_id, visibility, created_at)
SELECT changed.user_id, changed.company_id, changed.visibility, @created_at
FROM changed;

-- name: UserCompanyVisibilityOverrides :many
SELECT company_id
FROM user_company_visibility_overrides
WHERE user_id = @user_id
  AND visibility = true;

-- name: UserIndustryVisibilityOverrideUpsert :exec
WITH changed AS (
    INSERT INTO user_industry_visibility_overrides AS t (user_id, industry_id, visibility, created_at, updated_at)
    VALUES (@user_id, @industry_id, @visibility, @created_at, @created_at)
    ON CONFLICT (user_id, industry_id) DO UPDATE
        SET visibility = excluded.visibility,
            updated_at = excluded.updated_at
    WHERE t.visibility <> excluded.visibility
    RETURNING user_id, industry_id, visibility
)
INSERT INTO user_industry_visibility_override_history (user_id, industry_id, visibility, created_at)
SELECT changed.user_id, changed.industry_id, changed.visibility, @created_at
FROM changed;

-- name: UserIndustryVisibilityOverrides :many
SELECT industry_id
FROM user_industry_visibility_overrides
WHERE user_id = @user_id
  AND visibility = true;
