package cache

import (
	"context"
	"errors"
	"strings"
	"time"

	"github.com/go-dev-frame/sponge/pkg/cache"
	"github.com/go-dev-frame/sponge/pkg/encoding"

	"store/internal/database"
	"store/internal/model"
)

const (
	// cache prefix key, must end with a colon
	inventoryCheckCachePrefixKey = "inventoryCheck:"
	// InventoryCheckExpireTime expire time
	InventoryCheckExpireTime = 5 * time.Minute
)

var _ InventoryCheckCache = (*inventoryCheckCache)(nil)

// InventoryCheckCache cache interface
type InventoryCheckCache interface {
	Set(ctx context.Context, id string, data *model.InventoryCheck, duration time.Duration) error
	Get(ctx context.Context, id string) (*model.InventoryCheck, error)
	MultiGet(ctx context.Context, ids []string) (map[string]*model.InventoryCheck, error)
	MultiSet(ctx context.Context, data []*model.InventoryCheck, duration time.Duration) error
	Del(ctx context.Context, id string) error
	SetPlaceholder(ctx context.Context, id string) error
	IsPlaceholderErr(err error) bool
}

// inventoryCheckCache define a cache struct
type inventoryCheckCache struct {
	cache cache.Cache
}

// NewInventoryCheckCache new a cache
func NewInventoryCheckCache(cacheType *database.CacheType) InventoryCheckCache {
	jsonEncoding := encoding.JSONEncoding{}
	cachePrefix := ""

	cType := strings.ToLower(cacheType.CType)
	switch cType {
	case "redis":
		c := cache.NewRedisCache(cacheType.Rdb, cachePrefix, jsonEncoding, func() interface{} {
			return &model.InventoryCheck{}
		})
		return &inventoryCheckCache{cache: c}
	case "memory":
		c := cache.NewMemoryCache(cachePrefix, jsonEncoding, func() interface{} {
			return &model.InventoryCheck{}
		})
		return &inventoryCheckCache{cache: c}
	}

	return nil // no cache
}

// GetInventoryCheckCacheKey cache key
func (c *inventoryCheckCache) GetInventoryCheckCacheKey(id string) string {
	return inventoryCheckCachePrefixKey + id
}

// Set write to cache
func (c *inventoryCheckCache) Set(ctx context.Context, id string, data *model.InventoryCheck, duration time.Duration) error {
	if data == nil {
		return nil
	}
	cacheKey := c.GetInventoryCheckCacheKey(id)
	err := c.cache.Set(ctx, cacheKey, data, duration)
	if err != nil {
		return err
	}
	return nil
}

// Get cache value
func (c *inventoryCheckCache) Get(ctx context.Context, id string) (*model.InventoryCheck, error) {
	var data *model.InventoryCheck
	cacheKey := c.GetInventoryCheckCacheKey(id)
	err := c.cache.Get(ctx, cacheKey, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// MultiSet multiple set cache
func (c *inventoryCheckCache) MultiSet(ctx context.Context, data []*model.InventoryCheck, duration time.Duration) error {
	valMap := make(map[string]interface{})
	for _, v := range data {
		cacheKey := c.GetInventoryCheckCacheKey(v.ID)
		valMap[cacheKey] = v
	}

	err := c.cache.MultiSet(ctx, valMap, duration)
	if err != nil {
		return err
	}

	return nil
}

// MultiGet multiple get cache, return key in map is id value
func (c *inventoryCheckCache) MultiGet(ctx context.Context, ids []string) (map[string]*model.InventoryCheck, error) {
	var keys []string
	for _, v := range ids {
		cacheKey := c.GetInventoryCheckCacheKey(v)
		keys = append(keys, cacheKey)
	}

	itemMap := make(map[string]*model.InventoryCheck)
	err := c.cache.MultiGet(ctx, keys, itemMap)
	if err != nil {
		return nil, err
	}

	retMap := make(map[string]*model.InventoryCheck)
	for _, id := range ids {
		val, ok := itemMap[c.GetInventoryCheckCacheKey(id)]
		if ok {
			retMap[id] = val
		}
	}

	return retMap, nil
}

// Del delete cache
func (c *inventoryCheckCache) Del(ctx context.Context, id string) error {
	cacheKey := c.GetInventoryCheckCacheKey(id)
	err := c.cache.Del(ctx, cacheKey)
	if err != nil {
		return err
	}
	return nil
}

// SetPlaceholder set placeholder value to cache
func (c *inventoryCheckCache) SetPlaceholder(ctx context.Context, id string) error {
	cacheKey := c.GetInventoryCheckCacheKey(id)
	return c.cache.SetCacheWithNotFound(ctx, cacheKey)
}

// IsPlaceholderErr check if cache is placeholder error
func (c *inventoryCheckCache) IsPlaceholderErr(err error) bool {
	return errors.Is(err, cache.ErrPlaceholder)
}
