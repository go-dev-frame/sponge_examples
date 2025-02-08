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
	couponCachePrefixKey = "coupon:"
	// CouponExpireTime expire time
	CouponExpireTime = 5 * time.Minute
)

var _ CouponCache = (*couponCache)(nil)

// CouponCache cache interface
type CouponCache interface {
	Set(ctx context.Context, id string, data *model.Coupon, duration time.Duration) error
	Get(ctx context.Context, id string) (*model.Coupon, error)
	MultiGet(ctx context.Context, ids []string) (map[string]*model.Coupon, error)
	MultiSet(ctx context.Context, data []*model.Coupon, duration time.Duration) error
	Del(ctx context.Context, id string) error
	SetPlaceholder(ctx context.Context, id string) error
	IsPlaceholderErr(err error) bool
}

// couponCache define a cache struct
type couponCache struct {
	cache cache.Cache
}

// NewCouponCache new a cache
func NewCouponCache(cacheType *database.CacheType) CouponCache {
	jsonEncoding := encoding.JSONEncoding{}
	cachePrefix := ""

	cType := strings.ToLower(cacheType.CType)
	switch cType {
	case "redis":
		c := cache.NewRedisCache(cacheType.Rdb, cachePrefix, jsonEncoding, func() interface{} {
			return &model.Coupon{}
		})
		return &couponCache{cache: c}
	case "memory":
		c := cache.NewMemoryCache(cachePrefix, jsonEncoding, func() interface{} {
			return &model.Coupon{}
		})
		return &couponCache{cache: c}
	}

	return nil // no cache
}

// GetCouponCacheKey cache key
func (c *couponCache) GetCouponCacheKey(id string) string {
	return couponCachePrefixKey + id
}

// Set write to cache
func (c *couponCache) Set(ctx context.Context, id string, data *model.Coupon, duration time.Duration) error {
	if data == nil {
		return nil
	}
	cacheKey := c.GetCouponCacheKey(id)
	err := c.cache.Set(ctx, cacheKey, data, duration)
	if err != nil {
		return err
	}
	return nil
}

// Get cache value
func (c *couponCache) Get(ctx context.Context, id string) (*model.Coupon, error) {
	var data *model.Coupon
	cacheKey := c.GetCouponCacheKey(id)
	err := c.cache.Get(ctx, cacheKey, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// MultiSet multiple set cache
func (c *couponCache) MultiSet(ctx context.Context, data []*model.Coupon, duration time.Duration) error {
	valMap := make(map[string]interface{})
	for _, v := range data {
		cacheKey := c.GetCouponCacheKey(v.ID)
		valMap[cacheKey] = v
	}

	err := c.cache.MultiSet(ctx, valMap, duration)
	if err != nil {
		return err
	}

	return nil
}

// MultiGet multiple get cache, return key in map is id value
func (c *couponCache) MultiGet(ctx context.Context, ids []string) (map[string]*model.Coupon, error) {
	var keys []string
	for _, v := range ids {
		cacheKey := c.GetCouponCacheKey(v)
		keys = append(keys, cacheKey)
	}

	itemMap := make(map[string]*model.Coupon)
	err := c.cache.MultiGet(ctx, keys, itemMap)
	if err != nil {
		return nil, err
	}

	retMap := make(map[string]*model.Coupon)
	for _, id := range ids {
		val, ok := itemMap[c.GetCouponCacheKey(id)]
		if ok {
			retMap[id] = val
		}
	}

	return retMap, nil
}

// Del delete cache
func (c *couponCache) Del(ctx context.Context, id string) error {
	cacheKey := c.GetCouponCacheKey(id)
	err := c.cache.Del(ctx, cacheKey)
	if err != nil {
		return err
	}
	return nil
}

// SetPlaceholder set placeholder value to cache
func (c *couponCache) SetPlaceholder(ctx context.Context, id string) error {
	cacheKey := c.GetCouponCacheKey(id)
	return c.cache.SetCacheWithNotFound(ctx, cacheKey)
}

// IsPlaceholderErr check if cache is placeholder error
func (c *couponCache) IsPlaceholderErr(err error) bool {
	return errors.Is(err, cache.ErrPlaceholder)
}
