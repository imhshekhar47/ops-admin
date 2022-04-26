package service

/*
import (
	"fmt"
	"sync"
	"time"

	"github.com/sirupsen/logrus"
)

type Cache interface {
	Read(string) (interface{}, error)
	Write(string, interface{}, time.Duration) error
	Scan() error
}

type CacheService interface {
	Read(string) (interface{}, error)
	Write(string, interface{}, time.Duration) error
	Scan() error
}

// ---------- Implementation -----------
type CacheEntry struct {
	Key        string
	Data       interface{}
	CreatedOn  time.Time
	AccessedOn time.Time
	Validity   time.Duration
}

func (c *CacheEntry) IsValid() bool {
	return !time.Now().After(c.CreatedOn.Add(c.Validity))
}

type InmemoryCache struct {
	log   *logrus.Entry
	lock  sync.RWMutex
	store map[string]*CacheEntry
}

func NewInmemoryCache(logger *logrus.Logger) *InmemoryCache {

	return &InmemoryCache{
		log:   logger.WithField("origin", "InmemoryCache"),
		lock:  sync.RWMutex{},
		store: make(map[string]*CacheEntry),
	}
}

func (s *InmemoryCache) Read(key string) (interface{}, error) {
	s.log.Traceln("entry: Read(%s)", key)
	s.lock.RLock()
	defer s.lock.RUnlock()
	entry, found := s.store[key]
	if !found {
		s.log.Traceln("exit: Read(), Not found", key)
		return nil, fmt.Errorf("%s is Not in cache", key)
	}

	if !entry.IsValid() {
		s.log.Traceln("exit: Read(), Expired", key)
		return nil, fmt.Errorf("%s is Expired", key)
	} else {
		entry.AccessedOn = time.Now()
	}

	s.log.Println("exit: Read()")
	return entry.Data, nil
}

func (s *InmemoryCache) Write(key string, data interface{}, validity time.Duration) error {
	s.log.Traceln("entry: Write(%s)", key)
	s.lock.Lock()
	defer s.lock.Unlock()
	s.store[key] = &CacheEntry{
		Key:       key,
		Data:      data,
		CreatedOn: time.Now(),
		Validity:  validity,
	}

	s.log.Println("exit: Write()")
	return nil
}

func (s *InmemoryCache) Scan() error {
	s.log.Traceln("entry: Scan()")
	staleKeys := make([]string, 0)

	for key, entry := range s.store {
		if !entry.IsValid() {
			staleKeys = append(staleKeys, key)
		}
	}

	s.lock.Lock()
	for _, key := range staleKeys {
		s.log.Traceln("Deleting  entry '%s'", key)
		delete(s.store, key)
	}
	s.lock.Unlock()

	s.log.Traceln("exit: Scan()")
	return nil
}

type AdminCacheService struct {
	cache *Cache
}

func NewAdminCacheService(cache *Cache) *AdminCacheService {
	return &AdminCacheService{
		cache: cache,
	}
}

func (s *AdminCacheService) Read(key string) (interface{}, error) {
	return (*s.cache).Read(key)
}

func (s *AdminCacheService) Write(key string, data interface{}, validity time.Duration) error {
	return (*s.cache).Write(key, data, validity)
}

func (s *AdminCacheService) Scan() error {
	return (*s.cache).Scan()
}
*/
