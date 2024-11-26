// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: worker.queries.sql

package database

import (
	"context"
	"database/sql"
	"encoding/json"

	"github.com/google/uuid"
	"github.com/sqlc-dev/pqtype"
)

const updateTaskProcessing = `-- name: UpdateTaskProcessing :one
UPDATE tasks 
SET status = 'processing', 
  started_at = $1 
WHERE id = (
  SELECT id 
  FROM tasks 
  WHERE status = 'pending' 
  ORDER BY created_at 
  FOR UPDATE SKIP LOCKED 
  LIMIT 1
)
RETURNING id, object_id, input
`

type UpdateTaskProcessingRow struct {
	ID       *uuid.UUID      `json:"id"`
	ObjectID *uuid.UUID      `json:"object_id"`
	Input    json.RawMessage `json:"input"`
}

func (q *Queries) UpdateTaskProcessing(ctx context.Context, startedAt sql.NullTime) (UpdateTaskProcessingRow, error) {
	row := q.queryRow(ctx, q.updateTaskProcessingStmt, updateTaskProcessing, startedAt)
	var i UpdateTaskProcessingRow
	err := row.Scan(&i.ID, &i.ObjectID, &i.Input)
	return i, err
}

const updateTaskStatus = `-- name: UpdateTaskStatus :exec
UPDATE tasks 
SET status = $1, 
  output = $2, 
  error = $3, 
  completed_at = $4 
WHERE id = $5
`

type UpdateTaskStatusParams struct {
	Status      string                `json:"status"`
	Output      pqtype.NullRawMessage `json:"output"`
	Error       sql.NullString        `json:"error"`
	CompletedAt sql.NullTime          `json:"completed_at"`
	ID          *uuid.UUID            `json:"id"`
}

func (q *Queries) UpdateTaskStatus(ctx context.Context, arg UpdateTaskStatusParams) error {
	_, err := q.exec(ctx, q.updateTaskStatusStmt, updateTaskStatus,
		arg.Status,
		arg.Output,
		arg.Error,
		arg.CompletedAt,
		arg.ID,
	)
	return err
}
