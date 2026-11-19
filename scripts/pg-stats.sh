#!/usr/bin/env bash
# Read-only statistics of the database, for a quick look over ssh on the server.
#
#   ./scripts/pg-stats.sh [all|size|tables|users|wip-linkedin]
#
# The queries run with psql inside the Postgres container, which knows its own POSTGRES_USER and POSTGRES_DB,
# so it works the same in development and in production. The session is read-only: nothing can be changed.
# PG_CONTAINER overrides the container name.
set -euo pipefail

PG_CONTAINER="${PG_CONTAINER:-readytotouch_postgres_db}"

# psql_exec reads SQL from the standard input
psql_exec() {
	docker exec -i -e PGOPTIONS='-c default_transaction_read_only=on' "$PG_CONTAINER" \
		sh -c 'psql -U "$POSTGRES_USER" -d "$POSTGRES_DB" -X -q -v ON_ERROR_STOP=1 -P pager=off'
}

check_container() {
	if [ "$(docker inspect -f '{{.State.Running}}' "$PG_CONTAINER" 2>/dev/null || true)" != "true" ]; then
		echo "The container $PG_CONTAINER is not running." >&2
		exit 1
	fi
}

stats_size() {
	psql_exec <<'SQL'
\echo
\echo '== Database size'
SELECT current_database()                                  AS database,
       pg_size_pretty(pg_database_size(current_database())) AS size;
SQL
}

stats_tables() {
	# The number of rows is exact (count(*)), it is fast enough for a database of this size.
	psql_exec <<'SQL'
\echo
\echo '== Tables, the biggest first'
WITH tables AS (
    SELECT c.oid,
           c.relname,
           (xpath('/row/cnt/text()', query_to_xml(FORMAT('SELECT COUNT(*) AS cnt FROM %I.%I', n.nspname, c.relname), FALSE, TRUE, '')))[1]::TEXT::BIGINT AS row_count
    FROM pg_class c
             JOIN pg_namespace n ON n.oid = c.relnamespace
    WHERE c.relkind = 'r'
      AND n.nspname = 'public'
)
SELECT relname                                    AS "table",
       row_count                                  AS "rows",
       pg_size_pretty(pg_table_size(oid))         AS table_size,
       pg_size_pretty(pg_indexes_size(oid))       AS indexes_size,
       pg_size_pretty(pg_total_relation_size(oid)) AS total_size
FROM tables
ORDER BY pg_total_relation_size(oid) DESC, relname;

\echo
\echo '== Tables, total'
SELECT COUNT(*)                                                        AS tables,
       pg_size_pretty(SUM(pg_total_relation_size(c.oid))::BIGINT)      AS total_size
FROM pg_class c
         JOIN pg_namespace n ON n.oid = c.relnamespace
WHERE c.relkind = 'r'
  AND n.nspname = 'public';
SQL
}

stats_users() {
	psql_exec <<'SQL'
SET TIME ZONE 'UTC';

\echo
\echo '== Users, total'
SELECT COUNT(*)                                        AS total,
       COUNT(*) FILTER (WHERE deleted_at IS NULL)      AS active,
       COUNT(*) FILTER (WHERE deleted_at IS NOT NULL)  AS deleted
FROM users;

\echo
\echo '== Users registered by month (UTC)'
WITH months AS (
    SELECT GENERATE_SERIES(DATE_TRUNC('month', MIN(created_at)), DATE_TRUNC('month', NOW()), '1 month') AS month
    FROM users
),
registered AS (
    SELECT DATE_TRUNC('month', created_at) AS month,
           COUNT(*)                        AS new_users
    FROM users
    GROUP BY 1
)
SELECT TO_CHAR(m.month, 'YYYY-MM')                        AS month,
       COALESCE(r.new_users, 0)                           AS new_users,
       SUM(COALESCE(r.new_users, 0)) OVER (ORDER BY m.month) AS total
FROM months m
         LEFT JOIN registered r ON r.month = m.month
ORDER BY m.month;

\echo
\echo '== Social profiles by provider'
SELECT social_provider        AS provider,
       COUNT(*)               AS profiles,
       COUNT(DISTINCT user_id) AS users
FROM user_social_profiles
WHERE deleted_at IS NULL
GROUP BY social_provider
ORDER BY profiles DESC, provider;
SQL
}

