package database

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

// ============================================================
// PostgreSQL 연결 풀 — 데이터베이스 연결 관리
// ============================================================

// 학습 포인트 — 커넥션 풀(Connection Pool):
//   - DB 연결은 생성 비용이 크다 (TCP 핸드셰이크, 인증 등).
//   - 풀은 미리 여러 개의 연결을 만들어 두고, 요청마다 빌려주고 반납받는다.
//   - pgxpool은 자동으로 연결 수를 관리하고, 끊어진 연결을 재생성한다.
//   - 실무에서는 최대 연결 수, 유휴 시간 등을 환경에 맞게 튜닝한다.

// NewPostgresPool — PostgreSQL 커넥션 풀을 생성한다.
//
// dsn 예시: "postgres://user:password@localhost:5432/dbname?sslmode=disable"
func NewPostgresPool(ctx context.Context, dsn string) (*pgxpool.Pool, error) {
	config, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return nil, fmt.Errorf("DB 설정 파싱 실패: %w", err)
	}

	pool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		return nil, fmt.Errorf("DB 연결 풀 생성 실패: %w", err)
	}

	if err := pool.Ping(ctx); err != nil {
		pool.Close()
		return nil, fmt.Errorf("DB 연결 확인 실패: %w", err)
	}

	return pool, nil
}

// Migrate — 필요한 테이블을 생성한다 (자동 마이그레이션).
//
// 학습 포인트 — 마이그레이션:
//   - 실무에서는 golang-migrate, goose 같은 도구로 버전 관리한다.
//   - 여기서는 학습 목적으로 CREATE TABLE IF NOT EXISTS를 사용한다.
func Migrate(ctx context.Context, pool *pgxpool.Pool) error {
	query := `
		CREATE TABLE IF NOT EXISTS todos (
			id         BIGINT PRIMARY KEY,
			title      VARCHAR(200) NOT NULL,
			completed  BOOLEAN NOT NULL DEFAULT false,
			created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
			updated_at TIMESTAMPTZ NOT NULL DEFAULT now()
		);
	`
	_, err := pool.Exec(ctx, query)
	if err != nil {
		return fmt.Errorf("마이그레이션 실패: %w", err)
	}
	return nil
}
