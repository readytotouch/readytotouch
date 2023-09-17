INSERT INTO users (id, created_at, updated_at)
SELECT generate_series, '2023-09-17 12:00:00', '2023-09-17 12:00:00'
FROM generate_series(1, 1000 * 1000)
ON CONFLICT (id) DO UPDATE
    SET updated_at = '2023-09-17 12:00:00';
