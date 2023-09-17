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
