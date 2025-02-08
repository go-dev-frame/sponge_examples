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
	transferOrderCachePrefixKey = "transferOrder:"
	// TransferOrderExpireTime expire time
	TransferOrderExpireTime = 5 * time.Minute
)

var _ TransferOrderCache = (*transferOrderCache)(nil)

// TransferOrderCache cache interface
type TransferOrderCache interface {
	Set(ctx context.Context, id string, data *model.TransferOrder, duration time.Duration) error
	Get(ctx context.Context, id string) (*model.TransferOrder, error)
	MultiGet(ctx context.Context, ids []string) (map[string]*model.TransferOrder, error)
	MultiSet(ctx context.Context, data []*model.TransferOrder, duration time.Duration) error
	Del(ctx context.Context, id string) error
	SetPlaceholder(ctx context.Context, id string) error
	IsPlaceholderErr(err error) bool
}

// transferOrderCache define a cache struct
type transferOrderCache struct {
	cache cache.Cache
}

// NewTransferOrderCache new a cache
func NewTransferOrderCache(cacheType *database.CacheType) TransferOrderCache {
	jsonEncoding := encoding.JSONEncoding{}
	cachePrefix := ""

	cType := strings.ToLower(cacheType.CType)
	switch cType {
	case "redis":
		c := cache.NewRedisCache(cacheType.Rdb, cachePrefix, jsonEncoding, func() interface{} {
			return &model.TransferOrder{}
		})
		return &transferOrderCache{cache: c}
	case "memory":
		c := cache.NewMemoryCache(cachePrefix, jsonEncoding, func() interface{} {
			return &model.TransferOrder{}
		})
		return &transferOrderCache{cache: c}
	}

	return nil // no cache
}

// GetTransferOrderCacheKey cache key
func (c *transferOrderCache) GetTransferOrderCacheKey(id string) string {
	return transferOrderCachePrefixKey + id
}

// Set write to cache
func (c *transferOrderCache) Set(ctx context.Context, id string, data *model.TransferOrder, duration time.Duration) error {
	if data == nil {
		return nil
	}
	cacheKey := c.GetTransferOrderCacheKey(id)
	err := c.cache.Set(ctx, cacheKey, data, duration)
	if err != nil {
		return err
	}
	return nil
}

// Get cache value
func (c *transferOrderCache) Get(ctx context.Context, id string) (*model.TransferOrder, error) {
	var data *model.TransferOrder
	cacheKey := c.GetTransferOrderCacheKey(id)
	err := c.cache.Get(ctx, cacheKey, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// MultiSet multiple set cache
func (c *transferOrderCache) MultiSet(ctx context.Context, data []*model.TransferOrder, duration time.Duration) error {
	valMap := make(map[string]interface{})
	for _, v := range data {
		cacheKey := c.GetTransferOrderCacheKey(v.ID)
		valMap[cacheKey] = v
	}

	err := c.cache.MultiSet(ctx, valMap, duration)
	if err != nil {
		return err
	}

	return nil
}

// MultiGet multiple get cache, return key in map is id value
func (c *transferOrderCache) MultiGet(ctx context.Context, ids []string) (map[string]*model.TransferOrder, error) {
	var keys []string
	for _, v := range ids {
		cacheKey := c.GetTransferOrderCacheKey(v)
		keys = append(keys, cacheKey)
	}

	itemMap := make(map[string]*model.TransferOrder)
	err := c.cache.MultiGet(ctx, keys, itemMap)
	if err != nil {
		return nil, err
	}

	retMap := make(map[string]*model.TransferOrder)
	for _, id := range ids {
		val, ok := itemMap[c.GetTransferOrderCacheKey(id)]
		if ok {
			retMap[id] = val
		}
	}

	return retMap, nil
}

// Del delete cache
func (c *transferOrderCache) Del(ctx context.Context, id string) error {
	cacheKey := c.GetTransferOrderCacheKey(id)
	err := c.cache.Del(ctx, cacheKey)
	if err != nil {
		return err
	}
	return nil
}

// SetPlaceholder set placeholder value to cache
func (c *transferOrderCache) SetPlaceholder(ctx context.Context, id string) error {
	cacheKey := c.GetTransferOrderCacheKey(id)
	return c.cache.SetCacheWithNotFound(ctx, cacheKey)
}

// IsPlaceholderErr check if cache is placeholder error
func (c *transferOrderCache) IsPlaceholderErr(err error) bool {
	return errors.Is(err, cache.ErrPlaceholder)
}
