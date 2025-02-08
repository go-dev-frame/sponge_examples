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
	inventoryOperationCachePrefixKey = "inventoryOperation:"
	// InventoryOperationExpireTime expire time
	InventoryOperationExpireTime = 5 * time.Minute
)

var _ InventoryOperationCache = (*inventoryOperationCache)(nil)

// InventoryOperationCache cache interface
type InventoryOperationCache interface {
	Set(ctx context.Context, id uint64, data *model.InventoryOperation, duration time.Duration) error
	Get(ctx context.Context, id uint64) (*model.InventoryOperation, error)
	MultiGet(ctx context.Context, ids []uint64) (map[uint64]*model.InventoryOperation, error)
	MultiSet(ctx context.Context, data []*model.InventoryOperation, duration time.Duration) error
	Del(ctx context.Context, id uint64) error
	SetPlaceholder(ctx context.Context, id uint64) error
	IsPlaceholderErr(err error) bool
}

// inventoryOperationCache define a cache struct
type inventoryOperationCache struct {
	cache cache.Cache
}

// NewInventoryOperationCache new a cache
func NewInventoryOperationCache(cacheType *database.CacheType) InventoryOperationCache {
	jsonEncoding := encoding.JSONEncoding{}
	cachePrefix := ""

	cType := strings.ToLower(cacheType.CType)
	switch cType {
	case "redis":
		c := cache.NewRedisCache(cacheType.Rdb, cachePrefix, jsonEncoding, func() interface{} {
			return &model.InventoryOperation{}
		})
		return &inventoryOperationCache{cache: c}
	case "memory":
		c := cache.NewMemoryCache(cachePrefix, jsonEncoding, func() interface{} {
			return &model.InventoryOperation{}
		})
		return &inventoryOperationCache{cache: c}
	}

	return nil // no cache
}

// GetInventoryOperationCacheKey cache key
func (c *inventoryOperationCache) GetInventoryOperationCacheKey(id uint64) string {
	return inventoryOperationCachePrefixKey + utils.Uint64ToStr(id)
}

// Set write to cache
func (c *inventoryOperationCache) Set(ctx context.Context, id uint64, data *model.InventoryOperation, duration time.Duration) error {
	if data == nil || id == 0 {
		return nil
	}
	cacheKey := c.GetInventoryOperationCacheKey(id)
	err := c.cache.Set(ctx, cacheKey, data, duration)
	if err != nil {
		return err
	}
	return nil
}

// Get cache value
func (c *inventoryOperationCache) Get(ctx context.Context, id uint64) (*model.InventoryOperation, error) {
	var data *model.InventoryOperation
	cacheKey := c.GetInventoryOperationCacheKey(id)
	err := c.cache.Get(ctx, cacheKey, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// MultiSet multiple set cache
func (c *inventoryOperationCache) MultiSet(ctx context.Context, data []*model.InventoryOperation, duration time.Duration) error {
	valMap := make(map[string]interface{})
	for _, v := range data {
		cacheKey := c.GetInventoryOperationCacheKey(v.ID)
		valMap[cacheKey] = v
	}

	err := c.cache.MultiSet(ctx, valMap, duration)
	if err != nil {
		return err
	}

	return nil
}

// MultiGet multiple get cache, return key in map is id value
func (c *inventoryOperationCache) MultiGet(ctx context.Context, ids []uint64) (map[uint64]*model.InventoryOperation, error) {
	var keys []string
	for _, v := range ids {
		cacheKey := c.GetInventoryOperationCacheKey(v)
		keys = append(keys, cacheKey)
	}

	itemMap := make(map[string]*model.InventoryOperation)
	err := c.cache.MultiGet(ctx, keys, itemMap)
	if err != nil {
		return nil, err
	}

	retMap := make(map[uint64]*model.InventoryOperation)
	for _, id := range ids {
		val, ok := itemMap[c.GetInventoryOperationCacheKey(id)]
		if ok {
			retMap[id] = val
		}
	}

	return retMap, nil
}

// Del delete cache
func (c *inventoryOperationCache) Del(ctx context.Context, id uint64) error {
	cacheKey := c.GetInventoryOperationCacheKey(id)
	err := c.cache.Del(ctx, cacheKey)
	if err != nil {
		return err
	}
	return nil
}

// SetPlaceholder set placeholder value to cache
func (c *inventoryOperationCache) SetPlaceholder(ctx context.Context, id uint64) error {
	cacheKey := c.GetInventoryOperationCacheKey(id)
	return c.cache.SetCacheWithNotFound(ctx, cacheKey)
}

// IsPlaceholderErr check if cache is placeholder error
func (c *inventoryOperationCache) IsPlaceholderErr(err error) bool {
	return errors.Is(err, cache.ErrPlaceholder)
}
