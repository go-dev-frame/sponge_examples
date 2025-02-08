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
	purchaseOrderCachePrefixKey = "purchaseOrder:"
	// PurchaseOrderExpireTime expire time
	PurchaseOrderExpireTime = 5 * time.Minute
)

var _ PurchaseOrderCache = (*purchaseOrderCache)(nil)

// PurchaseOrderCache cache interface
type PurchaseOrderCache interface {
	Set(ctx context.Context, id string, data *model.PurchaseOrder, duration time.Duration) error
	Get(ctx context.Context, id string) (*model.PurchaseOrder, error)
	MultiGet(ctx context.Context, ids []string) (map[string]*model.PurchaseOrder, error)
	MultiSet(ctx context.Context, data []*model.PurchaseOrder, duration time.Duration) error
	Del(ctx context.Context, id string) error
	SetPlaceholder(ctx context.Context, id string) error
	IsPlaceholderErr(err error) bool
}

// purchaseOrderCache define a cache struct
type purchaseOrderCache struct {
	cache cache.Cache
}

// NewPurchaseOrderCache new a cache
func NewPurchaseOrderCache(cacheType *database.CacheType) PurchaseOrderCache {
	jsonEncoding := encoding.JSONEncoding{}
	cachePrefix := ""

	cType := strings.ToLower(cacheType.CType)
	switch cType {
	case "redis":
		c := cache.NewRedisCache(cacheType.Rdb, cachePrefix, jsonEncoding, func() interface{} {
			return &model.PurchaseOrder{}
		})
		return &purchaseOrderCache{cache: c}
	case "memory":
		c := cache.NewMemoryCache(cachePrefix, jsonEncoding, func() interface{} {
			return &model.PurchaseOrder{}
		})
		return &purchaseOrderCache{cache: c}
	}

	return nil // no cache
}

// GetPurchaseOrderCacheKey cache key
func (c *purchaseOrderCache) GetPurchaseOrderCacheKey(id string) string {
	return purchaseOrderCachePrefixKey + id
}

// Set write to cache
func (c *purchaseOrderCache) Set(ctx context.Context, id string, data *model.PurchaseOrder, duration time.Duration) error {
	if data == nil {
		return nil
	}
	cacheKey := c.GetPurchaseOrderCacheKey(id)
	err := c.cache.Set(ctx, cacheKey, data, duration)
	if err != nil {
		return err
	}
	return nil
}

// Get cache value
func (c *purchaseOrderCache) Get(ctx context.Context, id string) (*model.PurchaseOrder, error) {
	var data *model.PurchaseOrder
	cacheKey := c.GetPurchaseOrderCacheKey(id)
	err := c.cache.Get(ctx, cacheKey, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// MultiSet multiple set cache
func (c *purchaseOrderCache) MultiSet(ctx context.Context, data []*model.PurchaseOrder, duration time.Duration) error {
	valMap := make(map[string]interface{})
	for _, v := range data {
		cacheKey := c.GetPurchaseOrderCacheKey(v.ID)
		valMap[cacheKey] = v
	}

	err := c.cache.MultiSet(ctx, valMap, duration)
	if err != nil {
		return err
	}

	return nil
}

// MultiGet multiple get cache, return key in map is id value
func (c *purchaseOrderCache) MultiGet(ctx context.Context, ids []string) (map[string]*model.PurchaseOrder, error) {
	var keys []string
	for _, v := range ids {
		cacheKey := c.GetPurchaseOrderCacheKey(v)
		keys = append(keys, cacheKey)
	}

	itemMap := make(map[string]*model.PurchaseOrder)
	err := c.cache.MultiGet(ctx, keys, itemMap)
	if err != nil {
		return nil, err
	}

	retMap := make(map[string]*model.PurchaseOrder)
	for _, id := range ids {
		val, ok := itemMap[c.GetPurchaseOrderCacheKey(id)]
		if ok {
			retMap[id] = val
		}
	}

	return retMap, nil
}

// Del delete cache
func (c *purchaseOrderCache) Del(ctx context.Context, id string) error {
	cacheKey := c.GetPurchaseOrderCacheKey(id)
	err := c.cache.Del(ctx, cacheKey)
	if err != nil {
		return err
	}
	return nil
}

// SetPlaceholder set placeholder value to cache
func (c *purchaseOrderCache) SetPlaceholder(ctx context.Context, id string) error {
	cacheKey := c.GetPurchaseOrderCacheKey(id)
	return c.cache.SetCacheWithNotFound(ctx, cacheKey)
}

// IsPlaceholderErr check if cache is placeholder error
func (c *purchaseOrderCache) IsPlaceholderErr(err error) bool {
	return errors.Is(err, cache.ErrPlaceholder)
}
