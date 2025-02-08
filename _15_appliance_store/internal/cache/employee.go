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
	employeeCachePrefixKey = "employee:"
	// EmployeeExpireTime expire time
	EmployeeExpireTime = 5 * time.Minute
)

var _ EmployeeCache = (*employeeCache)(nil)

// EmployeeCache cache interface
type EmployeeCache interface {
	Set(ctx context.Context, id uint64, data *model.Employee, duration time.Duration) error
	Get(ctx context.Context, id uint64) (*model.Employee, error)
	MultiGet(ctx context.Context, ids []uint64) (map[uint64]*model.Employee, error)
	MultiSet(ctx context.Context, data []*model.Employee, duration time.Duration) error
	Del(ctx context.Context, id uint64) error
	SetPlaceholder(ctx context.Context, id uint64) error
	IsPlaceholderErr(err error) bool
}

// employeeCache define a cache struct
type employeeCache struct {
	cache cache.Cache
}

// NewEmployeeCache new a cache
func NewEmployeeCache(cacheType *database.CacheType) EmployeeCache {
	jsonEncoding := encoding.JSONEncoding{}
	cachePrefix := ""

	cType := strings.ToLower(cacheType.CType)
	switch cType {
	case "redis":
		c := cache.NewRedisCache(cacheType.Rdb, cachePrefix, jsonEncoding, func() interface{} {
			return &model.Employee{}
		})
		return &employeeCache{cache: c}
	case "memory":
		c := cache.NewMemoryCache(cachePrefix, jsonEncoding, func() interface{} {
			return &model.Employee{}
		})
		return &employeeCache{cache: c}
	}

	return nil // no cache
}

// GetEmployeeCacheKey cache key
func (c *employeeCache) GetEmployeeCacheKey(id uint64) string {
	return employeeCachePrefixKey + utils.Uint64ToStr(id)
}

// Set write to cache
func (c *employeeCache) Set(ctx context.Context, id uint64, data *model.Employee, duration time.Duration) error {
	if data == nil || id == 0 {
		return nil
	}
	cacheKey := c.GetEmployeeCacheKey(id)
	err := c.cache.Set(ctx, cacheKey, data, duration)
	if err != nil {
		return err
	}
	return nil
}

// Get cache value
func (c *employeeCache) Get(ctx context.Context, id uint64) (*model.Employee, error) {
	var data *model.Employee
	cacheKey := c.GetEmployeeCacheKey(id)
	err := c.cache.Get(ctx, cacheKey, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// MultiSet multiple set cache
func (c *employeeCache) MultiSet(ctx context.Context, data []*model.Employee, duration time.Duration) error {
	valMap := make(map[string]interface{})
	for _, v := range data {
		cacheKey := c.GetEmployeeCacheKey(v.ID)
		valMap[cacheKey] = v
	}

	err := c.cache.MultiSet(ctx, valMap, duration)
	if err != nil {
		return err
	}

	return nil
}

// MultiGet multiple get cache, return key in map is id value
func (c *employeeCache) MultiGet(ctx context.Context, ids []uint64) (map[uint64]*model.Employee, error) {
	var keys []string
	for _, v := range ids {
		cacheKey := c.GetEmployeeCacheKey(v)
		keys = append(keys, cacheKey)
	}

	itemMap := make(map[string]*model.Employee)
	err := c.cache.MultiGet(ctx, keys, itemMap)
	if err != nil {
		return nil, err
	}

	retMap := make(map[uint64]*model.Employee)
	for _, id := range ids {
		val, ok := itemMap[c.GetEmployeeCacheKey(id)]
		if ok {
			retMap[id] = val
		}
	}

	return retMap, nil
}

// Del delete cache
func (c *employeeCache) Del(ctx context.Context, id uint64) error {
	cacheKey := c.GetEmployeeCacheKey(id)
	err := c.cache.Del(ctx, cacheKey)
	if err != nil {
		return err
	}
	return nil
}

// SetPlaceholder set placeholder value to cache
func (c *employeeCache) SetPlaceholder(ctx context.Context, id uint64) error {
	cacheKey := c.GetEmployeeCacheKey(id)
	return c.cache.SetCacheWithNotFound(ctx, cacheKey)
}

// IsPlaceholderErr check if cache is placeholder error
func (c *employeeCache) IsPlaceholderErr(err error) bool {
	return errors.Is(err, cache.ErrPlaceholder)
}
