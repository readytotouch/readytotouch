-- name: CompanyViewDailyStatsUpsert :exec
INSERT INTO company_view_daily_stats AS cvds (company_id, created_at, view_count)
VALUES (@company_id, @created_at, 1)
ON CONFLICT (company_id, created_at) DO UPDATE
    SET view_count = cvds.view_count + 1;

-- name: CompanyViewDailyStats :many
SELECT days.day::DATE                    AS day,
       COALESCE(s.view_count, 0)::BIGINT AS view_count
FROM GENERATE_SERIES(
    sqlc.arg('from')::DATE,
    sqlc.arg('to')::DATE,
    '1 DAY'::INTERVAL
) AS days (day)
    LEFT JOIN company_view_daily_stats s ON (days.day = s.created_at AND s.company_id = @company_id)
ORDER BY days.day;

-- name: CompanyTotalViews :one
SELECT SUM(view_count)                                                       AS count,
       SUM(view_count) FILTER ( WHERE created_at >= sqlc.arg('from')::DATE ) AS count_since
FROM company_view_daily_stats s
WHERE s.company_id = @company_id;
