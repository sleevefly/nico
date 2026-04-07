package service

import (
	"sync"
	"time"
)

type lateEntry struct {
	result DspResult
	expire time.Time
}

type LateResponseStore struct {
	mu      sync.RWMutex
	data    map[string][]lateEntry
	ttl     time.Duration
	stopCh  chan struct{}
	stopWg  sync.WaitGroup
	ticker  *time.Ticker
	started bool
}

func NewLateResponseStore(ttl time.Duration) *LateResponseStore {
	if ttl <= 0 {
		ttl = 30 * time.Second
	}
	s := &LateResponseStore{
		data:   make(map[string][]lateEntry),
		ttl:    ttl,
		stopCh: make(chan struct{}),
		ticker: time.NewTicker(cleanupInterval(ttl)),
	}
	s.start()
	return s
}

func cleanupInterval(ttl time.Duration) time.Duration {
	if ttl <= time.Second {
		return 100 * time.Millisecond
	}
	iv := ttl / 2
	if iv > 5*time.Second {
		return 5 * time.Second
	}
	return iv
}

func (s *LateResponseStore) start() {
	s.mu.Lock()
	if s.started {
		s.mu.Unlock()
		return
	}
	s.started = true
	s.mu.Unlock()

	s.stopWg.Add(1)
	go func() {
		defer s.stopWg.Done()
		for {
			select {
			case <-s.ticker.C:
				s.cleanup()
			case <-s.stopCh:
				return
			}
		}
	}()
}

func (s *LateResponseStore) Stop() {
	s.mu.Lock()
	if !s.started {
		s.mu.Unlock()
		return
	}
	s.started = false
	s.mu.Unlock()

	s.ticker.Stop()
	close(s.stopCh)
	s.stopWg.Wait()
}

func (s *LateResponseStore) Put(key string, result DspResult) {
	if key == "" {
		return
	}
	entry := lateEntry{
		result: result,
		expire: time.Now().Add(s.ttl),
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	s.data[key] = append(s.data[key], entry)
}

func (s *LateResponseStore) GetByKey(key string) []DspResult {
	if key == "" {
		return nil
	}
	now := time.Now()
	s.mu.RLock()
	src := s.data[key]
	entries := make([]lateEntry, len(src))
	copy(entries, src)
	s.mu.RUnlock()

	out := make([]DspResult, 0, len(entries))
	for _, e := range entries {
		if now.Before(e.expire) {
			r := e.result
			r.IsFromCache = true
			out = append(out, r)
		}
	}
	return out
}

func (s *LateResponseStore) cleanup() {
	now := time.Now()
	s.mu.Lock()
	defer s.mu.Unlock()
	for key, entries := range s.data {
		alive := entries[:0]
		for _, e := range entries {
			if now.Before(e.expire) {
				alive = append(alive, e)
			}
		}
		if len(alive) == 0 {
			delete(s.data, key)
			continue
		}
		s.data[key] = alive
	}
}
