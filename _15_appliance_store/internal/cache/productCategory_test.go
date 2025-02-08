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

func newProductCategoryCache() *gotest.Cache {
	record1 := &model.ProductCategory{}
	record1.ID = 1
	record2 := &model.ProductCategory{}
	record2.ID = 2
	testData := map[string]interface{}{
		utils.Uint64ToStr(record1.ID): record1,
		utils.Uint64ToStr(record2.ID): record2,
	}

	c := gotest.NewCache(testData)
	c.ICache = NewProductCategoryCache(&database.CacheType{
		CType: "redis",
		Rdb:   c.RedisClient,
	})
	return c
}

func Test_productCategoryCache_Set(t *testing.T) {
	c := newProductCategoryCache()
	defer c.Close()

	record := c.TestDataSlice[0].(*model.ProductCategory)
	err := c.ICache.(ProductCategoryCache).Set(c.Ctx, record.ID, record, time.Hour)
	if err != nil {
		t.Fatal(err)
	}

	// nil data
	err = c.ICache.(ProductCategoryCache).Set(c.Ctx, 0, nil, time.Hour)
	assert.NoError(t, err)
}

func Test_productCategoryCache_Get(t *testing.T) {
	c := newProductCategoryCache()
	defer c.Close()

	record := c.TestDataSlice[0].(*model.ProductCategory)
	err := c.ICache.(ProductCategoryCache).Set(c.Ctx, record.ID, record, time.Hour)
	if err != nil {
		t.Fatal(err)
	}

	got, err := c.ICache.(ProductCategoryCache).Get(c.Ctx, record.ID)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, record, got)

	// zero key error
	_, err = c.ICache.(ProductCategoryCache).Get(c.Ctx, 0)
	assert.Error(t, err)
}

func Test_productCategoryCache_MultiGet(t *testing.T) {
	c := newProductCategoryCache()
	defer c.Close()

	var testData []*model.ProductCategory
	for _, data := range c.TestDataSlice {
		testData = append(testData, data.(*model.ProductCategory))
	}

	err := c.ICache.(ProductCategoryCache).MultiSet(c.Ctx, testData, time.Hour)
	if err != nil {
		t.Fatal(err)
	}

	got, err := c.ICache.(ProductCategoryCache).MultiGet(c.Ctx, c.GetIDs())
	if err != nil {
		t.Fatal(err)
	}

	expected := c.GetTestData()
	for k, v := range expected {
		assert.Equal(t, got[utils.StrToUint64(k)], v.(*model.ProductCategory))
	}
}

func Test_productCategoryCache_MultiSet(t *testing.T) {
	c := newProductCategoryCache()
	defer c.Close()

	var testData []*model.ProductCategory
	for _, data := range c.TestDataSlice {
		testData = append(testData, data.(*model.ProductCategory))
	}

	err := c.ICache.(ProductCategoryCache).MultiSet(c.Ctx, testData, time.Hour)
	if err != nil {
		t.Fatal(err)
	}
}

func Test_productCategoryCache_Del(t *testing.T) {
	c := newProductCategoryCache()
	defer c.Close()

	record := c.TestDataSlice[0].(*model.ProductCategory)
	err := c.ICache.(ProductCategoryCache).Del(c.Ctx, record.ID)
	if err != nil {
		t.Fatal(err)
	}
}

func Test_productCategoryCache_SetCacheWithNotFound(t *testing.T) {
	c := newProductCategoryCache()
	defer c.Close()

	record := c.TestDataSlice[0].(*model.ProductCategory)
	err := c.ICache.(ProductCategoryCache).SetPlaceholder(c.Ctx, record.ID)
	if err != nil {
		t.Fatal(err)
	}
	b := c.ICache.(ProductCategoryCache).IsPlaceholderErr(err)
	t.Log(b)
}

func TestNewProductCategoryCache(t *testing.T) {
	c := NewProductCategoryCache(&database.CacheType{
		CType: "",
	})
	assert.Nil(t, c)
	c = NewProductCategoryCache(&database.CacheType{
		CType: "memory",
	})
	assert.NotNil(t, c)
	c = NewProductCategoryCache(&database.CacheType{
		CType: "redis",
	})
	assert.NotNil(t, c)
}
