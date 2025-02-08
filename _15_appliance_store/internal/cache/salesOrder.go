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
	salesOrderCachePrefixKey = "salesOrder:"
	// SalesOrderExpireTime expire time
	SalesOrderExpireTime = 5 * time.Minute
)

var _ SalesOrderCache = (*salesOrderCache)(nil)

// SalesOrderCache cache interface
type SalesOrderCache interface {
	Set(ctx context.Context, id string, data *model.SalesOrder, duration time.Duration) error
	Get(ctx context.Context, id string) (*model.SalesOrder, error)
	MultiGet(ctx context.Context, ids []string) (map[string]*model.SalesOrder, error)
	MultiSet(ctx context.Context, data []*model.SalesOrder, duration time.Duration) error
	Del(ctx context.Context, id string) error
	SetPlaceholder(ctx context.Context, id string) error
	IsPlaceholderErr(err error) bool
}

// salesOrderCache define a cache struct
type salesOrderCache struct {
	cache cache.Cache
}

// NewSalesOrderCache new a cache
func NewSalesOrderCache(cacheType *database.CacheType) SalesOrderCache {
	jsonEncoding := encoding.JSONEncoding{}
	cachePrefix := ""

	cType := strings.ToLower(cacheType.CType)
	switch cType {
	case "redis":
		c := cache.NewRedisCache(cacheType.Rdb, cachePrefix, jsonEncoding, func() interface{} {
			return &model.SalesOrder{}
		})
		return &salesOrderCache{cache: c}
	case "memory":
		c := cache.NewMemoryCache(cachePrefix, jsonEncoding, func() interface{} {
			return &model.SalesOrder{}
		})
		return &salesOrderCache{cache: c}
	}

	return nil // no cache
}

// GetSalesOrderCacheKey cache key
func (c *salesOrderCache) GetSalesOrderCacheKey(id string) string {
	return salesOrderCachePrefixKey + id
}

// Set write to cache
func (c *salesOrderCache) Set(ctx context.Context, id string, data *model.SalesOrder, duration time.Duration) error {
	if data == nil {
		return nil
	}
	cacheKey := c.GetSalesOrderCacheKey(id)
	err := c.cache.Set(ctx, cacheKey, data, duration)
	if err != nil {
		return err
	}
	return nil
}

// Get cache value
func (c *salesOrderCache) Get(ctx context.Context, id string) (*model.SalesOrder, error) {
	var data *model.SalesOrder
	cacheKey := c.GetSalesOrderCacheKey(id)
	err := c.cache.Get(ctx, cacheKey, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// MultiSet multiple set cache
func (c *salesOrderCache) MultiSet(ctx context.Context, data []*model.SalesOrder, duration time.Duration) error {
	valMap := make(map[string]interface{})
	for _, v := range data {
		cacheKey := c.GetSalesOrderCacheKey(v.ID)
		valMap[cacheKey] = v
	}

	err := c.cache.MultiSet(ctx, valMap, duration)
	if err != nil {
		return err
	}

	return nil
}

// MultiGet multiple get cache, return key in map is id value
func (c *salesOrderCache) MultiGet(ctx context.Context, ids []string) (map[string]*model.SalesOrder, error) {
	var keys []string
	for _, v := range ids {
		cacheKey := c.GetSalesOrderCacheKey(v)
		keys = append(keys, cacheKey)
	}

	itemMap := make(map[string]*model.SalesOrder)
	err := c.cache.MultiGet(ctx, keys, itemMap)
	if err != nil {
		return nil, err
	}

	retMap := make(map[string]*model.SalesOrder)
	for _, id := range ids {
		val, ok := itemMap[c.GetSalesOrderCacheKey(id)]
		if ok {
			retMap[id] = val
		}
	}

	return retMap, nil
}

// Del delete cache
func (c *salesOrderCache) Del(ctx context.Context, id string) error {
	cacheKey := c.GetSalesOrderCacheKey(id)
	err := c.cache.Del(ctx, cacheKey)
	if err != nil {
		return err
	}
	return nil
}

// SetPlaceholder set placeholder value to cache
func (c *salesOrderCache) SetPlaceholder(ctx context.Context, id string) error {
	cacheKey := c.GetSalesOrderCacheKey(id)
	return c.cache.SetCacheWithNotFound(ctx, cacheKey)
}

// IsPlaceholderErr check if cache is placeholder error
func (c *salesOrderCache) IsPlaceholderErr(err error) bool {
	return errors.Is(err, cache.ErrPlaceholder)
}
