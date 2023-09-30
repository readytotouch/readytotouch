-- name: UserSocialProfileGet :one
SELECT id, user_id, email, username, name
FROM user_social_profiles usp
WHERE usp.social_provider = @social_provider
  AND usp.social_provider_user_id = @social_provider_user_id
  AND usp.deleted_at IS NULL
    FOR UPDATE;

-- name: UserSocialProfileUpdate :exec
UPDATE user_social_profiles
SET email      = @email,
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
SELECT u.id,
       s.social_provider,
       s.social_provider_user_id,
       s.username,
       s.name,
       u.created_at
FROM users u
         INNER JOIN user_social_profiles s ON u.id = s.user_id
WHERE s.deleted_at IS NULL
ORDER BY s.id DESC
LIMIT sqlc.arg('limit');
