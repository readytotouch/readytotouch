-- name: UsersNew :one
INSERT INTO users (created_at, updated_at)
VALUES (@created_at, @created_at)
RETURNING id;

-- name: UsersUpdate :exec
UPDATE users
SET updated_at = @updated_at
WHERE id = @id;

-- name: UserSocialProfileGetByID :one
SELECT usp.id, usp.user_id, usp.email, usp.username, usp.name
FROM user_social_profiles usp
WHERE usp.social_provider = @social_provider
  AND usp.social_provider_user_id = @social_provider_user_id
  AND usp.deleted_at IS NULL
    FOR UPDATE;

-- name: UserSocialProfileGetUserByEmail :one
SELECT usp.user_id
FROM user_social_profiles usp
WHERE usp.email = LOWER(@email)
  AND usp.deleted_at IS NULL
ORDER BY usp.id DESC
LIMIT 1;

-- name: UserSocialProfileNew :one
INSERT INTO user_social_profiles (user_id, social_provider, social_provider_user_id, email, username, name, created_at, updated_at)
VALUES (@user_id, @social_provider, @social_provider_user_id, @email, @username, @name, @created_at, @created_at)
RETURNING id;

-- name: UserSocialProfileUpdate :exec
UPDATE user_social_profiles
SET email      = LOWER(@email),
    username   = @username,
    name       = @name,
    updated_at = @updated_at
WHERE id = @id;

-- name: UserSocialProfileChangeHistoryNew :exec
INSERT INTO user_social_profile_change_history (user_id, user_social_profile_id, email, username, name, created_at)
VALUES (@user_id, @user_social_profile_id, @email, @username, @name, @created_at);

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
WHERE s.deleted_at IS NULL
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
