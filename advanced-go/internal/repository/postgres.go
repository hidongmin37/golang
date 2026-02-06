package repository

import (
	"context"
	"fmt"
	"time"

	"golang/advanced-go/internal/model"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

// ============================================================
// PostgreSQL 저장소 — 실제 DB를 사용하는 Todo CRUD
// ============================================================

// 학습 포인트 — pgx vs database/sql:
//   - database/sql: Go 표준 라이브러리. 범용적이지만 PostgreSQL 고유 기능 사용 불가.
//   - pgx: PostgreSQL 전용 드라이버. LISTEN/NOTIFY, COPY, 커스텀 타입 등 지원.
//   - pgxpool: pgx 기반 커넥션 풀. database/sql의 sql.DB에 해당.
//   - 실무에서 PostgreSQL만 쓴다면 pgx가 성능과 기능 면에서 유리하다.

// PostgresRepository — pgxpool 기반 PostgreSQL 저장소
type PostgresRepository struct {
	pool *pgxpool.Pool
}

// NewPostgresRepository — 새 PostgreSQL 저장소를 생성한다.
func NewPostgresRepository(pool *pgxpool.Pool) *PostgresRepository {
	return &PostgresRepository{pool: pool}
}

func (r *PostgresRepository) Create(ctx context.Context, todo *model.Todo) error {
	query := `
		INSERT INTO todos (id, title, completed, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5)
	`
	_, err := r.pool.Exec(ctx, query,
		todo.ID, todo.Title, todo.Completed, todo.CreatedAt, todo.UpdatedAt,
	)
	if err != nil {
		return fmt.Errorf("todo 생성 실패: %w", err)
	}
	return nil
}

func (r *PostgresRepository) GetAll(ctx context.Context) ([]*model.Todo, error) {
	query := `SELECT id, title, completed, created_at, updated_at FROM todos ORDER BY created_at DESC`

	rows, err := r.pool.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("todo 목록 조회 실패: %w", err)
	}
	defer rows.Close()

	var todos []*model.Todo
	for rows.Next() {
		todo := &model.Todo{}
		if err := rows.Scan(&todo.ID, &todo.Title, &todo.Completed, &todo.CreatedAt, &todo.UpdatedAt); err != nil {
			return nil, fmt.Errorf("todo 스캔 실패: %w", err)
		}
		todos = append(todos, todo)
	}

	return todos, rows.Err()
}

func (r *PostgresRepository) GetByID(ctx context.Context, id int64) (*model.Todo, error) {
	query := `SELECT id, title, completed, created_at, updated_at FROM todos WHERE id = $1`

	todo := &model.Todo{}
	err := r.pool.QueryRow(ctx, query, id).Scan(
		&todo.ID, &todo.Title, &todo.Completed, &todo.CreatedAt, &todo.UpdatedAt,
	)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, fmt.Errorf("todo %d을(를) 찾을 수 없습니다", id)
		}
		return nil, fmt.Errorf("todo 조회 실패: %w", err)
	}

	return todo, nil
}

func (r *PostgresRepository) Update(ctx context.Context, id int64, title *string, completed *bool) (*model.Todo, error) {
	// 먼저 존재 여부를 확인하고, 부분 업데이트를 수행한다.
	// COALESCE를 사용하여 nil인 필드는 기존 값을 유지한다.
	query := `
		UPDATE todos
		SET title      = COALESCE($2, title),
		    completed  = COALESCE($3, completed),
		    updated_at = $4
		WHERE id = $1
		RETURNING id, title, completed, created_at, updated_at
	`

	now := time.Now()
	todo := &model.Todo{}
	err := r.pool.QueryRow(ctx, query, id, title, completed, now).Scan(
		&todo.ID, &todo.Title, &todo.Completed, &todo.CreatedAt, &todo.UpdatedAt,
	)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, fmt.Errorf("todo %d을(를) 찾을 수 없습니다", id)
		}
		return nil, fmt.Errorf("todo 수정 실패: %w", err)
	}

	return todo, nil
}

func (r *PostgresRepository) Delete(ctx context.Context, id int64) error {
	query := `DELETE FROM todos WHERE id = $1`

	result, err := r.pool.Exec(ctx, query, id)
	if err != nil {
		return fmt.Errorf("todo 삭제 실패: %w", err)
	}

	if result.RowsAffected() == 0 {
		return fmt.Errorf("todo %d을(를) 찾을 수 없습니다", id)
	}

	return nil
}
