INSERT INTO vacancies (id, company_id, created_at, created_by)
SELECT vacancy_id,
       1 AS company_id, -- hardcode company_id for now
       '2024-09-17 12:00:00',
       1
FROM GENERATE_SERIES(1, 4096) AS vacancy_id
ON CONFLICT DO NOTHING;

SELECT SETVAL('vacancies_id_seq', (SELECT COALESCE(MAX(id), 1) FROM vacancies));
