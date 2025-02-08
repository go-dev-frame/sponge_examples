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

func newSalesOrderItemCache() *gotest.Cache {
	record1 := &model.SalesOrderItem{}
	record1.ID = 1
	record2 := &model.SalesOrderItem{}
	record2.ID = 2
	testData := map[string]interface{}{
		utils.Uint64ToStr(record1.ID): record1,
		utils.Uint64ToStr(record2.ID): record2,
	}

	c := gotest.NewCache(testData)
	c.ICache = NewSalesOrderItemCache(&database.CacheType{
		CType: "redis",
		Rdb:   c.RedisClient,
	})
	return c
}

func Test_salesOrderItemCache_Set(t *testing.T) {
	c := newSalesOrderItemCache()
	defer c.Close()

	record := c.TestDataSlice[0].(*model.SalesOrderItem)
	err := c.ICache.(SalesOrderItemCache).Set(c.Ctx, record.ID, record, time.Hour)
	if err != nil {
		t.Fatal(err)
	}

	// nil data
	err = c.ICache.(SalesOrderItemCache).Set(c.Ctx, 0, nil, time.Hour)
	assert.NoError(t, err)
}

func Test_salesOrderItemCache_Get(t *testing.T) {
	c := newSalesOrderItemCache()
	defer c.Close()

	record := c.TestDataSlice[0].(*model.SalesOrderItem)
	err := c.ICache.(SalesOrderItemCache).Set(c.Ctx, record.ID, record, time.Hour)
	if err != nil {
		t.Fatal(err)
	}

	got, err := c.ICache.(SalesOrderItemCache).Get(c.Ctx, record.ID)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, record, got)

	// zero key error
	_, err = c.ICache.(SalesOrderItemCache).Get(c.Ctx, 0)
	assert.Error(t, err)
}

func Test_salesOrderItemCache_MultiGet(t *testing.T) {
	c := newSalesOrderItemCache()
	defer c.Close()

	var testData []*model.SalesOrderItem
	for _, data := range c.TestDataSlice {
		testData = append(testData, data.(*model.SalesOrderItem))
	}

	err := c.ICache.(SalesOrderItemCache).MultiSet(c.Ctx, testData, time.Hour)
	if err != nil {
		t.Fatal(err)
	}

	got, err := c.ICache.(SalesOrderItemCache).MultiGet(c.Ctx, c.GetIDs())
	if err != nil {
		t.Fatal(err)
	}

	expected := c.GetTestData()
	for k, v := range expected {
		assert.Equal(t, got[utils.StrToUint64(k)], v.(*model.SalesOrderItem))
	}
}

func Test_salesOrderItemCache_MultiSet(t *testing.T) {
	c := newSalesOrderItemCache()
	defer c.Close()

	var testData []*model.SalesOrderItem
	for _, data := range c.TestDataSlice {
		testData = append(testData, data.(*model.SalesOrderItem))
	}

	err := c.ICache.(SalesOrderItemCache).MultiSet(c.Ctx, testData, time.Hour)
	if err != nil {
		t.Fatal(err)
	}
}

func Test_salesOrderItemCache_Del(t *testing.T) {
	c := newSalesOrderItemCache()
	defer c.Close()

	record := c.TestDataSlice[0].(*model.SalesOrderItem)
	err := c.ICache.(SalesOrderItemCache).Del(c.Ctx, record.ID)
	if err != nil {
		t.Fatal(err)
	}
}

func Test_salesOrderItemCache_SetCacheWithNotFound(t *testing.T) {
	c := newSalesOrderItemCache()
	defer c.Close()

	record := c.TestDataSlice[0].(*model.SalesOrderItem)
	err := c.ICache.(SalesOrderItemCache).SetPlaceholder(c.Ctx, record.ID)
	if err != nil {
		t.Fatal(err)
	}
	b := c.ICache.(SalesOrderItemCache).IsPlaceholderErr(err)
	t.Log(b)
}

func TestNewSalesOrderItemCache(t *testing.T) {
	c := NewSalesOrderItemCache(&database.CacheType{
		CType: "",
	})
	assert.Nil(t, c)
	c = NewSalesOrderItemCache(&database.CacheType{
		CType: "memory",
	})
	assert.NotNil(t, c)
	c = NewSalesOrderItemCache(&database.CacheType{
		CType: "redis",
	})
	assert.NotNil(t, c)
}
