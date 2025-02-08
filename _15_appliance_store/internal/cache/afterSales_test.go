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

func newAfterSalesCache() *gotest.Cache {
	record1 := &model.AfterSales{}
	record1.ID = 1
	record2 := &model.AfterSales{}
	record2.ID = 2
	testData := map[string]interface{}{
		utils.Uint64ToStr(record1.ID): record1,
		utils.Uint64ToStr(record2.ID): record2,
	}

	c := gotest.NewCache(testData)
	c.ICache = NewAfterSalesCache(&database.CacheType{
		CType: "redis",
		Rdb:   c.RedisClient,
	})
	return c
}

func Test_afterSalesCache_Set(t *testing.T) {
	c := newAfterSalesCache()
	defer c.Close()

	record := c.TestDataSlice[0].(*model.AfterSales)
	err := c.ICache.(AfterSalesCache).Set(c.Ctx, record.ID, record, time.Hour)
	if err != nil {
		t.Fatal(err)
	}

	// nil data
	err = c.ICache.(AfterSalesCache).Set(c.Ctx, 0, nil, time.Hour)
	assert.NoError(t, err)
}

func Test_afterSalesCache_Get(t *testing.T) {
	c := newAfterSalesCache()
	defer c.Close()

	record := c.TestDataSlice[0].(*model.AfterSales)
	err := c.ICache.(AfterSalesCache).Set(c.Ctx, record.ID, record, time.Hour)
	if err != nil {
		t.Fatal(err)
	}

	got, err := c.ICache.(AfterSalesCache).Get(c.Ctx, record.ID)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, record, got)

	// zero key error
	_, err = c.ICache.(AfterSalesCache).Get(c.Ctx, 0)
	assert.Error(t, err)
}

func Test_afterSalesCache_MultiGet(t *testing.T) {
	c := newAfterSalesCache()
	defer c.Close()

	var testData []*model.AfterSales
	for _, data := range c.TestDataSlice {
		testData = append(testData, data.(*model.AfterSales))
	}

	err := c.ICache.(AfterSalesCache).MultiSet(c.Ctx, testData, time.Hour)
	if err != nil {
		t.Fatal(err)
	}

	got, err := c.ICache.(AfterSalesCache).MultiGet(c.Ctx, c.GetIDs())
	if err != nil {
		t.Fatal(err)
	}

	expected := c.GetTestData()
	for k, v := range expected {
		assert.Equal(t, got[utils.StrToUint64(k)], v.(*model.AfterSales))
	}
}

func Test_afterSalesCache_MultiSet(t *testing.T) {
	c := newAfterSalesCache()
	defer c.Close()

	var testData []*model.AfterSales
	for _, data := range c.TestDataSlice {
		testData = append(testData, data.(*model.AfterSales))
	}

	err := c.ICache.(AfterSalesCache).MultiSet(c.Ctx, testData, time.Hour)
	if err != nil {
		t.Fatal(err)
	}
}

func Test_afterSalesCache_Del(t *testing.T) {
	c := newAfterSalesCache()
	defer c.Close()

	record := c.TestDataSlice[0].(*model.AfterSales)
	err := c.ICache.(AfterSalesCache).Del(c.Ctx, record.ID)
	if err != nil {
		t.Fatal(err)
	}
}

func Test_afterSalesCache_SetCacheWithNotFound(t *testing.T) {
	c := newAfterSalesCache()
	defer c.Close()

	record := c.TestDataSlice[0].(*model.AfterSales)
	err := c.ICache.(AfterSalesCache).SetPlaceholder(c.Ctx, record.ID)
	if err != nil {
		t.Fatal(err)
	}
	b := c.ICache.(AfterSalesCache).IsPlaceholderErr(err)
	t.Log(b)
}

func TestNewAfterSalesCache(t *testing.T) {
	c := NewAfterSalesCache(&database.CacheType{
		CType: "",
	})
	assert.Nil(t, c)
	c = NewAfterSalesCache(&database.CacheType{
		CType: "memory",
	})
	assert.NotNil(t, c)
	c = NewAfterSalesCache(&database.CacheType{
		CType: "redis",
	})
	assert.NotNil(t, c)
}
