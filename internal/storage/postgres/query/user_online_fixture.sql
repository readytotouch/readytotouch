-- name: UserOnlineFixtureUpsert :exec
INSERT INTO user_online (user_id, online)
SELECT generate_series,
       to_timestamp(@online::BIGINT)
FROM generate_series(1, @count::BIGINT)
ON CONFLICT (user_id) DO UPDATE
    SET online = excluded.online;

-- name: UserOnlineFixtureCount :one
SELECT COUNT(*) AS total,
       SUM(
               CASE online
                   WHEN TO_TIMESTAMP(@online::BIGINT) THEN 0
                   ELSE 1
                   END
           )    AS changed
FROM user_online;
