package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"golang/advanced-go/internal/model"
	"golang/advanced-go/internal/service"
)

// ============================================================
// HTTP 핸들러 — Todo API 엔드포인트
// ============================================================

// TodoHandler — TodoService에 의존하는 핸들러 모음
type TodoHandler struct {
	service *service.TodoService
}

// NewTodoHandler — 새 핸들러를 생성한다.
func NewTodoHandler(svc *service.TodoService) *TodoHandler {
	return &TodoHandler{service: svc}
}

// RegisterRoutes — 라우트를 ServeMux에 등록한다.
//
// 학습 포인트 — 라우트 등록을 핸들러 쪽에서 관리:
//   - main.go가 모든 라우트를 알 필요 없이, 핸들러가 자신의 라우트를 등록한다.
//   - 도메인별로 핸들러 파일이 늘어나도 main.go는 깔끔하게 유지된다.
func (h *TodoHandler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /health", h.HandleHealthCheck)
	mux.HandleFunc("GET /todos", h.HandleGetTodos)
	mux.HandleFunc("GET /todos/{id}", h.HandleGetTodo)
	mux.HandleFunc("POST /todos", h.HandleCreateTodo)
	mux.HandleFunc("PATCH /todos/{id}", h.HandleUpdateTodo)
	mux.HandleFunc("DELETE /todos/{id}", h.HandleDeleteTodo)
}

// respondJSON — JSON 응답을 보내는 헬퍼
func RespondJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

// respondError — 에러 응답을 보내는 헬퍼
func RespondError(w http.ResponseWriter, status int, message string) {
	RespondJSON(w, status, model.ErrorResponse{Error: message})
}

// parseID — URL 경로에서 ID를 파싱하는 헬퍼
func parseID(r *http.Request) (int64, error) {
	return strconv.ParseInt(r.PathValue("id"), 10, 64)
}

// HandleHealthCheck — GET /health 헬스체크
func (h *TodoHandler) HandleHealthCheck(w http.ResponseWriter, r *http.Request) {
	RespondJSON(w, http.StatusOK, map[string]string{"status": "ok"})
}

// HandleGetTodos — GET /todos 전체 조회
func (h *TodoHandler) HandleGetTodos(w http.ResponseWriter, r *http.Request) {
	todos, err := h.service.GetAll(r.Context())
	if err != nil {
		RespondError(w, http.StatusInternalServerError, "목록 조회에 실패했습니다")
		return
	}

	if todos == nil {
		todos = []*model.Todo{}
	}

	RespondJSON(w, http.StatusOK, todos)
}

// HandleGetTodo — GET /todos/{id} 단건 조회
func (h *TodoHandler) HandleGetTodo(w http.ResponseWriter, r *http.Request) {
	id, err := parseID(r)
	if err != nil {
		RespondError(w, http.StatusBadRequest, "유효하지 않은 ID입니다")
		return
	}

	todo, err := h.service.GetByID(r.Context(), id)
	if err != nil {
		RespondError(w, http.StatusNotFound, err.Error())
		return
	}

	RespondJSON(w, http.StatusOK, todo)
}

// HandleCreateTodo — POST /todos 생성
func (h *TodoHandler) HandleCreateTodo(w http.ResponseWriter, r *http.Request) {
	r.Body = http.MaxBytesReader(w, r.Body, 1<<20)

	var req model.CreateTodoRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		RespondError(w, http.StatusBadRequest, "잘못된 JSON 형식입니다")
		return
	}

	if msg := req.Validate(); msg != "" {
		RespondError(w, http.StatusBadRequest, msg)
		return
	}

	todo, err := h.service.Create(r.Context(), req.Title)
	if err != nil {
		RespondError(w, http.StatusInternalServerError, "생성에 실패했습니다")
		return
	}

	RespondJSON(w, http.StatusCreated, todo)
}

// HandleUpdateTodo — PATCH /todos/{id} 부분 수정
func (h *TodoHandler) HandleUpdateTodo(w http.ResponseWriter, r *http.Request) {
	id, err := parseID(r)
	if err != nil {
		RespondError(w, http.StatusBadRequest, "유효하지 않은 ID입니다")
		return
	}

	r.Body = http.MaxBytesReader(w, r.Body, 1<<20)

	var req model.UpdateTodoRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		RespondError(w, http.StatusBadRequest, "잘못된 JSON 형식입니다")
		return
	}

	if msg := req.Validate(); msg != "" {
		RespondError(w, http.StatusBadRequest, msg)
		return
	}

	todo, err := h.service.Update(r.Context(), id, req.Title, req.Completed)
	if err != nil {
		RespondError(w, http.StatusNotFound, err.Error())
		return
	}

	RespondJSON(w, http.StatusOK, todo)
}

// HandleDeleteTodo — DELETE /todos/{id} 삭제
func (h *TodoHandler) HandleDeleteTodo(w http.ResponseWriter, r *http.Request) {
	id, err := parseID(r)
	if err != nil {
		RespondError(w, http.StatusBadRequest, "유효하지 않은 ID입니다")
		return
	}

	if err := h.service.Delete(r.Context(), id); err != nil {
		RespondError(w, http.StatusNotFound, err.Error())
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
