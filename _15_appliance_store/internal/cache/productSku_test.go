package cache

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/go-dev-frame/sponge/pkg/gotest"
	"github.com/go-dev-frame/sponge/pkg/utils"

	"store/internal/database"
	"store/internal/model"
)

func newProductSkuCache() *gotest.Cache {
	record1 := &model.ProductSku{}
	record1.ID = 1
	record2 := &model.ProductSku{}
	record2.ID = 2
	testData := map[string]interface{}{
		utils.Uint64ToStr(record1.ID): record1,
		utils.Uint64ToStr(record2.ID): record2,
	}

	c := gotest.NewCache(testData)
	c.ICache = NewProductSkuCache(&database.CacheType{
		CType: "redis",
		Rdb:   c.RedisClient,
	})
	return c
}

func Test_productSkuCache_Set(t *testing.T) {
	c := newProductSkuCache()
	defer c.Close()

	record := c.TestDataSlice[0].(*model.ProductSku)
	err := c.ICache.(ProductSkuCache).Set(c.Ctx, record.ID, record, time.Hour)
	if err != nil {
		t.Fatal(err)
	}

	// nil data
	err = c.ICache.(ProductSkuCache).Set(c.Ctx, 0, nil, time.Hour)
	assert.NoError(t, err)
}

func Test_productSkuCache_Get(t *testing.T) {
	c := newProductSkuCache()
	defer c.Close()

	record := c.TestDataSlice[0].(*model.ProductSku)
	err := c.ICache.(ProductSkuCache).Set(c.Ctx, record.ID, record, time.Hour)
	if err != nil {
		t.Fatal(err)
	}

	got, err := c.ICache.(ProductSkuCache).Get(c.Ctx, record.ID)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, record, got)

	// zero key error
	_, err = c.ICache.(ProductSkuCache).Get(c.Ctx, 0)
	assert.Error(t, err)
}

func Test_productSkuCache_MultiGet(t *testing.T) {
	c := newProductSkuCache()
	defer c.Close()

	var testData []*model.ProductSku
	for _, data := range c.TestDataSlice {
		testData = append(testData, data.(*model.ProductSku))
	}

	err := c.ICache.(ProductSkuCache).MultiSet(c.Ctx, testData, time.Hour)
	if err != nil {
		t.Fatal(err)
	}

	got, err := c.ICache.(ProductSkuCache).MultiGet(c.Ctx, c.GetIDs())
	if err != nil {
		t.Fatal(err)
	}

	expected := c.GetTestData()
	for k, v := range expected {
		assert.Equal(t, got[utils.StrToUint64(k)], v.(*model.ProductSku))
	}
}

func Test_productSkuCache_MultiSet(t *testing.T) {
	c := newProductSkuCache()
	defer c.Close()

	var testData []*model.ProductSku
	for _, data := range c.TestDataSlice {
		testData = append(testData, data.(*model.ProductSku))
	}

	err := c.ICache.(ProductSkuCache).MultiSet(c.Ctx, testData, time.Hour)
	if err != nil {
		t.Fatal(err)
	}
}

func Test_productSkuCache_Del(t *testing.T) {
	c := newProductSkuCache()
	defer c.Close()

	record := c.TestDataSlice[0].(*model.ProductSku)
	err := c.ICache.(ProductSkuCache).Del(c.Ctx, record.ID)
	if err != nil {
		t.Fatal(err)
	}
}

func Test_productSkuCache_SetCacheWithNotFound(t *testing.T) {
	c := newProductSkuCache()
	defer c.Close()

	record := c.TestDataSlice[0].(*model.ProductSku)
	err := c.ICache.(ProductSkuCache).SetPlaceholder(c.Ctx, record.ID)
	if err != nil {
		t.Fatal(err)
	}
	b := c.ICache.(ProductSkuCache).IsPlaceholderErr(err)
	t.Log(b)
}

func TestNewProductSkuCache(t *testing.T) {
	c := NewProductSkuCache(&database.CacheType{
		CType: "",
	})
	assert.Nil(t, c)
	c = NewProductSkuCache(&database.CacheType{
		CType: "memory",
	})
	assert.NotNil(t, c)
	c = NewProductSkuCache(&database.CacheType{
		CType: "redis",
	})
	assert.NotNil(t, c)
}
