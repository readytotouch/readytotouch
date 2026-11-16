-- name: UserFavoriteVacanciesUpsert :exec
INSERT INTO user_favorite_vacancies AS t (user_id, vacancy_id, favorite, created_at, updated_at)
VALUES (@user_id, @vacancy_id, @favorite, @created_at, @created_at)
ON CONFLICT (user_id, vacancy_id)
    DO UPDATE
    SET favorite   = excluded.favorite,
        updated_at = excluded.updated_at
WHERE t.favorite <> excluded.favorite;

-- name: UserFavoriteVacancies :many
SELECT ufv.vacancy_id
FROM user_favorite_vacancies ufv
WHERE ufv.user_id = @user_id
  AND (@vacancy_ids_filter_exists::BOOLEAN = FALSE OR ufv.vacancy_id = ANY (@vacancy_ids::BIGINT[]))
  AND ufv.favorite = TRUE;
