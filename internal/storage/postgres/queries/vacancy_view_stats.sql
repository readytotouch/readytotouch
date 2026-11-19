-- name: VacancyViewUpsert :exec
WITH hourly AS (
    INSERT INTO user_vacancy_view_hourly_stats (vacancy_id, created_at, user_id)
    VALUES (@vacancy_id, @hour::TIMESTAMP, @user_id)
    ON CONFLICT (vacancy_id, created_at, user_id) DO NOTHING
    RETURNING vacancy_id, user_id
),
daily AS (
    INSERT INTO user_vacancy_view_daily_stats (vacancy_id, created_at, user_id)
    SELECT hourly.vacancy_id, @day::DATE, hourly.user_id
    FROM hourly
    ON CONFLICT (vacancy_id, created_at, user_id) DO NOTHING
    RETURNING vacancy_id, created_at
)
INSERT INTO vacancy_view_daily_stats AS t (vacancy_id, created_at, user_count)
SELECT daily.vacancy_id, daily.created_at, 1
FROM daily
ON CONFLICT (vacancy_id, created_at) DO UPDATE
    SET user_count = t.user_count + 1;

-- name: VacancyViewStats :many
SELECT s.vacancy_id,
       SUM(user_count) AS count
FROM vacancy_view_daily_stats s
WHERE s.vacancy_id = ANY (@vacancy_ids::BIGINT[])
  AND s.created_at >= sqlc.arg('from')::DATE
GROUP BY s.vacancy_id;