stats_wip_linkedin() {
	psql_exec <<'SQL'
SET TIME ZONE 'UTC';

\echo
\echo '== LinkedIn companies (wip), total'
SELECT (SELECT COUNT(*) FROM wip_linkedin_companies)                                     AS companies,
       (SELECT COUNT(*) FROM wip_linkedin_company_request_history)                       AS requests,
       (SELECT COUNT(DISTINCT created_by) FROM wip_linkedin_company_request_history)     AS users_requesting,
       (SELECT COUNT(DISTINCT user_id) FROM wip_user_to_linkedin_companies)              AS users_with_lists,
       (SELECT COUNT(*) FROM wip_user_to_linkedin_companies WHERE active)                AS active_list_items;

\echo
\echo '== LinkedIn companies (wip) by month (UTC)'
\echo 'users_adding / companies_added: the companies added to the personal lists (the first addition of a pair)'
\echo 'users_requesting / requests: the requests to the LinkedIn API (every company is requested once)'
WITH first_month AS (
    SELECT MIN(month) AS month
    FROM (SELECT DATE_TRUNC('month', MIN(created_at)) AS month FROM wip_user_to_linkedin_companies
          UNION ALL
          SELECT DATE_TRUNC('month', MIN(created_at)) FROM wip_linkedin_company_request_history
          UNION ALL
          SELECT DATE_TRUNC('month', MIN(created_at)) FROM wip_linkedin_companies) AS firsts
),
months AS (
    SELECT GENERATE_SERIES(month, DATE_TRUNC('month', NOW()), '1 month') AS month
    FROM first_month
),
lists AS (
    SELECT DATE_TRUNC('month', created_at) AS month,
           COUNT(DISTINCT user_id)         AS users_adding,
           COUNT(*)                        AS companies_added
    FROM wip_user_to_linkedin_companies
    GROUP BY 1
),
requests AS (
    SELECT DATE_TRUNC('month', created_at)                          AS month,
           COUNT(DISTINCT created_by)                               AS users_requesting,
           COUNT(*)                                                 AS requests,
           COUNT(*) FILTER (WHERE linkedin_company_id IS NOT NULL)  AS found,
           COUNT(*) FILTER (WHERE linkedin_company_id IS NULL)      AS not_found
    FROM wip_linkedin_company_request_history
    GROUP BY 1
),
companies AS (
    SELECT DATE_TRUNC('month', created_at) AS month,
           COUNT(*)                        AS new_companies
    FROM wip_linkedin_companies
    GROUP BY 1
)
SELECT TO_CHAR(m.month, 'YYYY-MM')              AS month,
       COALESCE(l.users_adding, 0)              AS users_adding,
       COALESCE(l.companies_added, 0)           AS companies_added,
       COALESCE(r.users_requesting, 0)          AS users_requesting,
       COALESCE(r.requests, 0)                  AS requests,
       COALESCE(r.found, 0)                     AS found,
       COALESCE(r.not_found, 0)                 AS not_found,
       COALESCE(c.new_companies, 0)             AS new_companies
FROM months m
         LEFT JOIN lists l ON l.month = m.month
         LEFT JOIN requests r ON r.month = m.month
         LEFT JOIN companies c ON c.month = m.month
ORDER BY m.month;
SQL
}

usage() {
	cat >&2 <<'USAGE'
Usage: scripts/pg-stats.sh [all|size|tables|users|wip-linkedin]

  all            everything below (default)
  size           the size of the database
  tables         the tables: rows and sizes
  users          the registered users: total, by month, by provider
  wip-linkedin   the usage of the LinkedIn companies feature: total and by month
USAGE
}

main() {
	local command="${1:-all}"

	case "$command" in
	all | size | tables | users | wip-linkedin) ;;
	-h | --help | help)
		usage
		exit 0
		;;
	*)
		usage
		exit 2
		;;
	esac

	check_container

	case "$command" in
	size) stats_size ;;
	tables) stats_tables ;;
	users) stats_users ;;
	wip-linkedin) stats_wip_linkedin ;;
	all)
		stats_size
		stats_tables
		stats_users
		stats_wip_linkedin
		;;
	esac
}

# Not run when the file is sourced (for the tests)
if [ "${BASH_SOURCE[0]}" = "$0" ]; then
	main "$@"
fi
