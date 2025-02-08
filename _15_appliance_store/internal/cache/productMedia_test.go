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

func newProductMediaCache() *gotest.Cache {
	record1 := &model.ProductMedia{}
	record1.ID = 1
	record2 := &model.ProductMedia{}
	record2.ID = 2
	testData := map[string]interface{}{
		utils.Uint64ToStr(record1.ID): record1,
		utils.Uint64ToStr(record2.ID): record2,
	}

	c := gotest.NewCache(testData)
	c.ICache = NewProductMediaCache(&database.CacheType{
		CType: "redis",
		Rdb:   c.RedisClient,
	})
	return c
}

func Test_productMediaCache_Set(t *testing.T) {
	c := newProductMediaCache()
	defer c.Close()

	record := c.TestDataSlice[0].(*model.ProductMedia)
	err := c.ICache.(ProductMediaCache).Set(c.Ctx, record.ID, record, time.Hour)
	if err != nil {
		t.Fatal(err)
	}

	// nil data
	err = c.ICache.(ProductMediaCache).Set(c.Ctx, 0, nil, time.Hour)
	assert.NoError(t, err)
}

func Test_productMediaCache_Get(t *testing.T) {
	c := newProductMediaCache()
	defer c.Close()

	record := c.TestDataSlice[0].(*model.ProductMedia)
	err := c.ICache.(ProductMediaCache).Set(c.Ctx, record.ID, record, time.Hour)
	if err != nil {
		t.Fatal(err)
	}

	got, err := c.ICache.(ProductMediaCache).Get(c.Ctx, record.ID)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, record, got)

	// zero key error
	_, err = c.ICache.(ProductMediaCache).Get(c.Ctx, 0)
	assert.Error(t, err)
}

func Test_productMediaCache_MultiGet(t *testing.T) {
	c := newProductMediaCache()
	defer c.Close()

	var testData []*model.ProductMedia
	for _, data := range c.TestDataSlice {
		testData = append(testData, data.(*model.ProductMedia))
	}

	err := c.ICache.(ProductMediaCache).MultiSet(c.Ctx, testData, time.Hour)
	if err != nil {
		t.Fatal(err)
	}

	got, err := c.ICache.(ProductMediaCache).MultiGet(c.Ctx, c.GetIDs())
	if err != nil {
		t.Fatal(err)
	}

	expected := c.GetTestData()
	for k, v := range expected {
		assert.Equal(t, got[utils.StrToUint64(k)], v.(*model.ProductMedia))
	}
}

func Test_productMediaCache_MultiSet(t *testing.T) {
	c := newProductMediaCache()
	defer c.Close()

	var testData []*model.ProductMedia
	for _, data := range c.TestDataSlice {
		testData = append(testData, data.(*model.ProductMedia))
	}

	err := c.ICache.(ProductMediaCache).MultiSet(c.Ctx, testData, time.Hour)
	if err != nil {
		t.Fatal(err)
	}
}

func Test_productMediaCache_Del(t *testing.T) {
	c := newProductMediaCache()
	defer c.Close()

	record := c.TestDataSlice[0].(*model.ProductMedia)
	err := c.ICache.(ProductMediaCache).Del(c.Ctx, record.ID)
	if err != nil {
		t.Fatal(err)
	}
}

func Test_productMediaCache_SetCacheWithNotFound(t *testing.T) {
	c := newProductMediaCache()
	defer c.Close()

	record := c.TestDataSlice[0].(*model.ProductMedia)
	err := c.ICache.(ProductMediaCache).SetPlaceholder(c.Ctx, record.ID)
	if err != nil {
		t.Fatal(err)
	}
	b := c.ICache.(ProductMediaCache).IsPlaceholderErr(err)
	t.Log(b)
}

func TestNewProductMediaCache(t *testing.T) {
	c := NewProductMediaCache(&database.CacheType{
		CType: "",
	})
	assert.Nil(t, c)
	c = NewProductMediaCache(&database.CacheType{
		CType: "memory",
	})
	assert.NotNil(t, c)
	c = NewProductMediaCache(&database.CacheType{
		CType: "redis",
	})
	assert.NotNil(t, c)
}
