package inmemory

import (
	"sync"
	"time"
)

type InMemoryStorage struct {
	mutex     sync.Mutex
	cache     map[uint64]bool
	cleanupCh chan struct{}
}

func NewInMemoryStorage(size uint, duration int64) *InMemoryStorage {
	storage := &InMemoryStorage{
		cache:     make(map[uint64]bool, size),
		cleanupCh: make(chan struct{}),
	}

	go storage.startCleanupTimer(duration)

	return storage
}

func (i *InMemoryStorage) Set(key uint64, value bool) {
	i.mutex.Lock()
	defer i.mutex.Unlock()
	i.cache[key] = value
}

func (i *InMemoryStorage) Get(key uint64) bool {
	i.mutex.Lock()
	defer i.mutex.Unlock()
	_, ok := i.cache[key]
	return ok
}

func (i *InMemoryStorage) startCleanupTimer(duration int64) {
	ticker := time.NewTicker(time.Duration(duration) * time.Minute)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			i.cleanup()
		case <-i.cleanupCh:
			return
		}
	}
}

func (i *InMemoryStorage) cleanup() {
	i.mutex.Lock()
	defer i.mutex.Unlock()

	i.cache = make(map[uint64]bool, 50) // Очищаем кеш
}
