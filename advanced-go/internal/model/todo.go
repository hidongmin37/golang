package model

import "time"

// ============================================================
// Todo 도메인 모델 및 요청/응답 DTO
// ============================================================

// Todo — 핵심 도메인 구조체
type Todo struct {
	ID        int64     `json:"id"`
	Title     string    `json:"title"`
	Completed bool      `json:"completed"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// CreateTodoRequest — Todo 생성 요청 DTO
type CreateTodoRequest struct {
	Title string `json:"title"`
}

// Validate — 생성 요청의 유효성을 검사한다.
func (r *CreateTodoRequest) Validate() string {
	if r.Title == "" {
		return "title은 필수입니다"
	}
	if len(r.Title) > 200 {
		return "title은 200자 이하여야 합니다"
	}
	return ""
}

// UpdateTodoRequest — Todo 수정 요청 DTO (부분 업데이트용)
//
// 포인터 필드를 사용하여 JSON에서 필드 미전송(nil)과 zero value를 구분한다.
type UpdateTodoRequest struct {
	Title     *string `json:"title"`
	Completed *bool   `json:"completed"`
}

// Validate — 수정 요청의 유효성을 검사한다.
func (r *UpdateTodoRequest) Validate() string {
	if r.Title == nil && r.Completed == nil {
		return "수정할 필드가 최소 하나 필요합니다"
	}
	if r.Title != nil && *r.Title == "" {
		return "title은 빈 문자열일 수 없습니다"
	}
	if r.Title != nil && len(*r.Title) > 200 {
		return "title은 200자 이하여야 합니다"
	}
	return ""
}

// ErrorResponse — 에러 응답 DTO
type ErrorResponse struct {
	Error string `json:"error"`
}
