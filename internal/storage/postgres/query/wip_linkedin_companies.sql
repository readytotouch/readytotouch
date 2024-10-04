-- name: WipLinkedInCompaniesGetByVanityName :one
SELECT id
FROM wip_linkedin_companies
WHERE vanity_name = @vanity_name;

-- name: WipLinkedInCompaniesNew :exec
INSERT INTO wip_linkedin_companies (id, vanity_name, name, payload, created_at, created_by)
VALUES (@id, @vanity_name, @name, @payload, @created_at, @created_by);

-- name: WipLinkedInCompanyRequestHistoryNew :exec
INSERT INTO wip_linkedin_company_request_history (vanity_name, linkedin_company_id, created_at, created_by)
VALUES (@vanity_name, @id, @created_at, @created_by);

-- name: WipLinkedInCompanyRequestHistoryExistsVanityName :one
SELECT EXISTS(
    SELECT
    FROM wip_linkedin_company_request_history
    WHERE vanity_name = @vanity_name
);

-- name: WipLinkedInCompanyRequestHistoryCount :one
SELECT COUNT(*) AS total
FROM wip_linkedin_company_request_history
WHERE created_by = @created_by;

-- name: WipUserToLinkedInCompaniesUpsert :exec
INSERT INTO wip_user_to_linkedin_companies AS t (user_id, linkedin_company_id, active, created_at, created_by,
                                                 updated_at, updated_by)
VALUES (@created_by, @linkedin_company_id, @active, @created_at, @created_by,
        @created_at, @created_by)
ON CONFLICT (user_id, linkedin_company_id) DO UPDATE
    SET active     = excluded.active,
        updated_at = excluded.updated_at,
        updated_by = excluded.updated_by;

-- name: WipUserLinkedInCompanies :many
SELECT wlc.id,
       wlc.vanity_name,
       wlc.name
FROM wip_user_to_linkedin_companies wu2lc
         INNER JOIN wip_linkedin_companies wlc ON wu2lc.linkedin_company_id = wlc.id
WHERE wu2lc.user_id = @user_id
  AND wu2lc.active = TRUE
ORDER BY wu2lc.updated_at DESC;
