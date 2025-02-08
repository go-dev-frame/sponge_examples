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
	serviceOrderCachePrefixKey = "serviceOrder:"
	// ServiceOrderExpireTime expire time
	ServiceOrderExpireTime = 5 * time.Minute
)

var _ ServiceOrderCache = (*serviceOrderCache)(nil)

// ServiceOrderCache cache interface
type ServiceOrderCache interface {
	Set(ctx context.Context, id string, data *model.ServiceOrder, duration time.Duration) error
	Get(ctx context.Context, id string) (*model.ServiceOrder, error)
	MultiGet(ctx context.Context, ids []string) (map[string]*model.ServiceOrder, error)
	MultiSet(ctx context.Context, data []*model.ServiceOrder, duration time.Duration) error
	Del(ctx context.Context, id string) error
	SetPlaceholder(ctx context.Context, id string) error
	IsPlaceholderErr(err error) bool
}

// serviceOrderCache define a cache struct
type serviceOrderCache struct {
	cache cache.Cache
}

// NewServiceOrderCache new a cache
func NewServiceOrderCache(cacheType *database.CacheType) ServiceOrderCache {
	jsonEncoding := encoding.JSONEncoding{}
	cachePrefix := ""

	cType := strings.ToLower(cacheType.CType)
	switch cType {
	case "redis":
		c := cache.NewRedisCache(cacheType.Rdb, cachePrefix, jsonEncoding, func() interface{} {
			return &model.ServiceOrder{}
		})
		return &serviceOrderCache{cache: c}
	case "memory":
		c := cache.NewMemoryCache(cachePrefix, jsonEncoding, func() interface{} {
			return &model.ServiceOrder{}
		})
		return &serviceOrderCache{cache: c}
	}

	return nil // no cache
}

// GetServiceOrderCacheKey cache key
func (c *serviceOrderCache) GetServiceOrderCacheKey(id string) string {
	return serviceOrderCachePrefixKey + id
}

// Set write to cache
func (c *serviceOrderCache) Set(ctx context.Context, id string, data *model.ServiceOrder, duration time.Duration) error {
	if data == nil {
		return nil
	}
	cacheKey := c.GetServiceOrderCacheKey(id)
	err := c.cache.Set(ctx, cacheKey, data, duration)
	if err != nil {
		return err
	}
	return nil
}

// Get cache value
func (c *serviceOrderCache) Get(ctx context.Context, id string) (*model.ServiceOrder, error) {
	var data *model.ServiceOrder
	cacheKey := c.GetServiceOrderCacheKey(id)
	err := c.cache.Get(ctx, cacheKey, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// MultiSet multiple set cache
func (c *serviceOrderCache) MultiSet(ctx context.Context, data []*model.ServiceOrder, duration time.Duration) error {
	valMap := make(map[string]interface{})
	for _, v := range data {
		cacheKey := c.GetServiceOrderCacheKey(v.ID)
		valMap[cacheKey] = v
	}

	err := c.cache.MultiSet(ctx, valMap, duration)
	if err != nil {
		return err
	}

	return nil
}

// MultiGet multiple get cache, return key in map is id value
func (c *serviceOrderCache) MultiGet(ctx context.Context, ids []string) (map[string]*model.ServiceOrder, error) {
	var keys []string
	for _, v := range ids {
		cacheKey := c.GetServiceOrderCacheKey(v)
		keys = append(keys, cacheKey)
	}

	itemMap := make(map[string]*model.ServiceOrder)
	err := c.cache.MultiGet(ctx, keys, itemMap)
	if err != nil {
		return nil, err
	}

	retMap := make(map[string]*model.ServiceOrder)
	for _, id := range ids {
		val, ok := itemMap[c.GetServiceOrderCacheKey(id)]
		if ok {
			retMap[id] = val
		}
	}

	return retMap, nil
}

// Del delete cache
func (c *serviceOrderCache) Del(ctx context.Context, id string) error {
	cacheKey := c.GetServiceOrderCacheKey(id)
	err := c.cache.Del(ctx, cacheKey)
	if err != nil {
		return err
	}
	return nil
}

// SetPlaceholder set placeholder value to cache
func (c *serviceOrderCache) SetPlaceholder(ctx context.Context, id string) error {
	cacheKey := c.GetServiceOrderCacheKey(id)
	return c.cache.SetCacheWithNotFound(ctx, cacheKey)
}

// IsPlaceholderErr check if cache is placeholder error
func (c *serviceOrderCache) IsPlaceholderErr(err error) bool {
	return errors.Is(err, cache.ErrPlaceholder)
}
