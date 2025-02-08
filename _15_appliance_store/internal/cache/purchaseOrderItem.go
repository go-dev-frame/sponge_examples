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
	purchaseOrderItemCachePrefixKey = "purchaseOrderItem:"
	// PurchaseOrderItemExpireTime expire time
	PurchaseOrderItemExpireTime = 5 * time.Minute
)

var _ PurchaseOrderItemCache = (*purchaseOrderItemCache)(nil)

// PurchaseOrderItemCache cache interface
type PurchaseOrderItemCache interface {
	Set(ctx context.Context, id uint64, data *model.PurchaseOrderItem, duration time.Duration) error
	Get(ctx context.Context, id uint64) (*model.PurchaseOrderItem, error)
	MultiGet(ctx context.Context, ids []uint64) (map[uint64]*model.PurchaseOrderItem, error)
	MultiSet(ctx context.Context, data []*model.PurchaseOrderItem, duration time.Duration) error
	Del(ctx context.Context, id uint64) error
	SetPlaceholder(ctx context.Context, id uint64) error
	IsPlaceholderErr(err error) bool
}

// purchaseOrderItemCache define a cache struct
type purchaseOrderItemCache struct {
	cache cache.Cache
}

// NewPurchaseOrderItemCache new a cache
func NewPurchaseOrderItemCache(cacheType *database.CacheType) PurchaseOrderItemCache {
	jsonEncoding := encoding.JSONEncoding{}
	cachePrefix := ""

	cType := strings.ToLower(cacheType.CType)
	switch cType {
	case "redis":
		c := cache.NewRedisCache(cacheType.Rdb, cachePrefix, jsonEncoding, func() interface{} {
			return &model.PurchaseOrderItem{}
		})
		return &purchaseOrderItemCache{cache: c}
	case "memory":
		c := cache.NewMemoryCache(cachePrefix, jsonEncoding, func() interface{} {
			return &model.PurchaseOrderItem{}
		})
		return &purchaseOrderItemCache{cache: c}
	}

	return nil // no cache
}

// GetPurchaseOrderItemCacheKey cache key
func (c *purchaseOrderItemCache) GetPurchaseOrderItemCacheKey(id uint64) string {
	return purchaseOrderItemCachePrefixKey + utils.Uint64ToStr(id)
}

// Set write to cache
func (c *purchaseOrderItemCache) Set(ctx context.Context, id uint64, data *model.PurchaseOrderItem, duration time.Duration) error {
	if data == nil || id == 0 {
		return nil
	}
	cacheKey := c.GetPurchaseOrderItemCacheKey(id)
	err := c.cache.Set(ctx, cacheKey, data, duration)
	if err != nil {
		return err
	}
	return nil
}

// Get cache value
func (c *purchaseOrderItemCache) Get(ctx context.Context, id uint64) (*model.PurchaseOrderItem, error) {
	var data *model.PurchaseOrderItem
	cacheKey := c.GetPurchaseOrderItemCacheKey(id)
	err := c.cache.Get(ctx, cacheKey, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// MultiSet multiple set cache
func (c *purchaseOrderItemCache) MultiSet(ctx context.Context, data []*model.PurchaseOrderItem, duration time.Duration) error {
	valMap := make(map[string]interface{})
	for _, v := range data {
		cacheKey := c.GetPurchaseOrderItemCacheKey(v.ID)
		valMap[cacheKey] = v
	}

	err := c.cache.MultiSet(ctx, valMap, duration)
	if err != nil {
		return err
	}

	return nil
}

// MultiGet multiple get cache, return key in map is id value
func (c *purchaseOrderItemCache) MultiGet(ctx context.Context, ids []uint64) (map[uint64]*model.PurchaseOrderItem, error) {
	var keys []string
	for _, v := range ids {
		cacheKey := c.GetPurchaseOrderItemCacheKey(v)
		keys = append(keys, cacheKey)
	}

	itemMap := make(map[string]*model.PurchaseOrderItem)
	err := c.cache.MultiGet(ctx, keys, itemMap)
	if err != nil {
		return nil, err
	}

	retMap := make(map[uint64]*model.PurchaseOrderItem)
	for _, id := range ids {
		val, ok := itemMap[c.GetPurchaseOrderItemCacheKey(id)]
		if ok {
			retMap[id] = val
		}
	}

	return retMap, nil
}

// Del delete cache
func (c *purchaseOrderItemCache) Del(ctx context.Context, id uint64) error {
	cacheKey := c.GetPurchaseOrderItemCacheKey(id)
	err := c.cache.Del(ctx, cacheKey)
	if err != nil {
		return err
	}
	return nil
}

// SetPlaceholder set placeholder value to cache
func (c *purchaseOrderItemCache) SetPlaceholder(ctx context.Context, id uint64) error {
	cacheKey := c.GetPurchaseOrderItemCacheKey(id)
	return c.cache.SetCacheWithNotFound(ctx, cacheKey)
}

// IsPlaceholderErr check if cache is placeholder error
func (c *purchaseOrderItemCache) IsPlaceholderErr(err error) bool {
	return errors.Is(err, cache.ErrPlaceholder)
}
