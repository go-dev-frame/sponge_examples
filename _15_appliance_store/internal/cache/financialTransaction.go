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
	financialTransactionCachePrefixKey = "financialTransaction:"
	// FinancialTransactionExpireTime expire time
	FinancialTransactionExpireTime = 5 * time.Minute
)

var _ FinancialTransactionCache = (*financialTransactionCache)(nil)

// FinancialTransactionCache cache interface
type FinancialTransactionCache interface {
	Set(ctx context.Context, id uint64, data *model.FinancialTransaction, duration time.Duration) error
	Get(ctx context.Context, id uint64) (*model.FinancialTransaction, error)
	MultiGet(ctx context.Context, ids []uint64) (map[uint64]*model.FinancialTransaction, error)
	MultiSet(ctx context.Context, data []*model.FinancialTransaction, duration time.Duration) error
	Del(ctx context.Context, id uint64) error
	SetPlaceholder(ctx context.Context, id uint64) error
	IsPlaceholderErr(err error) bool
}

// financialTransactionCache define a cache struct
type financialTransactionCache struct {
	cache cache.Cache
}

// NewFinancialTransactionCache new a cache
func NewFinancialTransactionCache(cacheType *database.CacheType) FinancialTransactionCache {
	jsonEncoding := encoding.JSONEncoding{}
	cachePrefix := ""

	cType := strings.ToLower(cacheType.CType)
	switch cType {
	case "redis":
		c := cache.NewRedisCache(cacheType.Rdb, cachePrefix, jsonEncoding, func() interface{} {
			return &model.FinancialTransaction{}
		})
		return &financialTransactionCache{cache: c}
	case "memory":
		c := cache.NewMemoryCache(cachePrefix, jsonEncoding, func() interface{} {
			return &model.FinancialTransaction{}
		})
		return &financialTransactionCache{cache: c}
	}

	return nil // no cache
}

// GetFinancialTransactionCacheKey cache key
func (c *financialTransactionCache) GetFinancialTransactionCacheKey(id uint64) string {
	return financialTransactionCachePrefixKey + utils.Uint64ToStr(id)
}

// Set write to cache
func (c *financialTransactionCache) Set(ctx context.Context, id uint64, data *model.FinancialTransaction, duration time.Duration) error {
	if data == nil || id == 0 {
		return nil
	}
	cacheKey := c.GetFinancialTransactionCacheKey(id)
	err := c.cache.Set(ctx, cacheKey, data, duration)
	if err != nil {
		return err
	}
	return nil
}

// Get cache value
func (c *financialTransactionCache) Get(ctx context.Context, id uint64) (*model.FinancialTransaction, error) {
	var data *model.FinancialTransaction
	cacheKey := c.GetFinancialTransactionCacheKey(id)
	err := c.cache.Get(ctx, cacheKey, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// MultiSet multiple set cache
func (c *financialTransactionCache) MultiSet(ctx context.Context, data []*model.FinancialTransaction, duration time.Duration) error {
	valMap := make(map[string]interface{})
	for _, v := range data {
		cacheKey := c.GetFinancialTransactionCacheKey(v.ID)
		valMap[cacheKey] = v
	}

	err := c.cache.MultiSet(ctx, valMap, duration)
	if err != nil {
		return err
	}

	return nil
}

// MultiGet multiple get cache, return key in map is id value
func (c *financialTransactionCache) MultiGet(ctx context.Context, ids []uint64) (map[uint64]*model.FinancialTransaction, error) {
	var keys []string
	for _, v := range ids {
		cacheKey := c.GetFinancialTransactionCacheKey(v)
		keys = append(keys, cacheKey)
	}

	itemMap := make(map[string]*model.FinancialTransaction)
	err := c.cache.MultiGet(ctx, keys, itemMap)
	if err != nil {
		return nil, err
	}

	retMap := make(map[uint64]*model.FinancialTransaction)
	for _, id := range ids {
		val, ok := itemMap[c.GetFinancialTransactionCacheKey(id)]
		if ok {
			retMap[id] = val
		}
	}

	return retMap, nil
}

// Del delete cache
func (c *financialTransactionCache) Del(ctx context.Context, id uint64) error {
	cacheKey := c.GetFinancialTransactionCacheKey(id)
	err := c.cache.Del(ctx, cacheKey)
	if err != nil {
		return err
	}
	return nil
}

// SetPlaceholder set placeholder value to cache
func (c *financialTransactionCache) SetPlaceholder(ctx context.Context, id uint64) error {
	cacheKey := c.GetFinancialTransactionCacheKey(id)
	return c.cache.SetCacheWithNotFound(ctx, cacheKey)
}

// IsPlaceholderErr check if cache is placeholder error
func (c *financialTransactionCache) IsPlaceholderErr(err error) bool {
	return errors.Is(err, cache.ErrPlaceholder)
}
