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
	inventoryCachePrefixKey = "inventory:"
	// InventoryExpireTime expire time
	InventoryExpireTime = 5 * time.Minute
)

var _ InventoryCache = (*inventoryCache)(nil)

// InventoryCache cache interface
type InventoryCache interface {
	Set(ctx context.Context, id uint64, data *model.Inventory, duration time.Duration) error
	Get(ctx context.Context, id uint64) (*model.Inventory, error)
	MultiGet(ctx context.Context, ids []uint64) (map[uint64]*model.Inventory, error)
	MultiSet(ctx context.Context, data []*model.Inventory, duration time.Duration) error
	Del(ctx context.Context, id uint64) error
	SetPlaceholder(ctx context.Context, id uint64) error
	IsPlaceholderErr(err error) bool
}

// inventoryCache define a cache struct
type inventoryCache struct {
	cache cache.Cache
}

// NewInventoryCache new a cache
func NewInventoryCache(cacheType *database.CacheType) InventoryCache {
	jsonEncoding := encoding.JSONEncoding{}
	cachePrefix := ""

	cType := strings.ToLower(cacheType.CType)
	switch cType {
	case "redis":
		c := cache.NewRedisCache(cacheType.Rdb, cachePrefix, jsonEncoding, func() interface{} {
			return &model.Inventory{}
		})
		return &inventoryCache{cache: c}
	case "memory":
		c := cache.NewMemoryCache(cachePrefix, jsonEncoding, func() interface{} {
			return &model.Inventory{}
		})
		return &inventoryCache{cache: c}
	}

	return nil // no cache
}

// GetInventoryCacheKey cache key
func (c *inventoryCache) GetInventoryCacheKey(id uint64) string {
	return inventoryCachePrefixKey + utils.Uint64ToStr(id)
}

// Set write to cache
func (c *inventoryCache) Set(ctx context.Context, id uint64, data *model.Inventory, duration time.Duration) error {
	if data == nil || id == 0 {
		return nil
	}
	cacheKey := c.GetInventoryCacheKey(id)
	err := c.cache.Set(ctx, cacheKey, data, duration)
	if err != nil {
		return err
	}
	return nil
}

// Get cache value
func (c *inventoryCache) Get(ctx context.Context, id uint64) (*model.Inventory, error) {
	var data *model.Inventory
	cacheKey := c.GetInventoryCacheKey(id)
	err := c.cache.Get(ctx, cacheKey, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// MultiSet multiple set cache
func (c *inventoryCache) MultiSet(ctx context.Context, data []*model.Inventory, duration time.Duration) error {
	valMap := make(map[string]interface{})
	for _, v := range data {
		cacheKey := c.GetInventoryCacheKey(v.ID)
		valMap[cacheKey] = v
	}

	err := c.cache.MultiSet(ctx, valMap, duration)
	if err != nil {
		return err
	}

	return nil
}

// MultiGet multiple get cache, return key in map is id value
func (c *inventoryCache) MultiGet(ctx context.Context, ids []uint64) (map[uint64]*model.Inventory, error) {
	var keys []string
	for _, v := range ids {
		cacheKey := c.GetInventoryCacheKey(v)
		keys = append(keys, cacheKey)
	}

	itemMap := make(map[string]*model.Inventory)
	err := c.cache.MultiGet(ctx, keys, itemMap)
	if err != nil {
		return nil, err
	}

	retMap := make(map[uint64]*model.Inventory)
	for _, id := range ids {
		val, ok := itemMap[c.GetInventoryCacheKey(id)]
		if ok {
			retMap[id] = val
		}
	}

	return retMap, nil
}

// Del delete cache
func (c *inventoryCache) Del(ctx context.Context, id uint64) error {
	cacheKey := c.GetInventoryCacheKey(id)
	err := c.cache.Del(ctx, cacheKey)
	if err != nil {
		return err
	}
	return nil
}

// SetPlaceholder set placeholder value to cache
func (c *inventoryCache) SetPlaceholder(ctx context.Context, id uint64) error {
	cacheKey := c.GetInventoryCacheKey(id)
	return c.cache.SetCacheWithNotFound(ctx, cacheKey)
}

// IsPlaceholderErr check if cache is placeholder error
func (c *inventoryCache) IsPlaceholderErr(err error) bool {
	return errors.Is(err, cache.ErrPlaceholder)
}
