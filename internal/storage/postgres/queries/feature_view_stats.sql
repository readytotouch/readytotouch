-- name: FeatureViewDailyStatsUpsert :exec
INSERT INTO feature_view_daily_stats AS fvds (feature, created_at, view_count)
VALUES (@feature, @created_at, 1)
ON CONFLICT (feature, created_at) DO UPDATE
    SET view_count = fvds.view_count + 1;

-- name: FeatureViewStatsUpsert :exec
INSERT INTO feature_view_stats AS fvs (feature, view_count)
VALUES (@feature, 1)
ON CONFLICT (feature) DO UPDATE
    SET view_count = fvs.view_count + 1;

-- name: FeatureViewDailyStats :many
SELECT days.day::DATE                    AS day,
       COALESCE(s.view_count, 0)::BIGINT AS view_count
FROM GENERATE_SERIES(
    sqlc.arg('from')::DATE,
    sqlc.arg('to')::DATE,
    '1 DAY'::INTERVAL
) AS days (day)
    LEFT JOIN feature_view_daily_stats s ON (days.day = s.created_at AND s.feature = @feature)
ORDER BY days.day;

-- name: FeatureViewStats :one
SELECT view_count
FROM feature_view_stats
WHERE feature = @feature;
