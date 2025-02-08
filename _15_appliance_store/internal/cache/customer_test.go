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

func newCustomerCache() *gotest.Cache {
	record1 := &model.Customer{}
	record1.ID = 1
	record2 := &model.Customer{}
	record2.ID = 2
	testData := map[string]interface{}{
		utils.Uint64ToStr(record1.ID): record1,
		utils.Uint64ToStr(record2.ID): record2,
	}

	c := gotest.NewCache(testData)
	c.ICache = NewCustomerCache(&database.CacheType{
		CType: "redis",
		Rdb:   c.RedisClient,
	})
	return c
}

func Test_customerCache_Set(t *testing.T) {
	c := newCustomerCache()
	defer c.Close()

	record := c.TestDataSlice[0].(*model.Customer)
	err := c.ICache.(CustomerCache).Set(c.Ctx, record.ID, record, time.Hour)
	if err != nil {
		t.Fatal(err)
	}

	// nil data
	err = c.ICache.(CustomerCache).Set(c.Ctx, 0, nil, time.Hour)
	assert.NoError(t, err)
}

func Test_customerCache_Get(t *testing.T) {
	c := newCustomerCache()
	defer c.Close()

	record := c.TestDataSlice[0].(*model.Customer)
	err := c.ICache.(CustomerCache).Set(c.Ctx, record.ID, record, time.Hour)
	if err != nil {
		t.Fatal(err)
	}

	got, err := c.ICache.(CustomerCache).Get(c.Ctx, record.ID)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, record, got)

	// zero key error
	_, err = c.ICache.(CustomerCache).Get(c.Ctx, 0)
	assert.Error(t, err)
}

func Test_customerCache_MultiGet(t *testing.T) {
	c := newCustomerCache()
	defer c.Close()

	var testData []*model.Customer
	for _, data := range c.TestDataSlice {
		testData = append(testData, data.(*model.Customer))
	}

	err := c.ICache.(CustomerCache).MultiSet(c.Ctx, testData, time.Hour)
	if err != nil {
		t.Fatal(err)
	}

	got, err := c.ICache.(CustomerCache).MultiGet(c.Ctx, c.GetIDs())
	if err != nil {
		t.Fatal(err)
	}

	expected := c.GetTestData()
	for k, v := range expected {
		assert.Equal(t, got[utils.StrToUint64(k)], v.(*model.Customer))
	}
}

func Test_customerCache_MultiSet(t *testing.T) {
	c := newCustomerCache()
	defer c.Close()

	var testData []*model.Customer
	for _, data := range c.TestDataSlice {
		testData = append(testData, data.(*model.Customer))
	}

	err := c.ICache.(CustomerCache).MultiSet(c.Ctx, testData, time.Hour)
	if err != nil {
		t.Fatal(err)
	}
}

func Test_customerCache_Del(t *testing.T) {
	c := newCustomerCache()
	defer c.Close()

	record := c.TestDataSlice[0].(*model.Customer)
	err := c.ICache.(CustomerCache).Del(c.Ctx, record.ID)
	if err != nil {
		t.Fatal(err)
	}
}

func Test_customerCache_SetCacheWithNotFound(t *testing.T) {
	c := newCustomerCache()
	defer c.Close()

	record := c.TestDataSlice[0].(*model.Customer)
	err := c.ICache.(CustomerCache).SetPlaceholder(c.Ctx, record.ID)
	if err != nil {
		t.Fatal(err)
	}
	b := c.ICache.(CustomerCache).IsPlaceholderErr(err)
	t.Log(b)
}

func TestNewCustomerCache(t *testing.T) {
	c := NewCustomerCache(&database.CacheType{
		CType: "",
	})
	assert.Nil(t, c)
	c = NewCustomerCache(&database.CacheType{
		CType: "memory",
	})
	assert.NotNil(t, c)
	c = NewCustomerCache(&database.CacheType{
		CType: "redis",
	})
	assert.NotNil(t, c)
}
