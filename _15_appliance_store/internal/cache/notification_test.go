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

func newNotificationCache() *gotest.Cache {
	record1 := &model.Notification{}
	record1.ID = 1
	record2 := &model.Notification{}
	record2.ID = 2
	testData := map[string]interface{}{
		utils.Uint64ToStr(record1.ID): record1,
		utils.Uint64ToStr(record2.ID): record2,
	}

	c := gotest.NewCache(testData)
	c.ICache = NewNotificationCache(&database.CacheType{
		CType: "redis",
		Rdb:   c.RedisClient,
	})
	return c
}

func Test_notificationCache_Set(t *testing.T) {
	c := newNotificationCache()
	defer c.Close()

	record := c.TestDataSlice[0].(*model.Notification)
	err := c.ICache.(NotificationCache).Set(c.Ctx, record.ID, record, time.Hour)
	if err != nil {
		t.Fatal(err)
	}

	// nil data
	err = c.ICache.(NotificationCache).Set(c.Ctx, 0, nil, time.Hour)
	assert.NoError(t, err)
}

func Test_notificationCache_Get(t *testing.T) {
	c := newNotificationCache()
	defer c.Close()

	record := c.TestDataSlice[0].(*model.Notification)
	err := c.ICache.(NotificationCache).Set(c.Ctx, record.ID, record, time.Hour)
	if err != nil {
		t.Fatal(err)
	}

	got, err := c.ICache.(NotificationCache).Get(c.Ctx, record.ID)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, record, got)

	// zero key error
	_, err = c.ICache.(NotificationCache).Get(c.Ctx, 0)
	assert.Error(t, err)
}

func Test_notificationCache_MultiGet(t *testing.T) {
	c := newNotificationCache()
	defer c.Close()

	var testData []*model.Notification
	for _, data := range c.TestDataSlice {
		testData = append(testData, data.(*model.Notification))
	}

	err := c.ICache.(NotificationCache).MultiSet(c.Ctx, testData, time.Hour)
	if err != nil {
		t.Fatal(err)
	}

	got, err := c.ICache.(NotificationCache).MultiGet(c.Ctx, c.GetIDs())
	if err != nil {
		t.Fatal(err)
	}

	expected := c.GetTestData()
	for k, v := range expected {
		assert.Equal(t, got[utils.StrToUint64(k)], v.(*model.Notification))
	}
}

func Test_notificationCache_MultiSet(t *testing.T) {
	c := newNotificationCache()
	defer c.Close()

	var testData []*model.Notification
	for _, data := range c.TestDataSlice {
		testData = append(testData, data.(*model.Notification))
	}

	err := c.ICache.(NotificationCache).MultiSet(c.Ctx, testData, time.Hour)
	if err != nil {
		t.Fatal(err)
	}
}

func Test_notificationCache_Del(t *testing.T) {
	c := newNotificationCache()
	defer c.Close()

	record := c.TestDataSlice[0].(*model.Notification)
	err := c.ICache.(NotificationCache).Del(c.Ctx, record.ID)
	if err != nil {
		t.Fatal(err)
	}
}

func Test_notificationCache_SetCacheWithNotFound(t *testing.T) {
	c := newNotificationCache()
	defer c.Close()

	record := c.TestDataSlice[0].(*model.Notification)
	err := c.ICache.(NotificationCache).SetPlaceholder(c.Ctx, record.ID)
	if err != nil {
		t.Fatal(err)
	}
	b := c.ICache.(NotificationCache).IsPlaceholderErr(err)
	t.Log(b)
}

func TestNewNotificationCache(t *testing.T) {
	c := NewNotificationCache(&database.CacheType{
		CType: "",
	})
	assert.Nil(t, c)
	c = NewNotificationCache(&database.CacheType{
		CType: "memory",
	})
	assert.NotNil(t, c)
	c = NewNotificationCache(&database.CacheType{
		CType: "redis",
	})
	assert.NotNil(t, c)
}
