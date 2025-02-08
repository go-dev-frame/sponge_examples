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

func newAccountPayableCache() *gotest.Cache {
	record1 := &model.AccountPayable{}
	record1.ID = 1
	record2 := &model.AccountPayable{}
	record2.ID = 2
	testData := map[string]interface{}{
		utils.Uint64ToStr(record1.ID): record1,
		utils.Uint64ToStr(record2.ID): record2,
	}

	c := gotest.NewCache(testData)
	c.ICache = NewAccountPayableCache(&database.CacheType{
		CType: "redis",
		Rdb:   c.RedisClient,
	})
	return c
}

func Test_accountPayableCache_Set(t *testing.T) {
	c := newAccountPayableCache()
	defer c.Close()

	record := c.TestDataSlice[0].(*model.AccountPayable)
	err := c.ICache.(AccountPayableCache).Set(c.Ctx, record.ID, record, time.Hour)
	if err != nil {
		t.Fatal(err)
	}

	// nil data
	err = c.ICache.(AccountPayableCache).Set(c.Ctx, 0, nil, time.Hour)
	assert.NoError(t, err)
}

func Test_accountPayableCache_Get(t *testing.T) {
	c := newAccountPayableCache()
	defer c.Close()

	record := c.TestDataSlice[0].(*model.AccountPayable)
	err := c.ICache.(AccountPayableCache).Set(c.Ctx, record.ID, record, time.Hour)
	if err != nil {
		t.Fatal(err)
	}

	got, err := c.ICache.(AccountPayableCache).Get(c.Ctx, record.ID)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, record, got)

	// zero key error
	_, err = c.ICache.(AccountPayableCache).Get(c.Ctx, 0)
	assert.Error(t, err)
}

func Test_accountPayableCache_MultiGet(t *testing.T) {
	c := newAccountPayableCache()
	defer c.Close()

	var testData []*model.AccountPayable
	for _, data := range c.TestDataSlice {
		testData = append(testData, data.(*model.AccountPayable))
	}

	err := c.ICache.(AccountPayableCache).MultiSet(c.Ctx, testData, time.Hour)
	if err != nil {
		t.Fatal(err)
	}

	got, err := c.ICache.(AccountPayableCache).MultiGet(c.Ctx, c.GetIDs())
	if err != nil {
		t.Fatal(err)
	}

	expected := c.GetTestData()
	for k, v := range expected {
		assert.Equal(t, got[utils.StrToUint64(k)], v.(*model.AccountPayable))
	}
}

func Test_accountPayableCache_MultiSet(t *testing.T) {
	c := newAccountPayableCache()
	defer c.Close()

	var testData []*model.AccountPayable
	for _, data := range c.TestDataSlice {
		testData = append(testData, data.(*model.AccountPayable))
	}

	err := c.ICache.(AccountPayableCache).MultiSet(c.Ctx, testData, time.Hour)
	if err != nil {
		t.Fatal(err)
	}
}

func Test_accountPayableCache_Del(t *testing.T) {
	c := newAccountPayableCache()
	defer c.Close()

	record := c.TestDataSlice[0].(*model.AccountPayable)
	err := c.ICache.(AccountPayableCache).Del(c.Ctx, record.ID)
	if err != nil {
		t.Fatal(err)
	}
}

func Test_accountPayableCache_SetCacheWithNotFound(t *testing.T) {
	c := newAccountPayableCache()
	defer c.Close()

	record := c.TestDataSlice[0].(*model.AccountPayable)
	err := c.ICache.(AccountPayableCache).SetPlaceholder(c.Ctx, record.ID)
	if err != nil {
		t.Fatal(err)
	}
	b := c.ICache.(AccountPayableCache).IsPlaceholderErr(err)
	t.Log(b)
}

func TestNewAccountPayableCache(t *testing.T) {
	c := NewAccountPayableCache(&database.CacheType{
		CType: "",
	})
	assert.Nil(t, c)
	c = NewAccountPayableCache(&database.CacheType{
		CType: "memory",
	})
	assert.NotNil(t, c)
	c = NewAccountPayableCache(&database.CacheType{
		CType: "redis",
	})
	assert.NotNil(t, c)
}
