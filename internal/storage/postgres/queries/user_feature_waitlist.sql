-- name: UserFeatureWaitlistUpsert :execrows
INSERT INTO user_feature_waitlist AS ufw (user_id, feature, active, created_at, updated_at)
VALUES (@user_id, @feature, @active, @created_at, @created_at)
ON CONFLICT (user_id, feature) DO UPDATE
    SET active     = excluded.active,
        updated_at = excluded.updated_at
WHERE ufw.active <> excluded.active;

-- name: UserFeatureWaitlistStatsUpsert :exec
INSERT INTO user_feature_waitlist_stats AS ufws (feature, user_count)
VALUES (@feature, @user_count)
ON CONFLICT (feature) DO UPDATE
    SET user_count = ufws.user_count + excluded.user_count;

-- name: UserFeatureWaitlist :one
SELECT active
FROM user_feature_waitlist
WHERE user_id = @user_id
  AND feature = @feature;

-- name: UserFeatureWaitlistDailyStats :many
WITH aggs AS (
    SELECT created_at::DATE AS day,
           COUNT(*)         AS subscriber_count
    FROM user_feature_waitlist
    WHERE feature = @feature
      AND created_at::DATE >= sqlc.arg('from')::DATE
    GROUP BY day
)
SELECT days.day::DATE                             AS day,
       COALESCE(aggs.subscriber_count, 0)::BIGINT AS subscriber_count
FROM GENERATE_SERIES(
    sqlc.arg('from')::DATE,
    sqlc.arg('to')::DATE,
    '1 DAY'::INTERVAL
) AS days (day)
    LEFT JOIN aggs ON (days.day = aggs.day)
ORDER BY days.day;

-- name: UserFeatureWaitlistStats :one
SELECT user_count
FROM user_feature_waitlist_stats
WHERE feature = @feature;
