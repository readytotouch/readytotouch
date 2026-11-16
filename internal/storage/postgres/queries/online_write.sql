-- name: UserOnlineHourlyStatsUpsert :execrows
INSERT INTO user_online_hourly_stats (user_id, created_at)
VALUES (UNNEST(@user_ids::BIGINT[]),
        UNNEST(@created_ats::TIMESTAMP[]))
ON CONFLICT (user_id, created_at) DO NOTHING;

-- name: UserOnlineDailyStatsUpsert :execrows
INSERT INTO user_online_daily_stats (created_at, user_id)
VALUES (UNNEST(@created_ats::DATE[]),
        UNNEST(@user_ids::BIGINT[]))
ON CONFLICT (created_at, user_id) DO NOTHING;

-- name: UserOnlineDailyCountStatsUpsert :exec
INSERT INTO user_online_daily_count_stats (created_at, user_count)
SELECT source.created_at, COUNT(*)
FROM user_online_daily_stats AS source
WHERE source.created_at = ANY (@created_ats::DATE[])
GROUP BY source.created_at
ORDER BY source.created_at
ON CONFLICT (created_at)
    DO UPDATE
    SET user_count = excluded.user_count;
