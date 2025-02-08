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

func newPromotionCache() *gotest.Cache {
	record1 := &model.Promotion{}
	record1.ID = 1
	record2 := &model.Promotion{}
	record2.ID = 2
	testData := map[string]interface{}{
		utils.Uint64ToStr(record1.ID): record1,
		utils.Uint64ToStr(record2.ID): record2,
	}

	c := gotest.NewCache(testData)
	c.ICache = NewPromotionCache(&database.CacheType{
		CType: "redis",
		Rdb:   c.RedisClient,
	})
	return c
}

func Test_promotionCache_Set(t *testing.T) {
	c := newPromotionCache()
	defer c.Close()

	record := c.TestDataSlice[0].(*model.Promotion)
	err := c.ICache.(PromotionCache).Set(c.Ctx, record.ID, record, time.Hour)
	if err != nil {
		t.Fatal(err)
	}

	// nil data
	err = c.ICache.(PromotionCache).Set(c.Ctx, 0, nil, time.Hour)
	assert.NoError(t, err)
}

func Test_promotionCache_Get(t *testing.T) {
	c := newPromotionCache()
	defer c.Close()

	record := c.TestDataSlice[0].(*model.Promotion)
	err := c.ICache.(PromotionCache).Set(c.Ctx, record.ID, record, time.Hour)
	if err != nil {
		t.Fatal(err)
	}

	got, err := c.ICache.(PromotionCache).Get(c.Ctx, record.ID)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, record, got)

	// zero key error
	_, err = c.ICache.(PromotionCache).Get(c.Ctx, 0)
	assert.Error(t, err)
}

func Test_promotionCache_MultiGet(t *testing.T) {
	c := newPromotionCache()
	defer c.Close()

	var testData []*model.Promotion
	for _, data := range c.TestDataSlice {
		testData = append(testData, data.(*model.Promotion))
	}

	err := c.ICache.(PromotionCache).MultiSet(c.Ctx, testData, time.Hour)
	if err != nil {
		t.Fatal(err)
	}

	got, err := c.ICache.(PromotionCache).MultiGet(c.Ctx, c.GetIDs())
	if err != nil {
		t.Fatal(err)
	}

	expected := c.GetTestData()
	for k, v := range expected {
		assert.Equal(t, got[utils.StrToUint64(k)], v.(*model.Promotion))
	}
}

func Test_promotionCache_MultiSet(t *testing.T) {
	c := newPromotionCache()
	defer c.Close()

	var testData []*model.Promotion
	for _, data := range c.TestDataSlice {
		testData = append(testData, data.(*model.Promotion))
	}

	err := c.ICache.(PromotionCache).MultiSet(c.Ctx, testData, time.Hour)
	if err != nil {
		t.Fatal(err)
	}
}

func Test_promotionCache_Del(t *testing.T) {
	c := newPromotionCache()
	defer c.Close()

	record := c.TestDataSlice[0].(*model.Promotion)
	err := c.ICache.(PromotionCache).Del(c.Ctx, record.ID)
	if err != nil {
		t.Fatal(err)
	}
}

func Test_promotionCache_SetCacheWithNotFound(t *testing.T) {
	c := newPromotionCache()
	defer c.Close()

	record := c.TestDataSlice[0].(*model.Promotion)
	err := c.ICache.(PromotionCache).SetPlaceholder(c.Ctx, record.ID)
	if err != nil {
		t.Fatal(err)
	}
	b := c.ICache.(PromotionCache).IsPlaceholderErr(err)
	t.Log(b)
}

func TestNewPromotionCache(t *testing.T) {
	c := NewPromotionCache(&database.CacheType{
		CType: "",
	})
	assert.Nil(t, c)
	c = NewPromotionCache(&database.CacheType{
		CType: "memory",
	})
	assert.NotNil(t, c)
	c = NewPromotionCache(&database.CacheType{
		CType: "redis",
	})
	assert.NotNil(t, c)
}
