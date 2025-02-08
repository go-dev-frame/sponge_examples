package cache

import (
	"context"
	"errors"
	"strings"
	"time"

	"github.com/go-dev-frame/sponge/pkg/cache"
	"github.com/go-dev-frame/sponge/pkg/encoding"
	"github.com/go-dev-frame/sponge/pkg/utils"

	"store/internal/database"
	"store/internal/model"
)

const (
	// cache prefix key, must end with a colon
	storeCachePrefixKey = "store:"
	// StoreExpireTime expire time
	StoreExpireTime = 5 * time.Minute
)

var _ StoreCache = (*storeCache)(nil)

// StoreCache cache interface
type StoreCache interface {
	Set(ctx context.Context, id uint64, data *model.Store, duration time.Duration) error
	Get(ctx context.Context, id uint64) (*model.Store, error)
	MultiGet(ctx context.Context, ids []uint64) (map[uint64]*model.Store, error)
	MultiSet(ctx context.Context, data []*model.Store, duration time.Duration) error
	Del(ctx context.Context, id uint64) error
	SetPlaceholder(ctx context.Context, id uint64) error
	IsPlaceholderErr(err error) bool
}

// storeCache define a cache struct
type storeCache struct {
	cache cache.Cache
}

// NewStoreCache new a cache
func NewStoreCache(cacheType *database.CacheType) StoreCache {
	jsonEncoding := encoding.JSONEncoding{}
	cachePrefix := ""

	cType := strings.ToLower(cacheType.CType)
	switch cType {
	case "redis":
		c := cache.NewRedisCache(cacheType.Rdb, cachePrefix, jsonEncoding, func() interface{} {
			return &model.Store{}
		})
		return &storeCache{cache: c}
	case "memory":
		c := cache.NewMemoryCache(cachePrefix, jsonEncoding, func() interface{} {
			return &model.Store{}
		})
		return &storeCache{cache: c}
	}

	return nil // no cache
}

// GetStoreCacheKey cache key
func (c *storeCache) GetStoreCacheKey(id uint64) string {
	return storeCachePrefixKey + utils.Uint64ToStr(id)
}

// Set write to cache
func (c *storeCache) Set(ctx context.Context, id uint64, data *model.Store, duration time.Duration) error {
	if data == nil || id == 0 {
		return nil
	}
	cacheKey := c.GetStoreCacheKey(id)
	err := c.cache.Set(ctx, cacheKey, data, duration)
	if err != nil {
		return err
	}
	return nil
}

// Get cache value
func (c *storeCache) Get(ctx context.Context, id uint64) (*model.Store, error) {
	var data *model.Store
	cacheKey := c.GetStoreCacheKey(id)
	err := c.cache.Get(ctx, cacheKey, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// MultiSet multiple set cache
func (c *storeCache) MultiSet(ctx context.Context, data []*model.Store, duration time.Duration) error {
	valMap := make(map[string]interface{})
	for _, v := range data {
		cacheKey := c.GetStoreCacheKey(v.ID)
		valMap[cacheKey] = v
	}

	err := c.cache.MultiSet(ctx, valMap, duration)
	if err != nil {
		return err
	}

	return nil
}

// MultiGet multiple get cache, return key in map is id value
func (c *storeCache) MultiGet(ctx context.Context, ids []uint64) (map[uint64]*model.Store, error) {
	var keys []string
	for _, v := range ids {
		cacheKey := c.GetStoreCacheKey(v)
		keys = append(keys, cacheKey)
	}

	itemMap := make(map[string]*model.Store)
	err := c.cache.MultiGet(ctx, keys, itemMap)
	if err != nil {
		return nil, err
	}

	retMap := make(map[uint64]*model.Store)
	for _, id := range ids {
		val, ok := itemMap[c.GetStoreCacheKey(id)]
		if ok {
			retMap[id] = val
		}
	}

	return retMap, nil
}

// Del delete cache
func (c *storeCache) Del(ctx context.Context, id uint64) error {
	cacheKey := c.GetStoreCacheKey(id)
	err := c.cache.Del(ctx, cacheKey)
	if err != nil {
		return err
	}
	return nil
}

// SetPlaceholder set placeholder value to cache
func (c *storeCache) SetPlaceholder(ctx context.Context, id uint64) error {
	cacheKey := c.GetStoreCacheKey(id)
	return c.cache.SetCacheWithNotFound(ctx, cacheKey)
}

// IsPlaceholderErr check if cache is placeholder error
func (c *storeCache) IsPlaceholderErr(err error) bool {
	return errors.Is(err, cache.ErrPlaceholder)
}
