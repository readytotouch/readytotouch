// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0
// source: feature_view_stats.sql

package dbs

import (
	"context"
	"time"
)

const featureViewDailyStats = `-- name: FeatureViewDailyStats :many
SELECT days.day::DATE                    AS day,
       COALESCE(s.view_count, 0)::BIGINT AS view_count
FROM GENERATE_SERIES(
    $1::DATE,
    $2::DATE,
    '1 DAY'::INTERVAL
) AS days (day)
    LEFT JOIN feature_view_daily_stats s ON (days.day = s.created_at AND s.feature = $3)
ORDER BY days.day
`

type FeatureViewDailyStatsParams struct {
	From    time.Time
	To      time.Time
	Feature FeatureWait
}

type FeatureViewDailyStatsRow struct {
	Day       time.Time
	ViewCount int64
}

func (q *Queries) FeatureViewDailyStats(ctx context.Context, arg FeatureViewDailyStatsParams) ([]FeatureViewDailyStatsRow, error) {
	rows, err := q.query(ctx, q.featureViewDailyStatsStmt, featureViewDailyStats, arg.From, arg.To, arg.Feature)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []FeatureViewDailyStatsRow
	for rows.Next() {
		var i FeatureViewDailyStatsRow
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

const featureViewDailyStatsUpsert = `-- name: FeatureViewDailyStatsUpsert :exec
INSERT INTO feature_view_daily_stats AS fvhs (feature, created_at, view_count)
VALUES ($1, $2, 1)
ON CONFLICT (feature, created_at) DO UPDATE
    SET view_count = fvhs.view_count + 1
`

type FeatureViewDailyStatsUpsertParams struct {
	Feature   FeatureWait
	CreatedAt time.Time
}

func (q *Queries) FeatureViewDailyStatsUpsert(ctx context.Context, arg FeatureViewDailyStatsUpsertParams) error {
	_, err := q.exec(ctx, q.featureViewDailyStatsUpsertStmt, featureViewDailyStatsUpsert, arg.Feature, arg.CreatedAt)
	return err
}

const featureViewStats = `-- name: FeatureViewStats :one
SELECT view_count
FROM feature_view_stats
WHERE feature = $1
`

func (q *Queries) FeatureViewStats(ctx context.Context, feature FeatureWait) (int64, error) {
	row := q.queryRow(ctx, q.featureViewStatsStmt, featureViewStats, feature)
	var view_count int64
	err := row.Scan(&view_count)
	return view_count, err
}

const featureViewStatsUpsert = `-- name: FeatureViewStatsUpsert :exec
INSERT INTO feature_view_stats AS fvhs (feature, view_count)
VALUES ($1, 1)
ON CONFLICT (feature) DO UPDATE
    SET view_count = fvhs.view_count + 1
`

func (q *Queries) FeatureViewStatsUpsert(ctx context.Context, feature FeatureWait) error {
	_, err := q.exec(ctx, q.featureViewStatsUpsertStmt, featureViewStatsUpsert, feature)
	return err
}
