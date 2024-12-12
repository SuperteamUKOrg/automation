// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package database

import (
	"context"
	"database/sql"
	"fmt"
)

type DBTX interface {
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

func New(db DBTX) *Queries {
	return &Queries{db: db}
}

func Prepare(ctx context.Context, db DBTX) (*Queries, error) {
	q := Queries{db: db}
	var err error
	if q.countObjectsStmt, err = db.PrepareContext(ctx, countObjects); err != nil {
		return nil, fmt.Errorf("error preparing query CountObjects: %w", err)
	}
	if q.countTasksStmt, err = db.PrepareContext(ctx, countTasks); err != nil {
		return nil, fmt.Errorf("error preparing query CountTasks: %w", err)
	}
	if q.createObjectStmt, err = db.PrepareContext(ctx, createObject); err != nil {
		return nil, fmt.Errorf("error preparing query CreateObject: %w", err)
	}
	if q.createScanLogStmt, err = db.PrepareContext(ctx, createScanLog); err != nil {
		return nil, fmt.Errorf("error preparing query CreateScanLog: %w", err)
	}
	if q.createTaskStmt, err = db.PrepareContext(ctx, createTask); err != nil {
		return nil, fmt.Errorf("error preparing query CreateTask: %w", err)
	}
	if q.getLatestScanTimeStmt, err = db.PrepareContext(ctx, getLatestScanTime); err != nil {
		return nil, fmt.Errorf("error preparing query GetLatestScanTime: %w", err)
	}
	if q.getObjectStmt, err = db.PrepareContext(ctx, getObject); err != nil {
		return nil, fmt.Errorf("error preparing query GetObject: %w", err)
	}
	if q.getStaleObjectsStmt, err = db.PrepareContext(ctx, getStaleObjects); err != nil {
		return nil, fmt.Errorf("error preparing query GetStaleObjects: %w", err)
	}
	if q.healthCheckStmt, err = db.PrepareContext(ctx, healthCheck); err != nil {
		return nil, fmt.Errorf("error preparing query HealthCheck: %w", err)
	}
	if q.listObjectsStmt, err = db.PrepareContext(ctx, listObjects); err != nil {
		return nil, fmt.Errorf("error preparing query ListObjects: %w", err)
	}
	if q.listTasksStmt, err = db.PrepareContext(ctx, listTasks); err != nil {
		return nil, fmt.Errorf("error preparing query ListTasks: %w", err)
	}
	if q.objectsSyncLast60daysStmt, err = db.PrepareContext(ctx, objectsSyncLast60days); err != nil {
		return nil, fmt.Errorf("error preparing query ObjectsSyncLast60days: %w", err)
	}
	if q.updateObjectLastSyncedAtStmt, err = db.PrepareContext(ctx, updateObjectLastSyncedAt); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateObjectLastSyncedAt: %w", err)
	}
	if q.updateTaskProcessingStmt, err = db.PrepareContext(ctx, updateTaskProcessing); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateTaskProcessing: %w", err)
	}
	if q.updateTaskStatusStmt, err = db.PrepareContext(ctx, updateTaskStatus); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateTaskStatus: %w", err)
	}
	return &q, nil
}

