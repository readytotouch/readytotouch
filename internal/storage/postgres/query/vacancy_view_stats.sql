-- name: UserVacancyViewHourlyStatsUpsert :execrows
INSERT INTO user_vacancy_view_hourly_stats (vacancy_id, created_at, user_id)
VALUES (@vacancy_id, @created_at, @user_id)
ON CONFLICT (vacancy_id, created_at, user_id) DO NOTHING;

-- name: UserVacancyViewDailyStatsUpsert :execrows
INSERT INTO user_vacancy_view_daily_stats (vacancy_id, created_at, user_id)
VALUES (@vacancy_id, @created_at, @user_id)
ON CONFLICT (vacancy_id, created_at, user_id) DO NOTHING;

-- name: VacancyViewDailyStatsUpsert :exec
INSERT INTO vacancy_view_daily_stats AS t (vacancy_id, created_at, user_count)
VALUES (@vacancy_id, @created_at, 1)
ON CONFLICT (vacancy_id, created_at) DO UPDATE
    SET user_count = t.user_count + 1;
