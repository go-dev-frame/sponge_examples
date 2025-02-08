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

func newInventoryCache() *gotest.Cache {
	record1 := &model.Inventory{}
	record1.ID = 1
	record2 := &model.Inventory{}
	record2.ID = 2
	testData := map[string]interface{}{
		utils.Uint64ToStr(record1.ID): record1,
		utils.Uint64ToStr(record2.ID): record2,
	}

	c := gotest.NewCache(testData)
	c.ICache = NewInventoryCache(&database.CacheType{
		CType: "redis",
		Rdb:   c.RedisClient,
	})
	return c
}

func Test_inventoryCache_Set(t *testing.T) {
	c := newInventoryCache()
	defer c.Close()

	record := c.TestDataSlice[0].(*model.Inventory)
	err := c.ICache.(InventoryCache).Set(c.Ctx, record.ID, record, time.Hour)
	if err != nil {
		t.Fatal(err)
	}

	// nil data
	err = c.ICache.(InventoryCache).Set(c.Ctx, 0, nil, time.Hour)
	assert.NoError(t, err)
}

func Test_inventoryCache_Get(t *testing.T) {
	c := newInventoryCache()
	defer c.Close()

	record := c.TestDataSlice[0].(*model.Inventory)
	err := c.ICache.(InventoryCache).Set(c.Ctx, record.ID, record, time.Hour)
	if err != nil {
		t.Fatal(err)
	}

	got, err := c.ICache.(InventoryCache).Get(c.Ctx, record.ID)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, record, got)

	// zero key error
	_, err = c.ICache.(InventoryCache).Get(c.Ctx, 0)
	assert.Error(t, err)
}

func Test_inventoryCache_MultiGet(t *testing.T) {
	c := newInventoryCache()
	defer c.Close()

	var testData []*model.Inventory
	for _, data := range c.TestDataSlice {
		testData = append(testData, data.(*model.Inventory))
	}

	err := c.ICache.(InventoryCache).MultiSet(c.Ctx, testData, time.Hour)
	if err != nil {
		t.Fatal(err)
	}

	got, err := c.ICache.(InventoryCache).MultiGet(c.Ctx, c.GetIDs())
	if err != nil {
		t.Fatal(err)
	}

	expected := c.GetTestData()
	for k, v := range expected {
		assert.Equal(t, got[utils.StrToUint64(k)], v.(*model.Inventory))
	}
}

func Test_inventoryCache_MultiSet(t *testing.T) {
	c := newInventoryCache()
	defer c.Close()

	var testData []*model.Inventory
	for _, data := range c.TestDataSlice {
		testData = append(testData, data.(*model.Inventory))
	}

	err := c.ICache.(InventoryCache).MultiSet(c.Ctx, testData, time.Hour)
	if err != nil {
		t.Fatal(err)
	}
}

func Test_inventoryCache_Del(t *testing.T) {
	c := newInventoryCache()
	defer c.Close()

	record := c.TestDataSlice[0].(*model.Inventory)
	err := c.ICache.(InventoryCache).Del(c.Ctx, record.ID)
	if err != nil {
		t.Fatal(err)
	}
}

func Test_inventoryCache_SetCacheWithNotFound(t *testing.T) {
	c := newInventoryCache()
	defer c.Close()

	record := c.TestDataSlice[0].(*model.Inventory)
	err := c.ICache.(InventoryCache).SetPlaceholder(c.Ctx, record.ID)
	if err != nil {
		t.Fatal(err)
	}
	b := c.ICache.(InventoryCache).IsPlaceholderErr(err)
	t.Log(b)
}

func TestNewInventoryCache(t *testing.T) {
	c := NewInventoryCache(&database.CacheType{
		CType: "",
	})
	assert.Nil(t, c)
	c = NewInventoryCache(&database.CacheType{
		CType: "memory",
	})
	assert.NotNil(t, c)
	c = NewInventoryCache(&database.CacheType{
		CType: "redis",
	})
	assert.NotNil(t, c)
}
