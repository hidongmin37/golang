package repository

import (
	"context"

	"golang/advanced-go/internal/model"
)

// ============================================================
// TodoRepository 인터페이스 — 저장소 추상화
// ============================================================

// 학습 포인트 — 인터페이스를 통한 의존성 역전(DIP):
//   - service 계층이 구체적인 구현(메모리, PostgreSQL)이 아닌 인터페이스에 의존한다.
//   - 이렇게 하면 저장소 구현을 교체해도 service 코드를 수정할 필요가 없다.
//   - 테스트 시 mock 구현을 주입할 수도 있다.
//   - Go에서는 인터페이스를 "사용하는 쪽"에서 정의하는 것이 관례지만,
//     여러 곳에서 공유할 때는 별도 파일에 두기도 한다.

// TodoRepository — Todo CRUD 저장소 인터페이스
type TodoRepository interface {
	Create(ctx context.Context, todo *model.Todo) error
	GetAll(ctx context.Context) ([]*model.Todo, error)
	GetByID(ctx context.Context, id int64) (*model.Todo, error)
	Update(ctx context.Context, id int64, title *string, completed *bool) (*model.Todo, error)
	Delete(ctx context.Context, id int64) error
}
