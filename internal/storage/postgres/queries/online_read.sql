-- name: UserOnlineHourlyStats :many
SELECT user_id, created_at
FROM user_online_hourly_stats
ORDER BY user_id, created_at;

-- name: UserOnlineDailyCountStats :many
SELECT days.day::DATE                    AS day,
       COALESCE(s.user_count, 0)::BIGINT AS user_count
FROM GENERATE_SERIES(
    sqlc.arg('from')::DATE,
    sqlc.arg('to')::DATE,
    '1 DAY'::INTERVAL
) AS days (day)
    LEFT JOIN user_online_daily_count_stats s ON (days.day = s.created_at)
ORDER BY days.day;
