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
	transferDetailCachePrefixKey = "transferDetail:"
	// TransferDetailExpireTime expire time
	TransferDetailExpireTime = 5 * time.Minute
)

var _ TransferDetailCache = (*transferDetailCache)(nil)

// TransferDetailCache cache interface
type TransferDetailCache interface {
	Set(ctx context.Context, transferID string, data *model.TransferDetail, duration time.Duration) error
	Get(ctx context.Context, transferID string) (*model.TransferDetail, error)
	MultiGet(ctx context.Context, transferIDs []string) (map[string]*model.TransferDetail, error)
	MultiSet(ctx context.Context, data []*model.TransferDetail, duration time.Duration) error
	Del(ctx context.Context, transferID string) error
	SetPlaceholder(ctx context.Context, transferID string) error
	IsPlaceholderErr(err error) bool
}

// transferDetailCache define a cache struct
type transferDetailCache struct {
	cache cache.Cache
}

// NewTransferDetailCache new a cache
func NewTransferDetailCache(cacheType *database.CacheType) TransferDetailCache {
	jsonEncoding := encoding.JSONEncoding{}
	cachePrefix := ""

	cType := strings.ToLower(cacheType.CType)
	switch cType {
	case "redis":
		c := cache.NewRedisCache(cacheType.Rdb, cachePrefix, jsonEncoding, func() interface{} {
			return &model.TransferDetail{}
		})
		return &transferDetailCache{cache: c}
	case "memory":
		c := cache.NewMemoryCache(cachePrefix, jsonEncoding, func() interface{} {
			return &model.TransferDetail{}
		})
		return &transferDetailCache{cache: c}
	}

	return nil // no cache
}

// GetTransferDetailCacheKey cache key
func (c *transferDetailCache) GetTransferDetailCacheKey(transferID string) string {
	return transferDetailCachePrefixKey + transferID
}

// Set write to cache
func (c *transferDetailCache) Set(ctx context.Context, transferID string, data *model.TransferDetail, duration time.Duration) error {
	if data == nil {
		return nil
	}
	cacheKey := c.GetTransferDetailCacheKey(transferID)
	err := c.cache.Set(ctx, cacheKey, data, duration)
	if err != nil {
		return err
	}
	return nil
}

// Get cache value
func (c *transferDetailCache) Get(ctx context.Context, transferID string) (*model.TransferDetail, error) {
	var data *model.TransferDetail
	cacheKey := c.GetTransferDetailCacheKey(transferID)
	err := c.cache.Get(ctx, cacheKey, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// MultiSet multiple set cache
func (c *transferDetailCache) MultiSet(ctx context.Context, data []*model.TransferDetail, duration time.Duration) error {
	valMap := make(map[string]interface{})
	for _, v := range data {
		cacheKey := c.GetTransferDetailCacheKey(v.TransferID)
		valMap[cacheKey] = v
	}

	err := c.cache.MultiSet(ctx, valMap, duration)
	if err != nil {
		return err
	}

	return nil
}

// MultiGet multiple get cache, return key in map is transferID value
func (c *transferDetailCache) MultiGet(ctx context.Context, transferIDs []string) (map[string]*model.TransferDetail, error) {
	var keys []string
	for _, v := range transferIDs {
		cacheKey := c.GetTransferDetailCacheKey(v)
		keys = append(keys, cacheKey)
	}

	itemMap := make(map[string]*model.TransferDetail)
	err := c.cache.MultiGet(ctx, keys, itemMap)
	if err != nil {
		return nil, err
	}

	retMap := make(map[string]*model.TransferDetail)
	for _, transferID := range transferIDs {
		val, ok := itemMap[c.GetTransferDetailCacheKey(transferID)]
		if ok {
			retMap[transferID] = val
		}
	}

	return retMap, nil
}

// Del delete cache
func (c *transferDetailCache) Del(ctx context.Context, transferID string) error {
	cacheKey := c.GetTransferDetailCacheKey(transferID)
	err := c.cache.Del(ctx, cacheKey)
	if err != nil {
		return err
	}
	return nil
}

// SetPlaceholder set placeholder value to cache
func (c *transferDetailCache) SetPlaceholder(ctx context.Context, transferID string) error {
	cacheKey := c.GetTransferDetailCacheKey(transferID)
	return c.cache.SetCacheWithNotFound(ctx, cacheKey)
}

// IsPlaceholderErr check if cache is placeholder error
func (c *transferDetailCache) IsPlaceholderErr(err error) bool {
	return errors.Is(err, cache.ErrPlaceholder)
}
