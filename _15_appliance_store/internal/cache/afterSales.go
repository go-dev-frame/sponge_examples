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
	afterSalesCachePrefixKey = "afterSales:"
	// AfterSalesExpireTime expire time
	AfterSalesExpireTime = 5 * time.Minute
)

var _ AfterSalesCache = (*afterSalesCache)(nil)

// AfterSalesCache cache interface
type AfterSalesCache interface {
	Set(ctx context.Context, id uint64, data *model.AfterSales, duration time.Duration) error
	Get(ctx context.Context, id uint64) (*model.AfterSales, error)
	MultiGet(ctx context.Context, ids []uint64) (map[uint64]*model.AfterSales, error)
	MultiSet(ctx context.Context, data []*model.AfterSales, duration time.Duration) error
	Del(ctx context.Context, id uint64) error
	SetPlaceholder(ctx context.Context, id uint64) error
	IsPlaceholderErr(err error) bool
}

// afterSalesCache define a cache struct
type afterSalesCache struct {
	cache cache.Cache
}

// NewAfterSalesCache new a cache
func NewAfterSalesCache(cacheType *database.CacheType) AfterSalesCache {
	jsonEncoding := encoding.JSONEncoding{}
	cachePrefix := ""

	cType := strings.ToLower(cacheType.CType)
	switch cType {
	case "redis":
		c := cache.NewRedisCache(cacheType.Rdb, cachePrefix, jsonEncoding, func() interface{} {
			return &model.AfterSales{}
		})
		return &afterSalesCache{cache: c}
	case "memory":
		c := cache.NewMemoryCache(cachePrefix, jsonEncoding, func() interface{} {
			return &model.AfterSales{}
		})
		return &afterSalesCache{cache: c}
	}

	return nil // no cache
}

// GetAfterSalesCacheKey cache key
func (c *afterSalesCache) GetAfterSalesCacheKey(id uint64) string {
	return afterSalesCachePrefixKey + utils.Uint64ToStr(id)
}

// Set write to cache
func (c *afterSalesCache) Set(ctx context.Context, id uint64, data *model.AfterSales, duration time.Duration) error {
	if data == nil || id == 0 {
		return nil
	}
	cacheKey := c.GetAfterSalesCacheKey(id)
	err := c.cache.Set(ctx, cacheKey, data, duration)
	if err != nil {
		return err
	}
	return nil
}

// Get cache value
func (c *afterSalesCache) Get(ctx context.Context, id uint64) (*model.AfterSales, error) {
	var data *model.AfterSales
	cacheKey := c.GetAfterSalesCacheKey(id)
	err := c.cache.Get(ctx, cacheKey, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// MultiSet multiple set cache
func (c *afterSalesCache) MultiSet(ctx context.Context, data []*model.AfterSales, duration time.Duration) error {
	valMap := make(map[string]interface{})
	for _, v := range data {
		cacheKey := c.GetAfterSalesCacheKey(v.ID)
		valMap[cacheKey] = v
	}

	err := c.cache.MultiSet(ctx, valMap, duration)
	if err != nil {
		return err
	}

	return nil
}

// MultiGet multiple get cache, return key in map is id value
func (c *afterSalesCache) MultiGet(ctx context.Context, ids []uint64) (map[uint64]*model.AfterSales, error) {
	var keys []string
	for _, v := range ids {
		cacheKey := c.GetAfterSalesCacheKey(v)
		keys = append(keys, cacheKey)
	}

	itemMap := make(map[string]*model.AfterSales)
	err := c.cache.MultiGet(ctx, keys, itemMap)
	if err != nil {
		return nil, err
	}

	retMap := make(map[uint64]*model.AfterSales)
	for _, id := range ids {
		val, ok := itemMap[c.GetAfterSalesCacheKey(id)]
		if ok {
			retMap[id] = val
		}
	}

	return retMap, nil
}

// Del delete cache
func (c *afterSalesCache) Del(ctx context.Context, id uint64) error {
	cacheKey := c.GetAfterSalesCacheKey(id)
	err := c.cache.Del(ctx, cacheKey)
	if err != nil {
		return err
	}
	return nil
}

// SetPlaceholder set placeholder value to cache
func (c *afterSalesCache) SetPlaceholder(ctx context.Context, id uint64) error {
	cacheKey := c.GetAfterSalesCacheKey(id)
	return c.cache.SetCacheWithNotFound(ctx, cacheKey)
}

// IsPlaceholderErr check if cache is placeholder error
func (c *afterSalesCache) IsPlaceholderErr(err error) bool {
	return errors.Is(err, cache.ErrPlaceholder)
}