func (q *Queries) Close() error {
	var err error
	if q.countObjectsStmt != nil {
		if cerr := q.countObjectsStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing countObjectsStmt: %w", cerr)
		}
	}
	if q.countTasksStmt != nil {
		if cerr := q.countTasksStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing countTasksStmt: %w", cerr)
		}
	}
	if q.createObjectStmt != nil {
		if cerr := q.createObjectStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createObjectStmt: %w", cerr)
		}
	}
	if q.createScanLogStmt != nil {
		if cerr := q.createScanLogStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createScanLogStmt: %w", cerr)
		}
	}
	if q.createTaskStmt != nil {
		if cerr := q.createTaskStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createTaskStmt: %w", cerr)
		}
	}
	if q.getLatestScanTimeStmt != nil {
		if cerr := q.getLatestScanTimeStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getLatestScanTimeStmt: %w", cerr)
		}
	}
	if q.getObjectStmt != nil {
		if cerr := q.getObjectStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getObjectStmt: %w", cerr)
		}
	}
	if q.getStaleObjectsStmt != nil {
		if cerr := q.getStaleObjectsStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getStaleObjectsStmt: %w", cerr)
		}
	}
	if q.healthCheckStmt != nil {
		if cerr := q.healthCheckStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing healthCheckStmt: %w", cerr)
		}
	}
	if q.listObjectsStmt != nil {
		if cerr := q.listObjectsStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing listObjectsStmt: %w", cerr)
		}
	}
	if q.listTasksStmt != nil {
		if cerr := q.listTasksStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing listTasksStmt: %w", cerr)
		}
	}
	if q.objectsSyncLast60daysStmt != nil {
		if cerr := q.objectsSyncLast60daysStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing objectsSyncLast60daysStmt: %w", cerr)
		}
	}
	if q.updateObjectLastSyncedAtStmt != nil {
		if cerr := q.updateObjectLastSyncedAtStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateObjectLastSyncedAtStmt: %w", cerr)
		}
	}
	if q.updateTaskProcessingStmt != nil {
		if cerr := q.updateTaskProcessingStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateTaskProcessingStmt: %w", cerr)
		}
	}
	if q.updateTaskStatusStmt != nil {
		if cerr := q.updateTaskStatusStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateTaskStatusStmt: %w", cerr)
		}
	}
	return err
}

func (q *Queries) exec(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (sql.Result, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).ExecContext(ctx, args...)
	case stmt != nil:
		return stmt.ExecContext(ctx, args...)
	default:
		return q.db.ExecContext(ctx, query, args...)
	}
}

func (q *Queries) query(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (*sql.Rows, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryContext(ctx, args...)
	default:
		return q.db.QueryContext(ctx, query, args...)
	}
}

func (q *Queries) queryRow(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) *sql.Row {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryRowContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryRowContext(ctx, args...)
	default:
		return q.db.QueryRowContext(ctx, query, args...)
	}
}

type Queries struct {
	db                           DBTX
	tx                           *sql.Tx
	countObjectsStmt             *sql.Stmt
	countTasksStmt               *sql.Stmt
	createObjectStmt             *sql.Stmt
	createScanLogStmt            *sql.Stmt
	createTaskStmt               *sql.Stmt
	getLatestScanTimeStmt        *sql.Stmt
	getObjectStmt                *sql.Stmt
	getStaleObjectsStmt          *sql.Stmt
	healthCheckStmt              *sql.Stmt
	listObjectsStmt              *sql.Stmt
	listTasksStmt                *sql.Stmt
	objectsSyncLast60daysStmt    *sql.Stmt
	updateObjectLastSyncedAtStmt *sql.Stmt
	updateTaskProcessingStmt     *sql.Stmt
	updateTaskStatusStmt         *sql.Stmt
}

func (q *Queries) WithTx(tx *sql.Tx) *Queries {
	return &Queries{
		db:                           tx,
		tx:                           tx,
		countObjectsStmt:             q.countObjectsStmt,
		countTasksStmt:               q.countTasksStmt,
		createObjectStmt:             q.createObjectStmt,
		createScanLogStmt:            q.createScanLogStmt,
		createTaskStmt:               q.createTaskStmt,
		getLatestScanTimeStmt:        q.getLatestScanTimeStmt,
		getObjectStmt:                q.getObjectStmt,
		getStaleObjectsStmt:          q.getStaleObjectsStmt,
		healthCheckStmt:              q.healthCheckStmt,
		listObjectsStmt:              q.listObjectsStmt,
		listTasksStmt:                q.listTasksStmt,
		objectsSyncLast60daysStmt:    q.objectsSyncLast60daysStmt,
		updateObjectLastSyncedAtStmt: q.updateObjectLastSyncedAtStmt,
		updateTaskProcessingStmt:     q.updateTaskProcessingStmt,
		updateTaskStatusStmt:         q.updateTaskStatusStmt,
	}
}
