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
	accountPayableCachePrefixKey = "accountPayable:"
	// AccountPayableExpireTime expire time
	AccountPayableExpireTime = 5 * time.Minute
)

var _ AccountPayableCache = (*accountPayableCache)(nil)

// AccountPayableCache cache interface
type AccountPayableCache interface {
	Set(ctx context.Context, id uint64, data *model.AccountPayable, duration time.Duration) error
	Get(ctx context.Context, id uint64) (*model.AccountPayable, error)
	MultiGet(ctx context.Context, ids []uint64) (map[uint64]*model.AccountPayable, error)
	MultiSet(ctx context.Context, data []*model.AccountPayable, duration time.Duration) error
	Del(ctx context.Context, id uint64) error
	SetPlaceholder(ctx context.Context, id uint64) error
	IsPlaceholderErr(err error) bool
}

// accountPayableCache define a cache struct
type accountPayableCache struct {
	cache cache.Cache
}

// NewAccountPayableCache new a cache
func NewAccountPayableCache(cacheType *database.CacheType) AccountPayableCache {
	jsonEncoding := encoding.JSONEncoding{}
	cachePrefix := ""

	cType := strings.ToLower(cacheType.CType)
	switch cType {
	case "redis":
		c := cache.NewRedisCache(cacheType.Rdb, cachePrefix, jsonEncoding, func() interface{} {
			return &model.AccountPayable{}
		})
		return &accountPayableCache{cache: c}
	case "memory":
		c := cache.NewMemoryCache(cachePrefix, jsonEncoding, func() interface{} {
			return &model.AccountPayable{}
		})
		return &accountPayableCache{cache: c}
	}

	return nil // no cache
}

// GetAccountPayableCacheKey cache key
func (c *accountPayableCache) GetAccountPayableCacheKey(id uint64) string {
	return accountPayableCachePrefixKey + utils.Uint64ToStr(id)
}

// Set write to cache
func (c *accountPayableCache) Set(ctx context.Context, id uint64, data *model.AccountPayable, duration time.Duration) error {
	if data == nil || id == 0 {
		return nil
	}
	cacheKey := c.GetAccountPayableCacheKey(id)
	err := c.cache.Set(ctx, cacheKey, data, duration)
	if err != nil {
		return err
	}
	return nil
}

// Get cache value
func (c *accountPayableCache) Get(ctx context.Context, id uint64) (*model.AccountPayable, error) {
	var data *model.AccountPayable
	cacheKey := c.GetAccountPayableCacheKey(id)
	err := c.cache.Get(ctx, cacheKey, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// MultiSet multiple set cache
func (c *accountPayableCache) MultiSet(ctx context.Context, data []*model.AccountPayable, duration time.Duration) error {
	valMap := make(map[string]interface{})
	for _, v := range data {
		cacheKey := c.GetAccountPayableCacheKey(v.ID)
		valMap[cacheKey] = v
	}

	err := c.cache.MultiSet(ctx, valMap, duration)
	if err != nil {
		return err
	}

	return nil
}

// MultiGet multiple get cache, return key in map is id value
func (c *accountPayableCache) MultiGet(ctx context.Context, ids []uint64) (map[uint64]*model.AccountPayable, error) {
	var keys []string
	for _, v := range ids {
		cacheKey := c.GetAccountPayableCacheKey(v)
		keys = append(keys, cacheKey)
	}

	itemMap := make(map[string]*model.AccountPayable)
	err := c.cache.MultiGet(ctx, keys, itemMap)
	if err != nil {
		return nil, err
	}

	retMap := make(map[uint64]*model.AccountPayable)
	for _, id := range ids {
		val, ok := itemMap[c.GetAccountPayableCacheKey(id)]
		if ok {
			retMap[id] = val
		}
	}

	return retMap, nil
}

// Del delete cache
func (c *accountPayableCache) Del(ctx context.Context, id uint64) error {
	cacheKey := c.GetAccountPayableCacheKey(id)
	err := c.cache.Del(ctx, cacheKey)
	if err != nil {
		return err
	}
	return nil
}

// SetPlaceholder set placeholder value to cache
func (c *accountPayableCache) SetPlaceholder(ctx context.Context, id uint64) error {
	cacheKey := c.GetAccountPayableCacheKey(id)
	return c.cache.SetCacheWithNotFound(ctx, cacheKey)
}

// IsPlaceholderErr check if cache is placeholder error
func (c *accountPayableCache) IsPlaceholderErr(err error) bool {
	return errors.Is(err, cache.ErrPlaceholder)
}
