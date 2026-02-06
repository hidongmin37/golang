package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"golang/advanced-go/internal/database"
	"golang/advanced-go/internal/handler"
	"golang/advanced-go/internal/middleware"
	"golang/advanced-go/internal/repository"
	"golang/advanced-go/internal/service"
	"golang/advanced-go/pkg/snowflake"
)

// ============================================================
// 메인 엔트리포인트 — 의존성 조립, 서버 설정, Graceful Shutdown
// ============================================================

func main() {
	ctx := context.Background()

	// --- Snowflake ID 생성기 ---
	idGen, err := snowflake.New(0)
	if err != nil {
		log.Fatalf("Snowflake 생성 실패: %v", err)
	}

	// --- PostgreSQL 연결 ---
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatal("DATABASE_URL 환경변수가 설정되지 않았습니다")
	}

	pool, err := database.NewPostgresPool(ctx, dsn)
	if err != nil {
		log.Fatalf("DB 연결 실패: %v", err)
	}
	defer pool.Close()

	if err := database.Migrate(ctx, pool); err != nil {
		log.Fatalf("마이그레이션 실패: %v", err)
	}

	repo := repository.NewPostgresRepository(pool)
	log.Println("저장소: PostgreSQL")

	// --- 의존성 조립 ---
	//
	// 학습 포인트 — 수동 의존성 주입(Manual DI):
	//   - Spring의 @Autowired 같은 프레임워크 DI가 아니라,
	//     main()에서 직접 생성하고 주입한다.
	//   - Go 커뮤니티에서는 이 "수동 DI" 방식을 선호한다.
	//     → 명시적이고, 디버깅이 쉽고, 매직이 없다.
	//   - 프로젝트가 커지면 google/wire 같은 코드 생성 도구를 사용하기도 한다.
	svc := service.NewTodoService(repo, idGen)
	h := handler.NewTodoHandler(svc)

	// --- 라우터 설정 ---
	mux := http.NewServeMux()
	h.RegisterRoutes(mux)

	// --- 미들웨어 체이닝 ---
	app := middleware.Chain(
		middleware.Recovery,
		middleware.Logging,
		middleware.ContentType,
	)(mux)

	// --- 서버 설정 ---
	srv := &http.Server{
		Addr:         ":8080",
		Handler:      app,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// --- Graceful Shutdown ---
	sigCtx, stop := signal.NotifyContext(ctx, os.Interrupt, syscall.SIGTERM)
	defer stop()

	go func() {
		log.Printf("서버 시작: http://localhost%s", srv.Addr)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("서버 에러: %v", err)
		}
	}()

	<-sigCtx.Done()
	log.Println("종료 시그널 수신, 서버를 정상 종료합니다...")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(shutdownCtx); err != nil {
		log.Fatalf("서버 강제 종료: %v", err)
	}

	log.Println("서버가 정상 종료되었습니다")
}
