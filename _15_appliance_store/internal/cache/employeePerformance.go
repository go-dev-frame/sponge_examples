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
	employeePerformanceCachePrefixKey = "employeePerformance:"
	// EmployeePerformanceExpireTime expire time
	EmployeePerformanceExpireTime = 5 * time.Minute
)

var _ EmployeePerformanceCache = (*employeePerformanceCache)(nil)

// EmployeePerformanceCache cache interface
type EmployeePerformanceCache interface {
	Set(ctx context.Context, id uint64, data *model.EmployeePerformance, duration time.Duration) error
	Get(ctx context.Context, id uint64) (*model.EmployeePerformance, error)
	MultiGet(ctx context.Context, ids []uint64) (map[uint64]*model.EmployeePerformance, error)
	MultiSet(ctx context.Context, data []*model.EmployeePerformance, duration time.Duration) error
	Del(ctx context.Context, id uint64) error
	SetPlaceholder(ctx context.Context, id uint64) error
	IsPlaceholderErr(err error) bool
}

// employeePerformanceCache define a cache struct
type employeePerformanceCache struct {
	cache cache.Cache
}

// NewEmployeePerformanceCache new a cache
func NewEmployeePerformanceCache(cacheType *database.CacheType) EmployeePerformanceCache {
	jsonEncoding := encoding.JSONEncoding{}
	cachePrefix := ""

	cType := strings.ToLower(cacheType.CType)
	switch cType {
	case "redis":
		c := cache.NewRedisCache(cacheType.Rdb, cachePrefix, jsonEncoding, func() interface{} {
			return &model.EmployeePerformance{}
		})
		return &employeePerformanceCache{cache: c}
	case "memory":
		c := cache.NewMemoryCache(cachePrefix, jsonEncoding, func() interface{} {
			return &model.EmployeePerformance{}
		})
		return &employeePerformanceCache{cache: c}
	}

	return nil // no cache
}

// GetEmployeePerformanceCacheKey cache key
func (c *employeePerformanceCache) GetEmployeePerformanceCacheKey(id uint64) string {
	return employeePerformanceCachePrefixKey + utils.Uint64ToStr(id)
}

// Set write to cache
func (c *employeePerformanceCache) Set(ctx context.Context, id uint64, data *model.EmployeePerformance, duration time.Duration) error {
	if data == nil || id == 0 {
		return nil
	}
	cacheKey := c.GetEmployeePerformanceCacheKey(id)
	err := c.cache.Set(ctx, cacheKey, data, duration)
	if err != nil {
		return err
	}
	return nil
}

// Get cache value
func (c *employeePerformanceCache) Get(ctx context.Context, id uint64) (*model.EmployeePerformance, error) {
	var data *model.EmployeePerformance
	cacheKey := c.GetEmployeePerformanceCacheKey(id)
	err := c.cache.Get(ctx, cacheKey, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// MultiSet multiple set cache
func (c *employeePerformanceCache) MultiSet(ctx context.Context, data []*model.EmployeePerformance, duration time.Duration) error {
	valMap := make(map[string]interface{})
	for _, v := range data {
		cacheKey := c.GetEmployeePerformanceCacheKey(v.ID)
		valMap[cacheKey] = v
	}

	err := c.cache.MultiSet(ctx, valMap, duration)
	if err != nil {
		return err
	}

	return nil
}

// MultiGet multiple get cache, return key in map is id value
func (c *employeePerformanceCache) MultiGet(ctx context.Context, ids []uint64) (map[uint64]*model.EmployeePerformance, error) {
	var keys []string
	for _, v := range ids {
		cacheKey := c.GetEmployeePerformanceCacheKey(v)
		keys = append(keys, cacheKey)
	}

	itemMap := make(map[string]*model.EmployeePerformance)
	err := c.cache.MultiGet(ctx, keys, itemMap)
	if err != nil {
		return nil, err
	}

	retMap := make(map[uint64]*model.EmployeePerformance)
	for _, id := range ids {
		val, ok := itemMap[c.GetEmployeePerformanceCacheKey(id)]
		if ok {
			retMap[id] = val
		}
	}

	return retMap, nil
}

// Del delete cache
func (c *employeePerformanceCache) Del(ctx context.Context, id uint64) error {
	cacheKey := c.GetEmployeePerformanceCacheKey(id)
	err := c.cache.Del(ctx, cacheKey)
	if err != nil {
		return err
	}
	return nil
}

// SetPlaceholder set placeholder value to cache
func (c *employeePerformanceCache) SetPlaceholder(ctx context.Context, id uint64) error {
	cacheKey := c.GetEmployeePerformanceCacheKey(id)
	return c.cache.SetCacheWithNotFound(ctx, cacheKey)
}

// IsPlaceholderErr check if cache is placeholder error
func (c *employeePerformanceCache) IsPlaceholderErr(err error) bool {
	return errors.Is(err, cache.ErrPlaceholder)
}
