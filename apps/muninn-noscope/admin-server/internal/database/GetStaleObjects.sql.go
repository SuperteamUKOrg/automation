// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: GetStaleObjects.sql

package database

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
)

const getStaleObjects = `-- name: GetStaleObjects :many
SELECT o.id 
FROM objects o
WHERE (o.last_synced_at < $1 OR o.last_synced_at IS NULL)
AND NOT EXISTS (
  SELECT 1 
  FROM tasks t 
  WHERE t.object_id = o.id 
  AND t.status = 'pending' OR t.status = 'processing'
)
`

func (q *Queries) GetStaleObjects(ctx context.Context, lastSyncedAt sql.NullTime) ([]*uuid.UUID, error) {
	rows, err := q.query(ctx, q.getStaleObjectsStmt, getStaleObjects, lastSyncedAt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []*uuid.UUID
	for rows.Next() {
		var id *uuid.UUID
		if err := rows.Scan(&id); err != nil {
			return nil, err
		}
		items = append(items, id)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}