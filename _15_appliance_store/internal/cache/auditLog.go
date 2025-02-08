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
	auditLogCachePrefixKey = "auditLog:"
	// AuditLogExpireTime expire time
	AuditLogExpireTime = 5 * time.Minute
)

var _ AuditLogCache = (*auditLogCache)(nil)

// AuditLogCache cache interface
type AuditLogCache interface {
	Set(ctx context.Context, id uint64, data *model.AuditLog, duration time.Duration) error
	Get(ctx context.Context, id uint64) (*model.AuditLog, error)
	MultiGet(ctx context.Context, ids []uint64) (map[uint64]*model.AuditLog, error)
	MultiSet(ctx context.Context, data []*model.AuditLog, duration time.Duration) error
	Del(ctx context.Context, id uint64) error
	SetPlaceholder(ctx context.Context, id uint64) error
	IsPlaceholderErr(err error) bool
}

// auditLogCache define a cache struct
type auditLogCache struct {
	cache cache.Cache
}

// NewAuditLogCache new a cache
func NewAuditLogCache(cacheType *database.CacheType) AuditLogCache {
	jsonEncoding := encoding.JSONEncoding{}
	cachePrefix := ""

	cType := strings.ToLower(cacheType.CType)
	switch cType {
	case "redis":
		c := cache.NewRedisCache(cacheType.Rdb, cachePrefix, jsonEncoding, func() interface{} {
			return &model.AuditLog{}
		})
		return &auditLogCache{cache: c}
	case "memory":
		c := cache.NewMemoryCache(cachePrefix, jsonEncoding, func() interface{} {
			return &model.AuditLog{}
		})
		return &auditLogCache{cache: c}
	}

	return nil // no cache
}

// GetAuditLogCacheKey cache key
func (c *auditLogCache) GetAuditLogCacheKey(id uint64) string {
	return auditLogCachePrefixKey + utils.Uint64ToStr(id)
}

// Set write to cache
func (c *auditLogCache) Set(ctx context.Context, id uint64, data *model.AuditLog, duration time.Duration) error {
	if data == nil || id == 0 {
		return nil
	}
	cacheKey := c.GetAuditLogCacheKey(id)
	err := c.cache.Set(ctx, cacheKey, data, duration)
	if err != nil {
		return err
	}
	return nil
}

// Get cache value
func (c *auditLogCache) Get(ctx context.Context, id uint64) (*model.AuditLog, error) {
	var data *model.AuditLog
	cacheKey := c.GetAuditLogCacheKey(id)
	err := c.cache.Get(ctx, cacheKey, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// MultiSet multiple set cache
func (c *auditLogCache) MultiSet(ctx context.Context, data []*model.AuditLog, duration time.Duration) error {
	valMap := make(map[string]interface{})
	for _, v := range data {
		cacheKey := c.GetAuditLogCacheKey(v.ID)
		valMap[cacheKey] = v
	}

	err := c.cache.MultiSet(ctx, valMap, duration)
	if err != nil {
		return err
	}

	return nil
}

// MultiGet multiple get cache, return key in map is id value
func (c *auditLogCache) MultiGet(ctx context.Context, ids []uint64) (map[uint64]*model.AuditLog, error) {
	var keys []string
	for _, v := range ids {
		cacheKey := c.GetAuditLogCacheKey(v)
		keys = append(keys, cacheKey)
	}

	itemMap := make(map[string]*model.AuditLog)
	err := c.cache.MultiGet(ctx, keys, itemMap)
	if err != nil {
		return nil, err
	}

	retMap := make(map[uint64]*model.AuditLog)
	for _, id := range ids {
		val, ok := itemMap[c.GetAuditLogCacheKey(id)]
		if ok {
			retMap[id] = val
		}
	}

	return retMap, nil
}

// Del delete cache
func (c *auditLogCache) Del(ctx context.Context, id uint64) error {
	cacheKey := c.GetAuditLogCacheKey(id)
	err := c.cache.Del(ctx, cacheKey)
	if err != nil {
		return err
	}
	return nil
}

// SetPlaceholder set placeholder value to cache
func (c *auditLogCache) SetPlaceholder(ctx context.Context, id uint64) error {
	cacheKey := c.GetAuditLogCacheKey(id)
	return c.cache.SetCacheWithNotFound(ctx, cacheKey)
}

// IsPlaceholderErr check if cache is placeholder error
func (c *auditLogCache) IsPlaceholderErr(err error) bool {
	return errors.Is(err, cache.ErrPlaceholder)
}
