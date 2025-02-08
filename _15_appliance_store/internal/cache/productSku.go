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
	productSkuCachePrefixKey = "productSku:"
	// ProductSkuExpireTime expire time
	ProductSkuExpireTime = 5 * time.Minute
)

var _ ProductSkuCache = (*productSkuCache)(nil)

// ProductSkuCache cache interface
type ProductSkuCache interface {
	Set(ctx context.Context, id uint64, data *model.ProductSku, duration time.Duration) error
	Get(ctx context.Context, id uint64) (*model.ProductSku, error)
	MultiGet(ctx context.Context, ids []uint64) (map[uint64]*model.ProductSku, error)
	MultiSet(ctx context.Context, data []*model.ProductSku, duration time.Duration) error
	Del(ctx context.Context, id uint64) error
	SetPlaceholder(ctx context.Context, id uint64) error
	IsPlaceholderErr(err error) bool
}

// productSkuCache define a cache struct
type productSkuCache struct {
	cache cache.Cache
}

// NewProductSkuCache new a cache
func NewProductSkuCache(cacheType *database.CacheType) ProductSkuCache {
	jsonEncoding := encoding.JSONEncoding{}
	cachePrefix := ""

	cType := strings.ToLower(cacheType.CType)
	switch cType {
	case "redis":
		c := cache.NewRedisCache(cacheType.Rdb, cachePrefix, jsonEncoding, func() interface{} {
			return &model.ProductSku{}
		})
		return &productSkuCache{cache: c}
	case "memory":
		c := cache.NewMemoryCache(cachePrefix, jsonEncoding, func() interface{} {
			return &model.ProductSku{}
		})
		return &productSkuCache{cache: c}
	}

	return nil // no cache
}

// GetProductSkuCacheKey cache key
func (c *productSkuCache) GetProductSkuCacheKey(id uint64) string {
	return productSkuCachePrefixKey + utils.Uint64ToStr(id)
}

// Set write to cache
func (c *productSkuCache) Set(ctx context.Context, id uint64, data *model.ProductSku, duration time.Duration) error {
	if data == nil || id == 0 {
		return nil
	}
	cacheKey := c.GetProductSkuCacheKey(id)
	err := c.cache.Set(ctx, cacheKey, data, duration)
	if err != nil {
		return err
	}
	return nil
}

// Get cache value
func (c *productSkuCache) Get(ctx context.Context, id uint64) (*model.ProductSku, error) {
	var data *model.ProductSku
	cacheKey := c.GetProductSkuCacheKey(id)
	err := c.cache.Get(ctx, cacheKey, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// MultiSet multiple set cache
func (c *productSkuCache) MultiSet(ctx context.Context, data []*model.ProductSku, duration time.Duration) error {
	valMap := make(map[string]interface{})
	for _, v := range data {
		cacheKey := c.GetProductSkuCacheKey(v.ID)
		valMap[cacheKey] = v
	}

	err := c.cache.MultiSet(ctx, valMap, duration)
	if err != nil {
		return err
	}

	return nil
}

// MultiGet multiple get cache, return key in map is id value
func (c *productSkuCache) MultiGet(ctx context.Context, ids []uint64) (map[uint64]*model.ProductSku, error) {
	var keys []string
	for _, v := range ids {
		cacheKey := c.GetProductSkuCacheKey(v)
		keys = append(keys, cacheKey)
	}

	itemMap := make(map[string]*model.ProductSku)
	err := c.cache.MultiGet(ctx, keys, itemMap)
	if err != nil {
		return nil, err
	}

	retMap := make(map[uint64]*model.ProductSku)
	for _, id := range ids {
		val, ok := itemMap[c.GetProductSkuCacheKey(id)]
		if ok {
			retMap[id] = val
		}
	}

	return retMap, nil
}

// Del delete cache
func (c *productSkuCache) Del(ctx context.Context, id uint64) error {
	cacheKey := c.GetProductSkuCacheKey(id)
	err := c.cache.Del(ctx, cacheKey)
	if err != nil {
		return err
	}
	return nil
}

// SetPlaceholder set placeholder value to cache
func (c *productSkuCache) SetPlaceholder(ctx context.Context, id uint64) error {
	cacheKey := c.GetProductSkuCacheKey(id)
	return c.cache.SetCacheWithNotFound(ctx, cacheKey)
}

// IsPlaceholderErr check if cache is placeholder error
func (c *productSkuCache) IsPlaceholderErr(err error) bool {
	return errors.Is(err, cache.ErrPlaceholder)
}
