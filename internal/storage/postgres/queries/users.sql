-- name: UserSocialProfileUpsert :one
WITH prev AS (
    SELECT usp.id, usp.user_id, usp.email, usp.username, usp.name
    FROM user_social_profiles usp
    WHERE usp.social_provider = @social_provider
      AND usp.social_provider_user_id = @social_provider_user_id
      AND usp.deleted_at IS NULL
        FOR UPDATE
),
changed AS (
    -- the email is compared case-insensitively, the username and the name are case-sensitive
    SELECT prev.id, prev.user_id, prev.email, prev.username, prev.name
    FROM prev
    WHERE LOWER(prev.email) IS DISTINCT FROM LOWER(@email::VARCHAR)
       OR prev.username IS DISTINCT FROM @username::VARCHAR
       OR prev.name IS DISTINCT FROM @name::VARCHAR
),
updated_profile AS (
    UPDATE user_social_profiles
    SET email      = LOWER(@email::VARCHAR),
        username   = @username::VARCHAR,
        name       = @name::VARCHAR,
        updated_at = @created_at
    FROM changed
    WHERE user_social_profiles.id = changed.id
),
updated_user AS (
    UPDATE users AS t
    SET updated_at = @created_at
    FROM changed
    WHERE t.id = changed.user_id
),
updated_history AS (
    -- the history is not written if only the case has changed
    INSERT INTO user_social_profile_change_history (user_id, user_social_profile_id, email, username, name, created_at)
    SELECT changed.user_id, changed.id, @email::VARCHAR, @username::VARCHAR, @name::VARCHAR, @created_at
    FROM changed
    WHERE LOWER(changed.email) IS DISTINCT FROM LOWER(@email::VARCHAR)
       OR LOWER(changed.username) IS DISTINCT FROM LOWER(@username::VARCHAR)
       OR LOWER(changed.name) IS DISTINCT FROM LOWER(@name::VARCHAR)
),
existing_user AS (
    SELECT usp.user_id AS id
    FROM user_social_profiles usp
    WHERE NOT EXISTS (SELECT 1 FROM prev)
      AND usp.email = LOWER(@email::VARCHAR)
      AND usp.deleted_at IS NULL
    ORDER BY usp.id DESC
    LIMIT 1
),
new_user AS (
    INSERT INTO users (created_at, updated_at)
    SELECT @created_at, @created_at
    WHERE NOT EXISTS (SELECT 1 FROM prev)
      AND NOT EXISTS (SELECT 1 FROM existing_user)
    RETURNING id
),
target_user AS (
    SELECT id FROM existing_user
    UNION ALL
    SELECT id FROM new_user
),
new_profile AS (
    INSERT INTO user_social_profiles (user_id, social_provider, social_provider_user_id, email, username, name, created_at, updated_at)
    SELECT target_user.id, @social_provider, @social_provider_user_id, @email::VARCHAR, @username::VARCHAR, @name::VARCHAR, @created_at, @created_at
    FROM target_user
    RETURNING id, user_id
),
new_history AS (
    INSERT INTO user_social_profile_change_history (user_id, user_social_profile_id, email, username, name, created_at)
    SELECT new_profile.user_id, new_profile.id, @email::VARCHAR, @username::VARCHAR, @name::VARCHAR, @created_at
    FROM new_profile
)
SELECT result.user_id::BIGINT AS user_id
FROM (
    SELECT prev.user_id FROM prev
    UNION ALL
    SELECT new_profile.user_id FROM new_profile
) AS result;

-- name: UserRegistrationDailyCountStats :many
SELECT days.day::DATE                    AS day,
       COALESCE(s.user_count, 0)::BIGINT AS user_count
FROM GENERATE_SERIES(
    sqlc.arg('from')::DATE,
    sqlc.arg('to')::DATE,
    '1 DAY'::INTERVAL
) AS days (day)
    LEFT JOIN (
        SELECT DATE_TRUNC('DAY', created_at) AS day,
               COUNT(*)                      AS user_count
        FROM users
        WHERE created_at >= sqlc.arg('from')::DATE
        GROUP BY day
    ) AS s ON (days.day = s.day)
ORDER BY days.day;

-- name: SocialUserProfiles :many
SELECT s.social_provider,
       s.social_provider_user_id,
       s.username,
       s.name,
       s.created_at
FROM user_social_profiles s
WHERE (
       @social_provider_filter_exists::BOOLEAN = FALSE
    OR s.social_provider = ANY (@social_providers::SOCIAL_PROVIDER[])
  )
  AND s.deleted_at IS NULL
ORDER BY s.id DESC
LIMIT sqlc.arg('limit');

-- name: SocialUserProfilesByUser :many
SELECT s.social_provider,
       s.social_provider_user_id,
       s.username,
       s.name,
       s.created_at
FROM user_social_profiles s
WHERE s.user_id = @user_id
  AND s.deleted_at IS NULL
ORDER BY s.id;
