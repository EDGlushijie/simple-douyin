// Code generated by sqlc. DO NOT EDIT.
// source: relations.sql

package mydb

import (
	"context"
)

const createRelation = `-- name: CreateRelation :exec
INSERT INTO relations (
  followed_id, follower_id
) VALUES (
  ?, ?
)
`

type CreateRelationParams struct {
	FollowedID int64 `json:"followed_id"`
	FollowerID int64 `json:"follower_id"`
}

func (q *Queries) CreateRelation(ctx context.Context, arg CreateRelationParams) error {
	_, err := q.db.ExecContext(ctx, createRelation, arg.FollowedID, arg.FollowerID)
	return err
}

const getFollowedCount = `-- name: GetFollowedCount :one
SELECT count(*) FROM relations
WHERE follower_id = ?
AND deleted = 0
`

func (q *Queries) GetFollowedCount(ctx context.Context, followerID int64) (int64, error) {
	row := q.db.QueryRowContext(ctx, getFollowedCount, followerID)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const getFollowedIdByFollower = `-- name: GetFollowedIdByFollower :many
SELECT followed_id FROM relations
WHERE  follower_id = ?
AND deleted = 0
`

func (q *Queries) GetFollowedIdByFollower(ctx context.Context, followerID int64) ([]int64, error) {
	rows, err := q.db.QueryContext(ctx, getFollowedIdByFollower, followerID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []int64
	for rows.Next() {
		var followed_id int64
		if err := rows.Scan(&followed_id); err != nil {
			return nil, err
		}
		items = append(items, followed_id)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getFollowerCount = `-- name: GetFollowerCount :one
SELECT count(*) FROM relations
WHERE followed_id = ?
AND deleted = 0
`

func (q *Queries) GetFollowerCount(ctx context.Context, followedID int64) (int64, error) {
	row := q.db.QueryRowContext(ctx, getFollowerCount, followedID)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const getFollowerIdByFollowed = `-- name: GetFollowerIdByFollowed :many
SELECT follower_id FROM relations
WHERE  followed_id = ?
AND deleted = 0
`

func (q *Queries) GetFollowerIdByFollowed(ctx context.Context, followedID int64) ([]int64, error) {
	rows, err := q.db.QueryContext(ctx, getFollowerIdByFollowed, followedID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []int64
	for rows.Next() {
		var follower_id int64
		if err := rows.Scan(&follower_id); err != nil {
			return nil, err
		}
		items = append(items, follower_id)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getRelationByID = `-- name: GetRelationByID :one
SELECT deleted FROM relations
WHERE followed_id = ?
AND follower_id = ?
`

type GetRelationByIDParams struct {
	FollowedID int64 `json:"followed_id"`
	FollowerID int64 `json:"follower_id"`
}

func (q *Queries) GetRelationByID(ctx context.Context, arg GetRelationByIDParams) (int32, error) {
	row := q.db.QueryRowContext(ctx, getRelationByID, arg.FollowedID, arg.FollowerID)
	var deleted int32
	err := row.Scan(&deleted)
	return deleted, err
}

const updateRelation = `-- name: UpdateRelation :exec
UPDATE relations SET deleted = ?
WHERE followed_id = ?
AND follower_id = ?
`

type UpdateRelationParams struct {
	Deleted    int32 `json:"deleted"`
	FollowedID int64 `json:"followed_id"`
	FollowerID int64 `json:"follower_id"`
}

func (q *Queries) UpdateRelation(ctx context.Context, arg UpdateRelationParams) error {
	_, err := q.db.ExecContext(ctx, updateRelation, arg.Deleted, arg.FollowedID, arg.FollowerID)
	return err
}
