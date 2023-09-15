-- name: UserOnlineHourlyStats :many
SELECT user_id, online
FROM user_online_hourly_stats
ORDER BY user_id, online;

-- name: UserOnlineHourlyStatsUpsert :execrows
INSERT INTO user_online_hourly_stats (user_id, online)
VALUES (unnest(@user_ids::BIGINT[]),
        unnest(@onlines::TIMESTAMP[]))
ON CONFLICT (user_id, online) DO NOTHING;

-- name: UserOnlineDailyStatsUpsert :execrows
INSERT INTO user_online_daily_stats (online, user_id)
VALUES (unnest(@onlines::DATE[]),
        unnest(@user_ids::BIGINT[]))
ON CONFLICT (online, user_id) DO NOTHING;

-- name: UserOnlineDailyCountStatsUpsert :exec
INSERT INTO user_online_daily_count_stats (online, user_count)
SELECT source.online, COUNT(*)
FROM user_online_daily_stats AS source
WHERE source.online >= @online
GROUP BY source.online
ORDER BY source.online
ON CONFLICT (online)
    DO UPDATE
    SET user_count = excluded.user_count;
