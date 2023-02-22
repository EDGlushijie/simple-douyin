// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.0
// source: message.sql

package mydb

import (
	"context"
	"database/sql"
)

const createMessage = `-- name: CreateMessage :execresult
INSERT INTO messages(
    from_user_id,to_user_id,content,create_time
)VALUES (?,?,?,?)
`

type CreateMessageParams struct {
	FromUserID int64  `json:"from_user_id"`
	ToUserID   int64  `json:"to_user_id"`
	Content    string `json:"content"`
	CreateTime int64  `json:"create_time"`
}

func (q *Queries) CreateMessage(ctx context.Context, arg CreateMessageParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, createMessage,
		arg.FromUserID,
		arg.ToUserID,
		arg.Content,
		arg.CreateTime,
	)
}

const listMessages = `-- name: ListMessages :many
SELECT id, to_user_id, from_user_id, content, create_time FROM messages
WHERE ((from_user_id =? AND to_user_id =?) OR (from_user_id =? AND to_user_id =?)) AND create_time >= ? ORDER BY create_time
`

type ListMessagesParams struct {
	FromUserID   int64 `json:"from_user_id"`
	ToUserID     int64 `json:"to_user_id"`
	FromUserID_2 int64 `json:"from_user_id_2"`
	ToUserID_2   int64 `json:"to_user_id_2"`
	CreateTime   int64 `json:"create_time"`
}

func (q *Queries) ListMessages(ctx context.Context, arg ListMessagesParams) ([]Message, error) {
	rows, err := q.db.QueryContext(ctx, listMessages,
		arg.FromUserID,
		arg.ToUserID,
		arg.FromUserID_2,
		arg.ToUserID_2,
		arg.CreateTime,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Message
	for rows.Next() {
		var i Message
		if err := rows.Scan(
			&i.ID,
			&i.ToUserID,
			&i.FromUserID,
			&i.Content,
			&i.CreateTime,
		); err != nil {
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
