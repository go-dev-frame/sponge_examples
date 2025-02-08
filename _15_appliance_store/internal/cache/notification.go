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
	notificationCachePrefixKey = "notification:"
	// NotificationExpireTime expire time
	NotificationExpireTime = 5 * time.Minute
)

var _ NotificationCache = (*notificationCache)(nil)

// NotificationCache cache interface
type NotificationCache interface {
	Set(ctx context.Context, id uint64, data *model.Notification, duration time.Duration) error
	Get(ctx context.Context, id uint64) (*model.Notification, error)
	MultiGet(ctx context.Context, ids []uint64) (map[uint64]*model.Notification, error)
	MultiSet(ctx context.Context, data []*model.Notification, duration time.Duration) error
	Del(ctx context.Context, id uint64) error
	SetPlaceholder(ctx context.Context, id uint64) error
	IsPlaceholderErr(err error) bool
}

// notificationCache define a cache struct
type notificationCache struct {
	cache cache.Cache
}

// NewNotificationCache new a cache
func NewNotificationCache(cacheType *database.CacheType) NotificationCache {
	jsonEncoding := encoding.JSONEncoding{}
	cachePrefix := ""

	cType := strings.ToLower(cacheType.CType)
	switch cType {
	case "redis":
		c := cache.NewRedisCache(cacheType.Rdb, cachePrefix, jsonEncoding, func() interface{} {
			return &model.Notification{}
		})
		return &notificationCache{cache: c}
	case "memory":
		c := cache.NewMemoryCache(cachePrefix, jsonEncoding, func() interface{} {
			return &model.Notification{}
		})
		return &notificationCache{cache: c}
	}

	return nil // no cache
}

// GetNotificationCacheKey cache key
func (c *notificationCache) GetNotificationCacheKey(id uint64) string {
	return notificationCachePrefixKey + utils.Uint64ToStr(id)
}

// Set write to cache
func (c *notificationCache) Set(ctx context.Context, id uint64, data *model.Notification, duration time.Duration) error {
	if data == nil || id == 0 {
		return nil
	}
	cacheKey := c.GetNotificationCacheKey(id)
	err := c.cache.Set(ctx, cacheKey, data, duration)
	if err != nil {
		return err
	}
	return nil
}

// Get cache value
func (c *notificationCache) Get(ctx context.Context, id uint64) (*model.Notification, error) {
	var data *model.Notification
	cacheKey := c.GetNotificationCacheKey(id)
	err := c.cache.Get(ctx, cacheKey, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// MultiSet multiple set cache
func (c *notificationCache) MultiSet(ctx context.Context, data []*model.Notification, duration time.Duration) error {
	valMap := make(map[string]interface{})
	for _, v := range data {
		cacheKey := c.GetNotificationCacheKey(v.ID)
		valMap[cacheKey] = v
	}

	err := c.cache.MultiSet(ctx, valMap, duration)
	if err != nil {
		return err
	}

	return nil
}

// MultiGet multiple get cache, return key in map is id value
func (c *notificationCache) MultiGet(ctx context.Context, ids []uint64) (map[uint64]*model.Notification, error) {
	var keys []string
	for _, v := range ids {
		cacheKey := c.GetNotificationCacheKey(v)
		keys = append(keys, cacheKey)
	}

	itemMap := make(map[string]*model.Notification)
	err := c.cache.MultiGet(ctx, keys, itemMap)
	if err != nil {
		return nil, err
	}

	retMap := make(map[uint64]*model.Notification)
	for _, id := range ids {
		val, ok := itemMap[c.GetNotificationCacheKey(id)]
		if ok {
			retMap[id] = val
		}
	}

	return retMap, nil
}

// Del delete cache
func (c *notificationCache) Del(ctx context.Context, id uint64) error {
	cacheKey := c.GetNotificationCacheKey(id)
	err := c.cache.Del(ctx, cacheKey)
	if err != nil {
		return err
	}
	return nil
}

// SetPlaceholder set placeholder value to cache
func (c *notificationCache) SetPlaceholder(ctx context.Context, id uint64) error {
	cacheKey := c.GetNotificationCacheKey(id)
	return c.cache.SetCacheWithNotFound(ctx, cacheKey)
}

// IsPlaceholderErr check if cache is placeholder error
func (c *notificationCache) IsPlaceholderErr(err error) bool {
	return errors.Is(err, cache.ErrPlaceholder)
}
