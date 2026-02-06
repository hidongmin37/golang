package service

import (
	"context"
	"time"

	"golang/advanced-go/internal/model"
	"golang/advanced-go/internal/repository"
	"golang/advanced-go/pkg/snowflake"
)

// ============================================================
// TodoService — 비즈니스 로직 계층
// ============================================================

// 학습 포인트 — Service 계층의 역할:
//   - Handler(HTTP) ↔ Service(비즈니스 로직) ↔ Repository(데이터 접근)
//   - Handler는 HTTP 요청/응답만 처리하고, 비즈니스 규칙은 Service에 위임한다.
//   - 이렇게 분리하면:
//     1. HTTP가 아닌 다른 인터페이스(gRPC, CLI)에서도 같은 로직을 재사용할 수 있다.
//     2. 비즈니스 로직만 독립적으로 테스트할 수 있다.
//     3. 트랜잭션 범위를 Service 메서드 단위로 관리할 수 있다.

// TodoService — Todo 비즈니스 로직을 담당하는 서비스
type TodoService struct {
	repo  repository.TodoRepository
	idGen *snowflake.Snowflake
}

// NewTodoService — 새 서비스를 생성한다.
func NewTodoService(repo repository.TodoRepository, idGen *snowflake.Snowflake) *TodoService {
	return &TodoService{
		repo:  repo,
		idGen: idGen,
	}
}

// Create — 새 Todo를 생성한다.
func (s *TodoService) Create(ctx context.Context, title string) (*model.Todo, error) {
	now := time.Now()
	todo := &model.Todo{
		ID:        s.idGen.Generate(),
		Title:     title,
		Completed: false,
		CreatedAt: now,
		UpdatedAt: now,
	}

	if err := s.repo.Create(ctx, todo); err != nil {
		return nil, err
	}
	return todo, nil
}

// GetAll — 모든 Todo를 조회한다.
func (s *TodoService) GetAll(ctx context.Context) ([]*model.Todo, error) {
	return s.repo.GetAll(ctx)
}

// GetByID — ID로 Todo를 조회한다.
func (s *TodoService) GetByID(ctx context.Context, id int64) (*model.Todo, error) {
	return s.repo.GetByID(ctx, id)
}

// Update — Todo를 부분 수정한다.
func (s *TodoService) Update(ctx context.Context, id int64, title *string, completed *bool) (*model.Todo, error) {
	return s.repo.Update(ctx, id, title, completed)
}

// Delete — Todo를 삭제한다.
func (s *TodoService) Delete(ctx context.Context, id int64) error {
	return s.repo.Delete(ctx, id)
}
