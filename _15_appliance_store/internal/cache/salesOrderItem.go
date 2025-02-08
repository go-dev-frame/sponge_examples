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
	salesOrderItemCachePrefixKey = "salesOrderItem:"
	// SalesOrderItemExpireTime expire time
	SalesOrderItemExpireTime = 5 * time.Minute
)

var _ SalesOrderItemCache = (*salesOrderItemCache)(nil)

// SalesOrderItemCache cache interface
type SalesOrderItemCache interface {
	Set(ctx context.Context, id uint64, data *model.SalesOrderItem, duration time.Duration) error
	Get(ctx context.Context, id uint64) (*model.SalesOrderItem, error)
	MultiGet(ctx context.Context, ids []uint64) (map[uint64]*model.SalesOrderItem, error)
	MultiSet(ctx context.Context, data []*model.SalesOrderItem, duration time.Duration) error
	Del(ctx context.Context, id uint64) error
	SetPlaceholder(ctx context.Context, id uint64) error
	IsPlaceholderErr(err error) bool
}

// salesOrderItemCache define a cache struct
type salesOrderItemCache struct {
	cache cache.Cache
}

// NewSalesOrderItemCache new a cache
func NewSalesOrderItemCache(cacheType *database.CacheType) SalesOrderItemCache {
	jsonEncoding := encoding.JSONEncoding{}
	cachePrefix := ""

	cType := strings.ToLower(cacheType.CType)
	switch cType {
	case "redis":
		c := cache.NewRedisCache(cacheType.Rdb, cachePrefix, jsonEncoding, func() interface{} {
			return &model.SalesOrderItem{}
		})
		return &salesOrderItemCache{cache: c}
	case "memory":
		c := cache.NewMemoryCache(cachePrefix, jsonEncoding, func() interface{} {
			return &model.SalesOrderItem{}
		})
		return &salesOrderItemCache{cache: c}
	}

	return nil // no cache
}

// GetSalesOrderItemCacheKey cache key
func (c *salesOrderItemCache) GetSalesOrderItemCacheKey(id uint64) string {
	return salesOrderItemCachePrefixKey + utils.Uint64ToStr(id)
}

// Set write to cache
func (c *salesOrderItemCache) Set(ctx context.Context, id uint64, data *model.SalesOrderItem, duration time.Duration) error {
	if data == nil || id == 0 {
		return nil
	}
	cacheKey := c.GetSalesOrderItemCacheKey(id)
	err := c.cache.Set(ctx, cacheKey, data, duration)
	if err != nil {
		return err
	}
	return nil
}

// Get cache value
func (c *salesOrderItemCache) Get(ctx context.Context, id uint64) (*model.SalesOrderItem, error) {
	var data *model.SalesOrderItem
	cacheKey := c.GetSalesOrderItemCacheKey(id)
	err := c.cache.Get(ctx, cacheKey, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// MultiSet multiple set cache
func (c *salesOrderItemCache) MultiSet(ctx context.Context, data []*model.SalesOrderItem, duration time.Duration) error {
	valMap := make(map[string]interface{})
	for _, v := range data {
		cacheKey := c.GetSalesOrderItemCacheKey(v.ID)
		valMap[cacheKey] = v
	}

	err := c.cache.MultiSet(ctx, valMap, duration)
	if err != nil {
		return err
	}

	return nil
}

// MultiGet multiple get cache, return key in map is id value
func (c *salesOrderItemCache) MultiGet(ctx context.Context, ids []uint64) (map[uint64]*model.SalesOrderItem, error) {
	var keys []string
	for _, v := range ids {
		cacheKey := c.GetSalesOrderItemCacheKey(v)
		keys = append(keys, cacheKey)
	}

	itemMap := make(map[string]*model.SalesOrderItem)
	err := c.cache.MultiGet(ctx, keys, itemMap)
	if err != nil {
		return nil, err
	}

	retMap := make(map[uint64]*model.SalesOrderItem)
	for _, id := range ids {
		val, ok := itemMap[c.GetSalesOrderItemCacheKey(id)]
		if ok {
			retMap[id] = val
		}
	}

	return retMap, nil
}

// Del delete cache
func (c *salesOrderItemCache) Del(ctx context.Context, id uint64) error {
	cacheKey := c.GetSalesOrderItemCacheKey(id)
	err := c.cache.Del(ctx, cacheKey)
	if err != nil {
		return err
	}
	return nil
}

// SetPlaceholder set placeholder value to cache
func (c *salesOrderItemCache) SetPlaceholder(ctx context.Context, id uint64) error {
	cacheKey := c.GetSalesOrderItemCacheKey(id)
	return c.cache.SetCacheWithNotFound(ctx, cacheKey)
}

// IsPlaceholderErr check if cache is placeholder error
func (c *salesOrderItemCache) IsPlaceholderErr(err error) bool {
	return errors.Is(err, cache.ErrPlaceholder)
}
