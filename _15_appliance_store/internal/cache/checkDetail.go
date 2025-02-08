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
	checkDetailCachePrefixKey = "checkDetail:"
	// CheckDetailExpireTime expire time
	CheckDetailExpireTime = 5 * time.Minute
)

var _ CheckDetailCache = (*checkDetailCache)(nil)

// CheckDetailCache cache interface
type CheckDetailCache interface {
	Set(ctx context.Context, checkID string, data *model.CheckDetail, duration time.Duration) error
	Get(ctx context.Context, checkID string) (*model.CheckDetail, error)
	MultiGet(ctx context.Context, checkIDs []string) (map[string]*model.CheckDetail, error)
	MultiSet(ctx context.Context, data []*model.CheckDetail, duration time.Duration) error
	Del(ctx context.Context, checkID string) error
	SetPlaceholder(ctx context.Context, checkID string) error
	IsPlaceholderErr(err error) bool
}

// checkDetailCache define a cache struct
type checkDetailCache struct {
	cache cache.Cache
}

// NewCheckDetailCache new a cache
func NewCheckDetailCache(cacheType *database.CacheType) CheckDetailCache {
	jsonEncoding := encoding.JSONEncoding{}
	cachePrefix := ""

	cType := strings.ToLower(cacheType.CType)
	switch cType {
	case "redis":
		c := cache.NewRedisCache(cacheType.Rdb, cachePrefix, jsonEncoding, func() interface{} {
			return &model.CheckDetail{}
		})
		return &checkDetailCache{cache: c}
	case "memory":
		c := cache.NewMemoryCache(cachePrefix, jsonEncoding, func() interface{} {
			return &model.CheckDetail{}
		})
		return &checkDetailCache{cache: c}
	}

	return nil // no cache
}

// GetCheckDetailCacheKey cache key
func (c *checkDetailCache) GetCheckDetailCacheKey(checkID string) string {
	return checkDetailCachePrefixKey + checkID
}

// Set write to cache
func (c *checkDetailCache) Set(ctx context.Context, checkID string, data *model.CheckDetail, duration time.Duration) error {
	if data == nil {
		return nil
	}
	cacheKey := c.GetCheckDetailCacheKey(checkID)
	err := c.cache.Set(ctx, cacheKey, data, duration)
	if err != nil {
		return err
	}
	return nil
}

// Get cache value
func (c *checkDetailCache) Get(ctx context.Context, checkID string) (*model.CheckDetail, error) {
	var data *model.CheckDetail
	cacheKey := c.GetCheckDetailCacheKey(checkID)
	err := c.cache.Get(ctx, cacheKey, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// MultiSet multiple set cache
func (c *checkDetailCache) MultiSet(ctx context.Context, data []*model.CheckDetail, duration time.Duration) error {
	valMap := make(map[string]interface{})
	for _, v := range data {
		cacheKey := c.GetCheckDetailCacheKey(v.CheckID)
		valMap[cacheKey] = v
	}

	err := c.cache.MultiSet(ctx, valMap, duration)
	if err != nil {
		return err
	}

	return nil
}

// MultiGet multiple get cache, return key in map is checkID value
func (c *checkDetailCache) MultiGet(ctx context.Context, checkIDs []string) (map[string]*model.CheckDetail, error) {
	var keys []string
	for _, v := range checkIDs {
		cacheKey := c.GetCheckDetailCacheKey(v)
		keys = append(keys, cacheKey)
	}

	itemMap := make(map[string]*model.CheckDetail)
	err := c.cache.MultiGet(ctx, keys, itemMap)
	if err != nil {
		return nil, err
	}

	retMap := make(map[string]*model.CheckDetail)
	for _, checkID := range checkIDs {
		val, ok := itemMap[c.GetCheckDetailCacheKey(checkID)]
		if ok {
			retMap[checkID] = val
		}
	}

	return retMap, nil
}

// Del delete cache
func (c *checkDetailCache) Del(ctx context.Context, checkID string) error {
	cacheKey := c.GetCheckDetailCacheKey(checkID)
	err := c.cache.Del(ctx, cacheKey)
	if err != nil {
		return err
	}
	return nil
}

// SetPlaceholder set placeholder value to cache
func (c *checkDetailCache) SetPlaceholder(ctx context.Context, checkID string) error {
	cacheKey := c.GetCheckDetailCacheKey(checkID)
	return c.cache.SetCacheWithNotFound(ctx, cacheKey)
}

// IsPlaceholderErr check if cache is placeholder error
func (c *checkDetailCache) IsPlaceholderErr(err error) bool {
	return errors.Is(err, cache.ErrPlaceholder)
}
