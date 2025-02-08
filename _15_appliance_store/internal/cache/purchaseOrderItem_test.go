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

func newPurchaseOrderItemCache() *gotest.Cache {
	record1 := &model.PurchaseOrderItem{}
	record1.ID = 1
	record2 := &model.PurchaseOrderItem{}
	record2.ID = 2
	testData := map[string]interface{}{
		utils.Uint64ToStr(record1.ID): record1,
		utils.Uint64ToStr(record2.ID): record2,
	}

	c := gotest.NewCache(testData)
	c.ICache = NewPurchaseOrderItemCache(&database.CacheType{
		CType: "redis",
		Rdb:   c.RedisClient,
	})
	return c
}

func Test_purchaseOrderItemCache_Set(t *testing.T) {
	c := newPurchaseOrderItemCache()
	defer c.Close()

	record := c.TestDataSlice[0].(*model.PurchaseOrderItem)
	err := c.ICache.(PurchaseOrderItemCache).Set(c.Ctx, record.ID, record, time.Hour)
	if err != nil {
		t.Fatal(err)
	}

	// nil data
	err = c.ICache.(PurchaseOrderItemCache).Set(c.Ctx, 0, nil, time.Hour)
	assert.NoError(t, err)
}

func Test_purchaseOrderItemCache_Get(t *testing.T) {
	c := newPurchaseOrderItemCache()
	defer c.Close()

	record := c.TestDataSlice[0].(*model.PurchaseOrderItem)
	err := c.ICache.(PurchaseOrderItemCache).Set(c.Ctx, record.ID, record, time.Hour)
	if err != nil {
		t.Fatal(err)
	}

	got, err := c.ICache.(PurchaseOrderItemCache).Get(c.Ctx, record.ID)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, record, got)

	// zero key error
	_, err = c.ICache.(PurchaseOrderItemCache).Get(c.Ctx, 0)
	assert.Error(t, err)
}

func Test_purchaseOrderItemCache_MultiGet(t *testing.T) {
	c := newPurchaseOrderItemCache()
	defer c.Close()

	var testData []*model.PurchaseOrderItem
	for _, data := range c.TestDataSlice {
		testData = append(testData, data.(*model.PurchaseOrderItem))
	}

	err := c.ICache.(PurchaseOrderItemCache).MultiSet(c.Ctx, testData, time.Hour)
	if err != nil {
		t.Fatal(err)
	}

	got, err := c.ICache.(PurchaseOrderItemCache).MultiGet(c.Ctx, c.GetIDs())
	if err != nil {
		t.Fatal(err)
	}

	expected := c.GetTestData()
	for k, v := range expected {
		assert.Equal(t, got[utils.StrToUint64(k)], v.(*model.PurchaseOrderItem))
	}
}

func Test_purchaseOrderItemCache_MultiSet(t *testing.T) {
	c := newPurchaseOrderItemCache()
	defer c.Close()

	var testData []*model.PurchaseOrderItem
	for _, data := range c.TestDataSlice {
		testData = append(testData, data.(*model.PurchaseOrderItem))
	}

	err := c.ICache.(PurchaseOrderItemCache).MultiSet(c.Ctx, testData, time.Hour)
	if err != nil {
		t.Fatal(err)
	}
}

func Test_purchaseOrderItemCache_Del(t *testing.T) {
	c := newPurchaseOrderItemCache()
	defer c.Close()

	record := c.TestDataSlice[0].(*model.PurchaseOrderItem)
	err := c.ICache.(PurchaseOrderItemCache).Del(c.Ctx, record.ID)
	if err != nil {
		t.Fatal(err)
	}
}

func Test_purchaseOrderItemCache_SetCacheWithNotFound(t *testing.T) {
	c := newPurchaseOrderItemCache()
	defer c.Close()

	record := c.TestDataSlice[0].(*model.PurchaseOrderItem)
	err := c.ICache.(PurchaseOrderItemCache).SetPlaceholder(c.Ctx, record.ID)
	if err != nil {
		t.Fatal(err)
	}
	b := c.ICache.(PurchaseOrderItemCache).IsPlaceholderErr(err)
	t.Log(b)
}

func TestNewPurchaseOrderItemCache(t *testing.T) {
	c := NewPurchaseOrderItemCache(&database.CacheType{
		CType: "",
	})
	assert.Nil(t, c)
	c = NewPurchaseOrderItemCache(&database.CacheType{
		CType: "memory",
	})
	assert.NotNil(t, c)
	c = NewPurchaseOrderItemCache(&database.CacheType{
		CType: "redis",
	})
	assert.NotNil(t, c)
}
