package client

import (
	"time"

	"github.com/maypok86/otter"
)

// Interface Check
var _ Cache = (*otterCacheStore)(nil)

// otterCacheStore implements Cache using Otter.
type otterCacheStore struct {
	cache otter.Cache[string, Compliance]
}

func (s *otterCacheStore) Get(key string) (Compliance, bool) {
	return s.cache.Get(key)
}

func (s *otterCacheStore) Set(key string, value Compliance) error {
	s.cache.Set(key, value)
	return nil
}

func (s *otterCacheStore) Delete(key string) error {
	s.cache.Delete(key)
	return nil
}

func NewOtterStore(ttl time.Duration, maxEntries int) (Cache, error) {
	cache, err := otter.MustBuilder[string, Compliance](maxEntries).
		CollectStats().
		WithTTL(ttl).
		Build()
	return &otterCacheStore{cache: cache}, err
}
