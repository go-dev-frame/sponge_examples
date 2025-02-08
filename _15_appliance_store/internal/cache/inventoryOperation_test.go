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

func newInventoryOperationCache() *gotest.Cache {
	record1 := &model.InventoryOperation{}
	record1.ID = 1
	record2 := &model.InventoryOperation{}
	record2.ID = 2
	testData := map[string]interface{}{
		utils.Uint64ToStr(record1.ID): record1,
		utils.Uint64ToStr(record2.ID): record2,
	}

	c := gotest.NewCache(testData)
	c.ICache = NewInventoryOperationCache(&database.CacheType{
		CType: "redis",
		Rdb:   c.RedisClient,
	})
	return c
}

func Test_inventoryOperationCache_Set(t *testing.T) {
	c := newInventoryOperationCache()
	defer c.Close()

	record := c.TestDataSlice[0].(*model.InventoryOperation)
	err := c.ICache.(InventoryOperationCache).Set(c.Ctx, record.ID, record, time.Hour)
	if err != nil {
		t.Fatal(err)
	}

	// nil data
	err = c.ICache.(InventoryOperationCache).Set(c.Ctx, 0, nil, time.Hour)
	assert.NoError(t, err)
}

func Test_inventoryOperationCache_Get(t *testing.T) {
	c := newInventoryOperationCache()
	defer c.Close()

	record := c.TestDataSlice[0].(*model.InventoryOperation)
	err := c.ICache.(InventoryOperationCache).Set(c.Ctx, record.ID, record, time.Hour)
	if err != nil {
		t.Fatal(err)
	}

	got, err := c.ICache.(InventoryOperationCache).Get(c.Ctx, record.ID)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, record, got)

	// zero key error
	_, err = c.ICache.(InventoryOperationCache).Get(c.Ctx, 0)
	assert.Error(t, err)
}

func Test_inventoryOperationCache_MultiGet(t *testing.T) {
	c := newInventoryOperationCache()
	defer c.Close()

	var testData []*model.InventoryOperation
	for _, data := range c.TestDataSlice {
		testData = append(testData, data.(*model.InventoryOperation))
	}

	err := c.ICache.(InventoryOperationCache).MultiSet(c.Ctx, testData, time.Hour)
	if err != nil {
		t.Fatal(err)
	}

	got, err := c.ICache.(InventoryOperationCache).MultiGet(c.Ctx, c.GetIDs())
	if err != nil {
		t.Fatal(err)
	}

	expected := c.GetTestData()
	for k, v := range expected {
		assert.Equal(t, got[utils.StrToUint64(k)], v.(*model.InventoryOperation))
	}
}

func Test_inventoryOperationCache_MultiSet(t *testing.T) {
	c := newInventoryOperationCache()
	defer c.Close()

	var testData []*model.InventoryOperation
	for _, data := range c.TestDataSlice {
		testData = append(testData, data.(*model.InventoryOperation))
	}

	err := c.ICache.(InventoryOperationCache).MultiSet(c.Ctx, testData, time.Hour)
	if err != nil {
		t.Fatal(err)
	}
}

func Test_inventoryOperationCache_Del(t *testing.T) {
	c := newInventoryOperationCache()
	defer c.Close()

	record := c.TestDataSlice[0].(*model.InventoryOperation)
	err := c.ICache.(InventoryOperationCache).Del(c.Ctx, record.ID)
	if err != nil {
		t.Fatal(err)
	}
}

func Test_inventoryOperationCache_SetCacheWithNotFound(t *testing.T) {
	c := newInventoryOperationCache()
	defer c.Close()

	record := c.TestDataSlice[0].(*model.InventoryOperation)
	err := c.ICache.(InventoryOperationCache).SetPlaceholder(c.Ctx, record.ID)
	if err != nil {
		t.Fatal(err)
	}
	b := c.ICache.(InventoryOperationCache).IsPlaceholderErr(err)
	t.Log(b)
}

func TestNewInventoryOperationCache(t *testing.T) {
	c := NewInventoryOperationCache(&database.CacheType{
		CType: "",
	})
	assert.Nil(t, c)
	c = NewInventoryOperationCache(&database.CacheType{
		CType: "memory",
	})
	assert.NotNil(t, c)
	c = NewInventoryOperationCache(&database.CacheType{
		CType: "redis",
	})
	assert.NotNil(t, c)
}
