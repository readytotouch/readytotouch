// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0
// source: company_view_daily_stats.sql

package dbs

import (
	"context"
	"time"
)

const companyTotalViews = `-- name: CompanyTotalViews :one
SELECT SUM(view_count)                                                       AS count,
       SUM(view_count) FILTER ( WHERE created_at >= $1::DATE ) AS count_since
FROM company_view_daily_stats s
WHERE s.company_id = $2
`

type CompanyTotalViewsParams struct {
	From      time.Time
	CompanyID int64
}

type CompanyTotalViewsRow struct {
	Count      int64
	CountSince int64
}

func (q *Queries) CompanyTotalViews(ctx context.Context, arg CompanyTotalViewsParams) (CompanyTotalViewsRow, error) {
	row := q.queryRow(ctx, q.companyTotalViewsStmt, companyTotalViews, arg.From, arg.CompanyID)
	var i CompanyTotalViewsRow
	err := row.Scan(&i.Count, &i.CountSince)
	return i, err
}

const companyViewDailyStats = `-- name: CompanyViewDailyStats :many
SELECT days.day::DATE                    AS day,
       COALESCE(s.view_count, 0)::BIGINT AS view_count
FROM GENERATE_SERIES(
    $1::DATE,
    $2::DATE,
    '1 DAY'::INTERVAL
) AS days (day)
    LEFT JOIN company_view_daily_stats s ON (days.day = s.created_at AND s.company_id = $3)
ORDER BY days.day
`

type CompanyViewDailyStatsParams struct {
	From      time.Time
	To        time.Time
	CompanyID int64
}

type CompanyViewDailyStatsRow struct {
	Day       time.Time
	ViewCount int64
}

func (q *Queries) CompanyViewDailyStats(ctx context.Context, arg CompanyViewDailyStatsParams) ([]CompanyViewDailyStatsRow, error) {
	rows, err := q.query(ctx, q.companyViewDailyStatsStmt, companyViewDailyStats, arg.From, arg.To, arg.CompanyID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []CompanyViewDailyStatsRow
	for rows.Next() {
		var i CompanyViewDailyStatsRow
		if err := rows.Scan(&i.Day, &i.ViewCount); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const companyViewDailyStatsUpsert = `-- name: CompanyViewDailyStatsUpsert :exec
INSERT INTO company_view_daily_stats AS cvds (company_id, created_at, view_count)
VALUES ($1, $2, 1)
ON CONFLICT (company_id, created_at) DO UPDATE
    SET view_count = cvds.view_count + 1
`

type CompanyViewDailyStatsUpsertParams struct {
	CompanyID int64
	CreatedAt time.Time
}

func (q *Queries) CompanyViewDailyStatsUpsert(ctx context.Context, arg CompanyViewDailyStatsUpsertParams) error {
	_, err := q.exec(ctx, q.companyViewDailyStatsUpsertStmt, companyViewDailyStatsUpsert, arg.CompanyID, arg.CreatedAt)
	return err
}
