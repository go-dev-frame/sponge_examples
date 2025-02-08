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
	customerCachePrefixKey = "customer:"
	// CustomerExpireTime expire time
	CustomerExpireTime = 5 * time.Minute
)

var _ CustomerCache = (*customerCache)(nil)

// CustomerCache cache interface
type CustomerCache interface {
	Set(ctx context.Context, id uint64, data *model.Customer, duration time.Duration) error
	Get(ctx context.Context, id uint64) (*model.Customer, error)
	MultiGet(ctx context.Context, ids []uint64) (map[uint64]*model.Customer, error)
	MultiSet(ctx context.Context, data []*model.Customer, duration time.Duration) error
	Del(ctx context.Context, id uint64) error
	SetPlaceholder(ctx context.Context, id uint64) error
	IsPlaceholderErr(err error) bool
}

// customerCache define a cache struct
type customerCache struct {
	cache cache.Cache
}

// NewCustomerCache new a cache
func NewCustomerCache(cacheType *database.CacheType) CustomerCache {
	jsonEncoding := encoding.JSONEncoding{}
	cachePrefix := ""

	cType := strings.ToLower(cacheType.CType)
	switch cType {
	case "redis":
		c := cache.NewRedisCache(cacheType.Rdb, cachePrefix, jsonEncoding, func() interface{} {
			return &model.Customer{}
		})
		return &customerCache{cache: c}
	case "memory":
		c := cache.NewMemoryCache(cachePrefix, jsonEncoding, func() interface{} {
			return &model.Customer{}
		})
		return &customerCache{cache: c}
	}

	return nil // no cache
}

// GetCustomerCacheKey cache key
func (c *customerCache) GetCustomerCacheKey(id uint64) string {
	return customerCachePrefixKey + utils.Uint64ToStr(id)
}

// Set write to cache
func (c *customerCache) Set(ctx context.Context, id uint64, data *model.Customer, duration time.Duration) error {
	if data == nil || id == 0 {
		return nil
	}
	cacheKey := c.GetCustomerCacheKey(id)
	err := c.cache.Set(ctx, cacheKey, data, duration)
	if err != nil {
		return err
	}
	return nil
}

// Get cache value
func (c *customerCache) Get(ctx context.Context, id uint64) (*model.Customer, error) {
	var data *model.Customer
	cacheKey := c.GetCustomerCacheKey(id)
	err := c.cache.Get(ctx, cacheKey, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// MultiSet multiple set cache
func (c *customerCache) MultiSet(ctx context.Context, data []*model.Customer, duration time.Duration) error {
	valMap := make(map[string]interface{})
	for _, v := range data {
		cacheKey := c.GetCustomerCacheKey(v.ID)
		valMap[cacheKey] = v
	}

	err := c.cache.MultiSet(ctx, valMap, duration)
	if err != nil {
		return err
	}

	return nil
}

// MultiGet multiple get cache, return key in map is id value
func (c *customerCache) MultiGet(ctx context.Context, ids []uint64) (map[uint64]*model.Customer, error) {
	var keys []string
	for _, v := range ids {
		cacheKey := c.GetCustomerCacheKey(v)
		keys = append(keys, cacheKey)
	}

	itemMap := make(map[string]*model.Customer)
	err := c.cache.MultiGet(ctx, keys, itemMap)
	if err != nil {
		return nil, err
	}

	retMap := make(map[uint64]*model.Customer)
	for _, id := range ids {
		val, ok := itemMap[c.GetCustomerCacheKey(id)]
		if ok {
			retMap[id] = val
		}
	}

	return retMap, nil
}

// Del delete cache
func (c *customerCache) Del(ctx context.Context, id uint64) error {
	cacheKey := c.GetCustomerCacheKey(id)
	err := c.cache.Del(ctx, cacheKey)
	if err != nil {
		return err
	}
	return nil
}

// SetPlaceholder set placeholder value to cache
func (c *customerCache) SetPlaceholder(ctx context.Context, id uint64) error {
	cacheKey := c.GetCustomerCacheKey(id)
	return c.cache.SetCacheWithNotFound(ctx, cacheKey)
}

// IsPlaceholderErr check if cache is placeholder error
func (c *customerCache) IsPlaceholderErr(err error) bool {
	return errors.Is(err, cache.ErrPlaceholder)
}
