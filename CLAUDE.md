# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Overview

Go 학습용 리포지토리. Go 기본 문법, 자료구조, 동시성, 웹 스크래핑 등을 주제별 패키지로 정리.

## Commands

```bash
go run main.go                        # 루트 엔트리포인트 실행
go run ./basic/accounts/              # 개별 패키지 실행 (main 함수 있는 패키지)
go build ./...                        # 전체 빌드 확인
go vet ./...                          # 정적 분석
```

## Module

- 모듈명: `golang`
- Go 1.25.6
- 외부 의존성: `goquery` (웹 스크래핑용)

## Structure

- `main.go` — 인터페이스, 타입 단언(type assertion), 에러 처리 예제
- `basic/accounts/` — 구조체 메서드, 포인터 리시버, `Stringer` 인터페이스 구현
- `basic/mydict/` — 커스텀 타입(`map` 기반), CRUD 에러 핸들링 패턴
- `basic/channel/` — goroutine, buffered channel
- `basic/url/` — HTTP 요청, goroutine+channel 병렬 처리, goquery 웹 스크래핑
- `basic/something/` — 패키지 export 기본
- `advanced-go/` — (비어있음, 향후 확장)

## Conventions

- 커밋 메시지: 한국어, 명령형
