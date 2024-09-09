INSERT INTO feature_view_stats (feature, view_count)
VALUES ('organizer_golang_vacancies', 1234)
ON CONFLICT (feature) DO UPDATE
    SET view_count = excluded.view_count;

INSERT INTO feature_view_daily_stats (feature, created_at, view_count)
SELECT 'organizer_golang_vacancies'::FEATURE_WAIT, g.day, FLOOR(RANDOM() * 50)::INT + 50
FROM (
    SELECT day::DATE
    FROM GENERATE_SERIES(
        (DATE_TRUNC('DAY', NOW()) - INTERVAL '1 MONTH')::TIMESTAMP,
        (DATE_TRUNC('DAY', NOW()))::TIMESTAMP,
        '1 DAY'::INTERVAL
    ) AS day
) AS g
ON CONFLICT (feature, created_at) DO UPDATE
    SET view_count = excluded.view_count;