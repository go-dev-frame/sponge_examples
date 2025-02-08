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

func newFinancialTransactionCache() *gotest.Cache {
	record1 := &model.FinancialTransaction{}
	record1.ID = 1
	record2 := &model.FinancialTransaction{}
	record2.ID = 2
	testData := map[string]interface{}{
		utils.Uint64ToStr(record1.ID): record1,
		utils.Uint64ToStr(record2.ID): record2,
	}

	c := gotest.NewCache(testData)
	c.ICache = NewFinancialTransactionCache(&database.CacheType{
		CType: "redis",
		Rdb:   c.RedisClient,
	})
	return c
}

func Test_financialTransactionCache_Set(t *testing.T) {
	c := newFinancialTransactionCache()
	defer c.Close()

	record := c.TestDataSlice[0].(*model.FinancialTransaction)
	err := c.ICache.(FinancialTransactionCache).Set(c.Ctx, record.ID, record, time.Hour)
	if err != nil {
		t.Fatal(err)
	}

	// nil data
	err = c.ICache.(FinancialTransactionCache).Set(c.Ctx, 0, nil, time.Hour)
	assert.NoError(t, err)
}

func Test_financialTransactionCache_Get(t *testing.T) {
	c := newFinancialTransactionCache()
	defer c.Close()

	record := c.TestDataSlice[0].(*model.FinancialTransaction)
	err := c.ICache.(FinancialTransactionCache).Set(c.Ctx, record.ID, record, time.Hour)
	if err != nil {
		t.Fatal(err)
	}

	got, err := c.ICache.(FinancialTransactionCache).Get(c.Ctx, record.ID)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, record, got)

	// zero key error
	_, err = c.ICache.(FinancialTransactionCache).Get(c.Ctx, 0)
	assert.Error(t, err)
}

func Test_financialTransactionCache_MultiGet(t *testing.T) {
	c := newFinancialTransactionCache()
	defer c.Close()

	var testData []*model.FinancialTransaction
	for _, data := range c.TestDataSlice {
		testData = append(testData, data.(*model.FinancialTransaction))
	}

	err := c.ICache.(FinancialTransactionCache).MultiSet(c.Ctx, testData, time.Hour)
	if err != nil {
		t.Fatal(err)
	}

	got, err := c.ICache.(FinancialTransactionCache).MultiGet(c.Ctx, c.GetIDs())
	if err != nil {
		t.Fatal(err)
	}

	expected := c.GetTestData()
	for k, v := range expected {
		assert.Equal(t, got[utils.StrToUint64(k)], v.(*model.FinancialTransaction))
	}
}

func Test_financialTransactionCache_MultiSet(t *testing.T) {
	c := newFinancialTransactionCache()
	defer c.Close()

	var testData []*model.FinancialTransaction
	for _, data := range c.TestDataSlice {
		testData = append(testData, data.(*model.FinancialTransaction))
	}

	err := c.ICache.(FinancialTransactionCache).MultiSet(c.Ctx, testData, time.Hour)
	if err != nil {
		t.Fatal(err)
	}
}

func Test_financialTransactionCache_Del(t *testing.T) {
	c := newFinancialTransactionCache()
	defer c.Close()

	record := c.TestDataSlice[0].(*model.FinancialTransaction)
	err := c.ICache.(FinancialTransactionCache).Del(c.Ctx, record.ID)
	if err != nil {
		t.Fatal(err)
	}
}

func Test_financialTransactionCache_SetCacheWithNotFound(t *testing.T) {
	c := newFinancialTransactionCache()
	defer c.Close()

	record := c.TestDataSlice[0].(*model.FinancialTransaction)
	err := c.ICache.(FinancialTransactionCache).SetPlaceholder(c.Ctx, record.ID)
	if err != nil {
		t.Fatal(err)
	}
	b := c.ICache.(FinancialTransactionCache).IsPlaceholderErr(err)
	t.Log(b)
}

func TestNewFinancialTransactionCache(t *testing.T) {
	c := NewFinancialTransactionCache(&database.CacheType{
		CType: "",
	})
	assert.Nil(t, c)
	c = NewFinancialTransactionCache(&database.CacheType{
		CType: "memory",
	})
	assert.NotNil(t, c)
	c = NewFinancialTransactionCache(&database.CacheType{
		CType: "redis",
	})
	assert.NotNil(t, c)
}
