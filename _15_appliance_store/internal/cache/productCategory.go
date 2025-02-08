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
	productCategoryCachePrefixKey = "productCategory:"
	// ProductCategoryExpireTime expire time
	ProductCategoryExpireTime = 5 * time.Minute
)

var _ ProductCategoryCache = (*productCategoryCache)(nil)

// ProductCategoryCache cache interface
type ProductCategoryCache interface {
	Set(ctx context.Context, id uint64, data *model.ProductCategory, duration time.Duration) error
	Get(ctx context.Context, id uint64) (*model.ProductCategory, error)
	MultiGet(ctx context.Context, ids []uint64) (map[uint64]*model.ProductCategory, error)
	MultiSet(ctx context.Context, data []*model.ProductCategory, duration time.Duration) error
	Del(ctx context.Context, id uint64) error
	SetPlaceholder(ctx context.Context, id uint64) error
	IsPlaceholderErr(err error) bool
}

// productCategoryCache define a cache struct
type productCategoryCache struct {
	cache cache.Cache
}

// NewProductCategoryCache new a cache
func NewProductCategoryCache(cacheType *database.CacheType) ProductCategoryCache {
	jsonEncoding := encoding.JSONEncoding{}
	cachePrefix := ""

	cType := strings.ToLower(cacheType.CType)
	switch cType {
	case "redis":
		c := cache.NewRedisCache(cacheType.Rdb, cachePrefix, jsonEncoding, func() interface{} {
			return &model.ProductCategory{}
		})
		return &productCategoryCache{cache: c}
	case "memory":
		c := cache.NewMemoryCache(cachePrefix, jsonEncoding, func() interface{} {
			return &model.ProductCategory{}
		})
		return &productCategoryCache{cache: c}
	}

	return nil // no cache
}

// GetProductCategoryCacheKey cache key
func (c *productCategoryCache) GetProductCategoryCacheKey(id uint64) string {
	return productCategoryCachePrefixKey + utils.Uint64ToStr(id)
}

// Set write to cache
func (c *productCategoryCache) Set(ctx context.Context, id uint64, data *model.ProductCategory, duration time.Duration) error {
	if data == nil || id == 0 {
		return nil
	}
	cacheKey := c.GetProductCategoryCacheKey(id)
	err := c.cache.Set(ctx, cacheKey, data, duration)
	if err != nil {
		return err
	}
	return nil
}

// Get cache value
func (c *productCategoryCache) Get(ctx context.Context, id uint64) (*model.ProductCategory, error) {
	var data *model.ProductCategory
	cacheKey := c.GetProductCategoryCacheKey(id)
	err := c.cache.Get(ctx, cacheKey, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// MultiSet multiple set cache
func (c *productCategoryCache) MultiSet(ctx context.Context, data []*model.ProductCategory, duration time.Duration) error {
	valMap := make(map[string]interface{})
	for _, v := range data {
		cacheKey := c.GetProductCategoryCacheKey(v.ID)
		valMap[cacheKey] = v
	}

	err := c.cache.MultiSet(ctx, valMap, duration)
	if err != nil {
		return err
	}

	return nil
}

// MultiGet multiple get cache, return key in map is id value
func (c *productCategoryCache) MultiGet(ctx context.Context, ids []uint64) (map[uint64]*model.ProductCategory, error) {
	var keys []string
	for _, v := range ids {
		cacheKey := c.GetProductCategoryCacheKey(v)
		keys = append(keys, cacheKey)
	}

	itemMap := make(map[string]*model.ProductCategory)
	err := c.cache.MultiGet(ctx, keys, itemMap)
	if err != nil {
		return nil, err
	}

	retMap := make(map[uint64]*model.ProductCategory)
	for _, id := range ids {
		val, ok := itemMap[c.GetProductCategoryCacheKey(id)]
		if ok {
			retMap[id] = val
		}
	}

	return retMap, nil
}

// Del delete cache
func (c *productCategoryCache) Del(ctx context.Context, id uint64) error {
	cacheKey := c.GetProductCategoryCacheKey(id)
	err := c.cache.Del(ctx, cacheKey)
	if err != nil {
		return err
	}
	return nil
}

// SetPlaceholder set placeholder value to cache
func (c *productCategoryCache) SetPlaceholder(ctx context.Context, id uint64) error {
	cacheKey := c.GetProductCategoryCacheKey(id)
	return c.cache.SetCacheWithNotFound(ctx, cacheKey)
}

// IsPlaceholderErr check if cache is placeholder error
func (c *productCategoryCache) IsPlaceholderErr(err error) bool {
	return errors.Is(err, cache.ErrPlaceholder)
}
