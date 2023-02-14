// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0

package mydb

import (
	"database/sql"
)

type User struct {
	UserID        int64         `json:"user_id"`
	Name          string        `json:"name"`
	Password      string        `json:"password"`
	FollowCount   sql.NullInt64 `json:"follow_count"`
	FollowerCount sql.NullInt64 `json:"follower_count"`
}

type Video struct {
	VideoID       int64         `json:"video_id"`
	Author        int64         `json:"author"`
	PlayUrl       string        `json:"play_url"`
	CoverUrl      string        `json:"cover_url"`
	FavoriteCount sql.NullInt64 `json:"favorite_count"`
	CommentCount  sql.NullInt64 `json:"comment_count"`
	Title         string        `json:"title"`
	CreatedAt     sql.NullTime  `json:"created_at"`
}
