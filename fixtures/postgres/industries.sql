INSERT INTO industries (id, created_at, created_by)
SELECT industry_id,
       '2026-02-15 12:00:00',
       1
FROM GENERATE_SERIES(1, 256) AS industry_id
ON CONFLICT DO NOTHING;

SELECT SETVAL('industries_id_seq', (SELECT COALESCE(MAX(id), 1) FROM industries));
