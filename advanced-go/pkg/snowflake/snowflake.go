package snowflake

import (
	"fmt"
	"sync"
	"time"
)

// ============================================================
// Snowflake ID 생성기 — 분산 환경에서 충돌 없는 고유 ID
// ============================================================

// 학습 포인트 — Snowflake란?
//   - Twitter(현 X)가 2010년에 만든 분산 ID 생성 알고리즘이다.
//   - 64비트 int 하나에 "시간 + 서버번호 + 순번"을 비트로 나눠 담는다.
//   - 각 서버가 독립적으로 ID를 생성해도, 서버번호가 다르면 절대 충돌하지 않는다.
//   - DB 왕복 없이 애플리케이션에서 바로 생성 → 매우 빠르다.
//
// 64비트 구조:
//
//	0 | 00000000 00000000 00000000 00000000 00000000 0 | 00000000 00 | 00000000 0000
//	↑   ↑ 41비트: 타임스탬프 (밀리초)                    ↑ 10비트     ↑ 12비트
//	부호  약 69년간 사용 가능                             노드 ID      시퀀스
//	(항상 0)                                            (0~1023)     (0~4095/ms)

const (
	epoch int64 = 1704067200000 // 2024-01-01 00:00:00 UTC

	nodeBits     = 10
	sequenceBits = 12

	nodeShift      = sequenceBits
	timestampShift = nodeBits + sequenceBits

	MaxNode     = (1 << nodeBits) - 1     // 1023
	maxSequence = (1 << sequenceBits) - 1 // 4095
)

// Snowflake — Snowflake ID 생성기
type Snowflake struct {
	mu       sync.Mutex
	nodeID   int64
	lastTime int64
	sequence int64
}

// New — 새 Snowflake 생성기를 만든다.
func New(nodeID int64) (*Snowflake, error) {
	if nodeID < 0 || nodeID > MaxNode {
		return nil, fmt.Errorf("nodeID는 0~%d 범위여야 합니다 (입력: %d)", MaxNode, nodeID)
	}
	return &Snowflake{nodeID: nodeID}, nil
}

// Generate — 새 Snowflake ID를 생성한다.
func (s *Snowflake) Generate() int64 {
	s.mu.Lock()
	defer s.mu.Unlock()

	now := time.Now().UnixMilli() - epoch

	if now == s.lastTime {
		s.sequence = (s.sequence + 1) & maxSequence
		if s.sequence == 0 {
			for now <= s.lastTime {
				now = time.Now().UnixMilli() - epoch
			}
		}
	} else {
		s.sequence = 0
	}

	s.lastTime = now

	return (now << timestampShift) | (s.nodeID << nodeShift) | s.sequence
}

// Decompose — Snowflake ID를 각 구성요소로 분해한다 (디버깅용).
func Decompose(id int64) (timestamp, nodeID, sequence int64) {
	timestamp = (id >> timestampShift) + epoch
	nodeID = (id >> nodeShift) & MaxNode
	sequence = id & maxSequence
	return
}
