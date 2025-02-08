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
	supplierCachePrefixKey = "supplier:"
	// SupplierExpireTime expire time
	SupplierExpireTime = 5 * time.Minute
)

var _ SupplierCache = (*supplierCache)(nil)

// SupplierCache cache interface
type SupplierCache interface {
	Set(ctx context.Context, id uint64, data *model.Supplier, duration time.Duration) error
	Get(ctx context.Context, id uint64) (*model.Supplier, error)
	MultiGet(ctx context.Context, ids []uint64) (map[uint64]*model.Supplier, error)
	MultiSet(ctx context.Context, data []*model.Supplier, duration time.Duration) error
	Del(ctx context.Context, id uint64) error
	SetPlaceholder(ctx context.Context, id uint64) error
	IsPlaceholderErr(err error) bool
}

// supplierCache define a cache struct
type supplierCache struct {
	cache cache.Cache
}

// NewSupplierCache new a cache
func NewSupplierCache(cacheType *database.CacheType) SupplierCache {
	jsonEncoding := encoding.JSONEncoding{}
	cachePrefix := ""

	cType := strings.ToLower(cacheType.CType)
	switch cType {
	case "redis":
		c := cache.NewRedisCache(cacheType.Rdb, cachePrefix, jsonEncoding, func() interface{} {
			return &model.Supplier{}
		})
		return &supplierCache{cache: c}
	case "memory":
		c := cache.NewMemoryCache(cachePrefix, jsonEncoding, func() interface{} {
			return &model.Supplier{}
		})
		return &supplierCache{cache: c}
	}

	return nil // no cache
}

// GetSupplierCacheKey cache key
func (c *supplierCache) GetSupplierCacheKey(id uint64) string {
	return supplierCachePrefixKey + utils.Uint64ToStr(id)
}

// Set write to cache
func (c *supplierCache) Set(ctx context.Context, id uint64, data *model.Supplier, duration time.Duration) error {
	if data == nil || id == 0 {
		return nil
	}
	cacheKey := c.GetSupplierCacheKey(id)
	err := c.cache.Set(ctx, cacheKey, data, duration)
	if err != nil {
		return err
	}
	return nil
}

// Get cache value
func (c *supplierCache) Get(ctx context.Context, id uint64) (*model.Supplier, error) {
	var data *model.Supplier
	cacheKey := c.GetSupplierCacheKey(id)
	err := c.cache.Get(ctx, cacheKey, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// MultiSet multiple set cache
func (c *supplierCache) MultiSet(ctx context.Context, data []*model.Supplier, duration time.Duration) error {
	valMap := make(map[string]interface{})
	for _, v := range data {
		cacheKey := c.GetSupplierCacheKey(v.ID)
		valMap[cacheKey] = v
	}

	err := c.cache.MultiSet(ctx, valMap, duration)
	if err != nil {
		return err
	}

	return nil
}

// MultiGet multiple get cache, return key in map is id value
func (c *supplierCache) MultiGet(ctx context.Context, ids []uint64) (map[uint64]*model.Supplier, error) {
	var keys []string
	for _, v := range ids {
		cacheKey := c.GetSupplierCacheKey(v)
		keys = append(keys, cacheKey)
	}

	itemMap := make(map[string]*model.Supplier)
	err := c.cache.MultiGet(ctx, keys, itemMap)
	if err != nil {
		return nil, err
	}

	retMap := make(map[uint64]*model.Supplier)
	for _, id := range ids {
		val, ok := itemMap[c.GetSupplierCacheKey(id)]
		if ok {
			retMap[id] = val
		}
	}

	return retMap, nil
}

// Del delete cache
func (c *supplierCache) Del(ctx context.Context, id uint64) error {
	cacheKey := c.GetSupplierCacheKey(id)
	err := c.cache.Del(ctx, cacheKey)
	if err != nil {
		return err
	}
	return nil
}

// SetPlaceholder set placeholder value to cache
func (c *supplierCache) SetPlaceholder(ctx context.Context, id uint64) error {
	cacheKey := c.GetSupplierCacheKey(id)
	return c.cache.SetCacheWithNotFound(ctx, cacheKey)
}

// IsPlaceholderErr check if cache is placeholder error
func (c *supplierCache) IsPlaceholderErr(err error) bool {
	return errors.Is(err, cache.ErrPlaceholder)
}
