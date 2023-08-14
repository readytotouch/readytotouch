-- name: UserOnlineUpsert :exec
INSERT INTO user_online (user_id, online)
VALUES (@user_id, @online)
ON CONFLICT (user_id) DO UPDATE
    SET online = @online;

-- name: UserOnlineBatchUpdate :exec
UPDATE user_online AS to_t
SET online = from_t.online
FROM (
         SELECT unnest(@user_ids::BIGINT[])   AS user_id,
                unnest(@onlines::TIMESTAMP[]) AS online
     ) AS from_t
WHERE to_t.user_id = from_t.user_id;
