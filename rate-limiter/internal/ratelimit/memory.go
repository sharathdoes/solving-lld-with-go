package ratelimit

import (
	"context"
	"sync"
)


type MemoryLimiter struct {
	buckets map[string]*Bucket
	mu sync.RWMutex
	capacity   int64
	refillrate float64
}

func NewMemoryLimiter(capacity int64, refillRate float64) *MemoryLimiter {
	return &MemoryLimiter{
		buckets:    make(map[string]*Bucket),
		capacity:  capacity,
		refillrate: refillRate,
	}
}

func (m *MemoryLimiter) Allow (ctx context.Context, key string) (bool, error){
	m.mu.RLock()
	bucket, exists:=m.buckets[key]
	m.mu.RUnlock()
	if !exists {
		m.mu.Lock()
		bucket=NewBucket(int(m.capacity), float64(m.capacity), m.refillrate)
		m.buckets[key]=bucket
		m.mu.Unlock()
	}
	return bucket.Allow(), nil
}