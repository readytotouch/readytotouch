-- name: UserOnlineHourlyStats :many
SELECT user_id, online
FROM user_online_hourly_stats
ORDER BY user_id, online;

-- name: UserOnlineDailyCountStats :many
SELECT days.online::DATE                 AS online,
       COALESCE(s.user_count, 0)::BIGINT AS user_count
FROM GENERATE_SERIES(
    sqlc.arg('from')::DATE,
    sqlc.arg('to')::DATE,
    '1 DAY'::INTERVAL
) AS days (online)
    LEFT JOIN user_online_daily_count_stats s ON (days.online = s.online)
ORDER BY days.online;
