TRUNCATE user_feature_waitlist;

INSERT INTO user_feature_waitlist (user_id, feature, active, created_at, updated_at)
SELECT user_id,
       'organizer_golang_vacancies'::FEATURE_WAIT,
       TRUE,
       LEAST(NOW() - INTERVAL '1 MONTH' + INTERVAL '1 HOUR' * user_id, NOW()),
       LEAST(NOW() - INTERVAL '1 MONTH' + INTERVAL '1 HOUR' * user_id, NOW())
FROM GENERATE_SERIES(1, 800) AS user_id
ON CONFLICT (user_id, feature) DO UPDATE
    SET updated_at = excluded.updated_at;

INSERT INTO user_feature_waitlist_stats (feature, user_count)
SELECT feature, COUNT(*)
FROM user_feature_waitlist
GROUP BY feature
ON CONFLICT (feature) DO UPDATE
    SET user_count = excluded.user_count;
