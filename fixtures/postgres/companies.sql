INSERT INTO companies (id, created_at, created_by)
SELECT company_id,
       '2024-09-17 12:00:00',
       1
FROM GENERATE_SERIES(1, 2048) AS company_id
ON CONFLICT DO NOTHING;

SELECT SETVAL('companies_id_seq', (SELECT COALESCE(MAX(id), 1) FROM companies));
