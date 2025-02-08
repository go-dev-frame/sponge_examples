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

func newSupplierCache() *gotest.Cache {
	record1 := &model.Supplier{}
	record1.ID = 1
	record2 := &model.Supplier{}
	record2.ID = 2
	testData := map[string]interface{}{
		utils.Uint64ToStr(record1.ID): record1,
		utils.Uint64ToStr(record2.ID): record2,
	}

	c := gotest.NewCache(testData)
	c.ICache = NewSupplierCache(&database.CacheType{
		CType: "redis",
		Rdb:   c.RedisClient,
	})
	return c
}

func Test_supplierCache_Set(t *testing.T) {
	c := newSupplierCache()
	defer c.Close()

	record := c.TestDataSlice[0].(*model.Supplier)
	err := c.ICache.(SupplierCache).Set(c.Ctx, record.ID, record, time.Hour)
	if err != nil {
		t.Fatal(err)
	}

	// nil data
	err = c.ICache.(SupplierCache).Set(c.Ctx, 0, nil, time.Hour)
	assert.NoError(t, err)
}

func Test_supplierCache_Get(t *testing.T) {
	c := newSupplierCache()
	defer c.Close()

	record := c.TestDataSlice[0].(*model.Supplier)
	err := c.ICache.(SupplierCache).Set(c.Ctx, record.ID, record, time.Hour)
	if err != nil {
		t.Fatal(err)
	}

	got, err := c.ICache.(SupplierCache).Get(c.Ctx, record.ID)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, record, got)

	// zero key error
	_, err = c.ICache.(SupplierCache).Get(c.Ctx, 0)
	assert.Error(t, err)
}

func Test_supplierCache_MultiGet(t *testing.T) {
	c := newSupplierCache()
	defer c.Close()

	var testData []*model.Supplier
	for _, data := range c.TestDataSlice {
		testData = append(testData, data.(*model.Supplier))
	}

	err := c.ICache.(SupplierCache).MultiSet(c.Ctx, testData, time.Hour)
	if err != nil {
		t.Fatal(err)
	}

	got, err := c.ICache.(SupplierCache).MultiGet(c.Ctx, c.GetIDs())
	if err != nil {
		t.Fatal(err)
	}

	expected := c.GetTestData()
	for k, v := range expected {
		assert.Equal(t, got[utils.StrToUint64(k)], v.(*model.Supplier))
	}
}

func Test_supplierCache_MultiSet(t *testing.T) {
	c := newSupplierCache()
	defer c.Close()

	var testData []*model.Supplier
	for _, data := range c.TestDataSlice {
		testData = append(testData, data.(*model.Supplier))
	}

	err := c.ICache.(SupplierCache).MultiSet(c.Ctx, testData, time.Hour)
	if err != nil {
		t.Fatal(err)
	}
}

func Test_supplierCache_Del(t *testing.T) {
	c := newSupplierCache()
	defer c.Close()

	record := c.TestDataSlice[0].(*model.Supplier)
	err := c.ICache.(SupplierCache).Del(c.Ctx, record.ID)
	if err != nil {
		t.Fatal(err)
	}
}

func Test_supplierCache_SetCacheWithNotFound(t *testing.T) {
	c := newSupplierCache()
	defer c.Close()

	record := c.TestDataSlice[0].(*model.Supplier)
	err := c.ICache.(SupplierCache).SetPlaceholder(c.Ctx, record.ID)
	if err != nil {
		t.Fatal(err)
	}
	b := c.ICache.(SupplierCache).IsPlaceholderErr(err)
	t.Log(b)
}

func TestNewSupplierCache(t *testing.T) {
	c := NewSupplierCache(&database.CacheType{
		CType: "",
	})
	assert.Nil(t, c)
	c = NewSupplierCache(&database.CacheType{
		CType: "memory",
	})
	assert.NotNil(t, c)
	c = NewSupplierCache(&database.CacheType{
		CType: "redis",
	})
	assert.NotNil(t, c)
}
