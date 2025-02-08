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
	productMediaCachePrefixKey = "productMedia:"
	// ProductMediaExpireTime expire time
	ProductMediaExpireTime = 5 * time.Minute
)

var _ ProductMediaCache = (*productMediaCache)(nil)

// ProductMediaCache cache interface
type ProductMediaCache interface {
	Set(ctx context.Context, id uint64, data *model.ProductMedia, duration time.Duration) error
	Get(ctx context.Context, id uint64) (*model.ProductMedia, error)
	MultiGet(ctx context.Context, ids []uint64) (map[uint64]*model.ProductMedia, error)
	MultiSet(ctx context.Context, data []*model.ProductMedia, duration time.Duration) error
	Del(ctx context.Context, id uint64) error
	SetPlaceholder(ctx context.Context, id uint64) error
	IsPlaceholderErr(err error) bool
}

// productMediaCache define a cache struct
type productMediaCache struct {
	cache cache.Cache
}

// NewProductMediaCache new a cache
func NewProductMediaCache(cacheType *database.CacheType) ProductMediaCache {
	jsonEncoding := encoding.JSONEncoding{}
	cachePrefix := ""

	cType := strings.ToLower(cacheType.CType)
	switch cType {
	case "redis":
		c := cache.NewRedisCache(cacheType.Rdb, cachePrefix, jsonEncoding, func() interface{} {
			return &model.ProductMedia{}
		})
		return &productMediaCache{cache: c}
	case "memory":
		c := cache.NewMemoryCache(cachePrefix, jsonEncoding, func() interface{} {
			return &model.ProductMedia{}
		})
		return &productMediaCache{cache: c}
	}

	return nil // no cache
}

// GetProductMediaCacheKey cache key
func (c *productMediaCache) GetProductMediaCacheKey(id uint64) string {
	return productMediaCachePrefixKey + utils.Uint64ToStr(id)
}

// Set write to cache
func (c *productMediaCache) Set(ctx context.Context, id uint64, data *model.ProductMedia, duration time.Duration) error {
	if data == nil || id == 0 {
		return nil
	}
	cacheKey := c.GetProductMediaCacheKey(id)
	err := c.cache.Set(ctx, cacheKey, data, duration)
	if err != nil {
		return err
	}
	return nil
}

// Get cache value
func (c *productMediaCache) Get(ctx context.Context, id uint64) (*model.ProductMedia, error) {
	var data *model.ProductMedia
	cacheKey := c.GetProductMediaCacheKey(id)
	err := c.cache.Get(ctx, cacheKey, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// MultiSet multiple set cache
func (c *productMediaCache) MultiSet(ctx context.Context, data []*model.ProductMedia, duration time.Duration) error {
	valMap := make(map[string]interface{})
	for _, v := range data {
		cacheKey := c.GetProductMediaCacheKey(v.ID)
		valMap[cacheKey] = v
	}

	err := c.cache.MultiSet(ctx, valMap, duration)
	if err != nil {
		return err
	}

	return nil
}

// MultiGet multiple get cache, return key in map is id value
func (c *productMediaCache) MultiGet(ctx context.Context, ids []uint64) (map[uint64]*model.ProductMedia, error) {
	var keys []string
	for _, v := range ids {
		cacheKey := c.GetProductMediaCacheKey(v)
		keys = append(keys, cacheKey)
	}

	itemMap := make(map[string]*model.ProductMedia)
	err := c.cache.MultiGet(ctx, keys, itemMap)
	if err != nil {
		return nil, err
	}

	retMap := make(map[uint64]*model.ProductMedia)
	for _, id := range ids {
		val, ok := itemMap[c.GetProductMediaCacheKey(id)]
		if ok {
			retMap[id] = val
		}
	}

	return retMap, nil
}

// Del delete cache
func (c *productMediaCache) Del(ctx context.Context, id uint64) error {
	cacheKey := c.GetProductMediaCacheKey(id)
	err := c.cache.Del(ctx, cacheKey)
	if err != nil {
		return err
	}
	return nil
}

// SetPlaceholder set placeholder value to cache
func (c *productMediaCache) SetPlaceholder(ctx context.Context, id uint64) error {
	cacheKey := c.GetProductMediaCacheKey(id)
	return c.cache.SetCacheWithNotFound(ctx, cacheKey)
}

// IsPlaceholderErr check if cache is placeholder error
func (c *productMediaCache) IsPlaceholderErr(err error) bool {
	return errors.Is(err, cache.ErrPlaceholder)
}
