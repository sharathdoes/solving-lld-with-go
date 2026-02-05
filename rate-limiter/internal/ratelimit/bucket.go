package ratelimit

import (
	"sync"
	"time"
)

type Bucket struct {
	capacity   int64
	tokens     float64
	refillrate float64
	lastrefill time.Time
	mu         sync.Mutex
}

func NewBucket(capacity int, tokens float64, refillrate float64) *Bucket {
	return &Bucket{capacity: int64(capacity), tokens: tokens, refillrate: refillrate}
}

func (b *Bucket) Allow() bool {
	b.mu.Lock()
	defer b.mu.Unlock()
	now := time.Now()
	elapsed := now.Sub(b.lastrefill).Seconds()
	b.tokens += elapsed * b.refillrate
	if b.tokens > float64(b.capacity) {
		b.tokens = float64(b.capacity)
	}

	if b.tokens > 1 {
		b.tokens-=1
		return true
	} else {
		return false
	}

}
