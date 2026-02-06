package middleware

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"golang/advanced-go/internal/model"
)

// ============================================================
// HTTP 미들웨어 — 요청 전/후 공통 처리
// ============================================================

// responseRecorder — 상태 코드를 기록하는 ResponseWriter 래퍼
type responseRecorder struct {
	http.ResponseWriter
	statusCode int
}

// WriteHeader — 상태 코드를 기록한 뒤 원본에 위임한다.
func (rec *responseRecorder) WriteHeader(code int) {
	rec.statusCode = code
	rec.ResponseWriter.WriteHeader(code)
}

// Logging — 요청/응답 로깅 미들웨어
func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		rec := &responseRecorder{
			ResponseWriter: w,
			statusCode:     http.StatusOK,
		}

		next.ServeHTTP(rec, r)

		log.Printf("%s %s → %d (%v)",
			r.Method,
			r.URL.Path,
			rec.statusCode,
			time.Since(start),
		)
	})
}

// Recovery — 패닉 복구 미들웨어
func Recovery(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("[PANIC] %s %s: %v", r.Method, r.URL.Path, err)
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(w).Encode(model.ErrorResponse{
					Error: "내부 서버 오류가 발생했습니다",
				})
			}
		}()

		next.ServeHTTP(w, r)
	})
}

// ContentType — Content-Type 헤더를 설정하는 미들웨어
func ContentType(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

// Chain — 여러 미들웨어를 하나로 합성한다.
//
// 적용 순서: Chain(A, B, C)(handler) = A(B(C(handler)))
// 요청: A → B → C → handler / 응답: handler → C → B → A
func Chain(middlewares ...func(http.Handler) http.Handler) func(http.Handler) http.Handler {
	return func(final http.Handler) http.Handler {
		for i := len(middlewares) - 1; i >= 0; i-- {
			final = middlewares[i](final)
		}
		return final
	}
}
