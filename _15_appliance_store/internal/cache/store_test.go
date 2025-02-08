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

func newStoreCache() *gotest.Cache {
	record1 := &model.Store{}
	record1.ID = 1
	record2 := &model.Store{}
	record2.ID = 2
	testData := map[string]interface{}{
		utils.Uint64ToStr(record1.ID): record1,
		utils.Uint64ToStr(record2.ID): record2,
	}

	c := gotest.NewCache(testData)
	c.ICache = NewStoreCache(&database.CacheType{
		CType: "redis",
		Rdb:   c.RedisClient,
	})
	return c
}

func Test_storeCache_Set(t *testing.T) {
	c := newStoreCache()
	defer c.Close()

	record := c.TestDataSlice[0].(*model.Store)
	err := c.ICache.(StoreCache).Set(c.Ctx, record.ID, record, time.Hour)
	if err != nil {
		t.Fatal(err)
	}

	// nil data
	err = c.ICache.(StoreCache).Set(c.Ctx, 0, nil, time.Hour)
	assert.NoError(t, err)
}

func Test_storeCache_Get(t *testing.T) {
	c := newStoreCache()
	defer c.Close()

	record := c.TestDataSlice[0].(*model.Store)
	err := c.ICache.(StoreCache).Set(c.Ctx, record.ID, record, time.Hour)
	if err != nil {
		t.Fatal(err)
	}

	got, err := c.ICache.(StoreCache).Get(c.Ctx, record.ID)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, record, got)

	// zero key error
	_, err = c.ICache.(StoreCache).Get(c.Ctx, 0)
	assert.Error(t, err)
}

func Test_storeCache_MultiGet(t *testing.T) {
	c := newStoreCache()
	defer c.Close()

	var testData []*model.Store
	for _, data := range c.TestDataSlice {
		testData = append(testData, data.(*model.Store))
	}

	err := c.ICache.(StoreCache).MultiSet(c.Ctx, testData, time.Hour)
	if err != nil {
		t.Fatal(err)
	}

	got, err := c.ICache.(StoreCache).MultiGet(c.Ctx, c.GetIDs())
	if err != nil {
		t.Fatal(err)
	}

	expected := c.GetTestData()
	for k, v := range expected {
		assert.Equal(t, got[utils.StrToUint64(k)], v.(*model.Store))
	}
}

func Test_storeCache_MultiSet(t *testing.T) {
	c := newStoreCache()
	defer c.Close()

	var testData []*model.Store
	for _, data := range c.TestDataSlice {
		testData = append(testData, data.(*model.Store))
	}

	err := c.ICache.(StoreCache).MultiSet(c.Ctx, testData, time.Hour)
	if err != nil {
		t.Fatal(err)
	}
}

func Test_storeCache_Del(t *testing.T) {
	c := newStoreCache()
	defer c.Close()

	record := c.TestDataSlice[0].(*model.Store)
	err := c.ICache.(StoreCache).Del(c.Ctx, record.ID)
	if err != nil {
		t.Fatal(err)
	}
}

func Test_storeCache_SetCacheWithNotFound(t *testing.T) {
	c := newStoreCache()
	defer c.Close()

	record := c.TestDataSlice[0].(*model.Store)
	err := c.ICache.(StoreCache).SetPlaceholder(c.Ctx, record.ID)
	if err != nil {
		t.Fatal(err)
	}
	b := c.ICache.(StoreCache).IsPlaceholderErr(err)
	t.Log(b)
}

func TestNewStoreCache(t *testing.T) {
	c := NewStoreCache(&database.CacheType{
		CType: "",
	})
	assert.Nil(t, c)
	c = NewStoreCache(&database.CacheType{
		CType: "memory",
	})
	assert.NotNil(t, c)
	c = NewStoreCache(&database.CacheType{
		CType: "redis",
	})
	assert.NotNil(t, c)
}
