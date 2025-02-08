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
	promotionCachePrefixKey = "promotion:"
	// PromotionExpireTime expire time
	PromotionExpireTime = 5 * time.Minute
)

var _ PromotionCache = (*promotionCache)(nil)

// PromotionCache cache interface
type PromotionCache interface {
	Set(ctx context.Context, id uint64, data *model.Promotion, duration time.Duration) error
	Get(ctx context.Context, id uint64) (*model.Promotion, error)
	MultiGet(ctx context.Context, ids []uint64) (map[uint64]*model.Promotion, error)
	MultiSet(ctx context.Context, data []*model.Promotion, duration time.Duration) error
	Del(ctx context.Context, id uint64) error
	SetPlaceholder(ctx context.Context, id uint64) error
	IsPlaceholderErr(err error) bool
}

// promotionCache define a cache struct
type promotionCache struct {
	cache cache.Cache
}

// NewPromotionCache new a cache
func NewPromotionCache(cacheType *database.CacheType) PromotionCache {
	jsonEncoding := encoding.JSONEncoding{}
	cachePrefix := ""

	cType := strings.ToLower(cacheType.CType)
	switch cType {
	case "redis":
		c := cache.NewRedisCache(cacheType.Rdb, cachePrefix, jsonEncoding, func() interface{} {
			return &model.Promotion{}
		})
		return &promotionCache{cache: c}
	case "memory":
		c := cache.NewMemoryCache(cachePrefix, jsonEncoding, func() interface{} {
			return &model.Promotion{}
		})
		return &promotionCache{cache: c}
	}

	return nil // no cache
}

// GetPromotionCacheKey cache key
func (c *promotionCache) GetPromotionCacheKey(id uint64) string {
	return promotionCachePrefixKey + utils.Uint64ToStr(id)
}

// Set write to cache
func (c *promotionCache) Set(ctx context.Context, id uint64, data *model.Promotion, duration time.Duration) error {
	if data == nil || id == 0 {
		return nil
	}
	cacheKey := c.GetPromotionCacheKey(id)
	err := c.cache.Set(ctx, cacheKey, data, duration)
	if err != nil {
		return err
	}
	return nil
}

// Get cache value
func (c *promotionCache) Get(ctx context.Context, id uint64) (*model.Promotion, error) {
	var data *model.Promotion
	cacheKey := c.GetPromotionCacheKey(id)
	err := c.cache.Get(ctx, cacheKey, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// MultiSet multiple set cache
func (c *promotionCache) MultiSet(ctx context.Context, data []*model.Promotion, duration time.Duration) error {
	valMap := make(map[string]interface{})
	for _, v := range data {
		cacheKey := c.GetPromotionCacheKey(v.ID)
		valMap[cacheKey] = v
	}

	err := c.cache.MultiSet(ctx, valMap, duration)
	if err != nil {
		return err
	}

	return nil
}

// MultiGet multiple get cache, return key in map is id value
func (c *promotionCache) MultiGet(ctx context.Context, ids []uint64) (map[uint64]*model.Promotion, error) {
	var keys []string
	for _, v := range ids {
		cacheKey := c.GetPromotionCacheKey(v)
		keys = append(keys, cacheKey)
	}

	itemMap := make(map[string]*model.Promotion)
	err := c.cache.MultiGet(ctx, keys, itemMap)
	if err != nil {
		return nil, err
	}

	retMap := make(map[uint64]*model.Promotion)
	for _, id := range ids {
		val, ok := itemMap[c.GetPromotionCacheKey(id)]
		if ok {
			retMap[id] = val
		}
	}

	return retMap, nil
}

// Del delete cache
func (c *promotionCache) Del(ctx context.Context, id uint64) error {
	cacheKey := c.GetPromotionCacheKey(id)
	err := c.cache.Del(ctx, cacheKey)
	if err != nil {
		return err
	}
	return nil
}

// SetPlaceholder set placeholder value to cache
func (c *promotionCache) SetPlaceholder(ctx context.Context, id uint64) error {
	cacheKey := c.GetPromotionCacheKey(id)
	return c.cache.SetCacheWithNotFound(ctx, cacheKey)
}

// IsPlaceholderErr check if cache is placeholder error
func (c *promotionCache) IsPlaceholderErr(err error) bool {
	return errors.Is(err, cache.ErrPlaceholder)
}
