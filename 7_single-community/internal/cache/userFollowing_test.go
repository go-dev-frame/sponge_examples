package cache

import (
	"testing"
	"time"

	"community/internal/model"

	"github.com/zhufuyi/sponge/pkg/gotest"
	"github.com/zhufuyi/sponge/pkg/utils"

	"github.com/stretchr/testify/assert"
)

func newUserFollowingCache() *gotest.Cache {
	record1 := &model.UserFollowing{}
	record1.ID = 1
	record2 := &model.UserFollowing{}
	record2.ID = 2
	testData := map[string]interface{}{
		utils.Uint64ToStr(record1.ID): record1,
		utils.Uint64ToStr(record2.ID): record2,
	}

	c := gotest.NewCache(testData)
	c.ICache = NewUserFollowingCache(&model.CacheType{
		CType: "redis",
		Rdb:   c.RedisClient,
	})
	return c
}

func Test_userFollowingCache_Set(t *testing.T) {
	c := newUserFollowingCache()
	defer c.Close()

	record := c.TestDataSlice[0].(*model.UserFollowing)
	err := c.ICache.(UserFollowingCache).Set(c.Ctx, record.ID, record, time.Hour)
	if err != nil {
		t.Fatal(err)
	}

	// nil data
	err = c.ICache.(UserFollowingCache).Set(c.Ctx, 0, nil, time.Hour)
	assert.NoError(t, err)
}

func Test_userFollowingCache_Get(t *testing.T) {
	c := newUserFollowingCache()
	defer c.Close()

	record := c.TestDataSlice[0].(*model.UserFollowing)
	err := c.ICache.(UserFollowingCache).Set(c.Ctx, record.ID, record, time.Hour)
	if err != nil {
		t.Fatal(err)
	}

	got, err := c.ICache.(UserFollowingCache).Get(c.Ctx, record.ID)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, record, got)

	// zero key error
	_, err = c.ICache.(UserFollowingCache).Get(c.Ctx, 0)
	assert.Error(t, err)
}

func Test_userFollowingCache_MultiGet(t *testing.T) {
	c := newUserFollowingCache()
	defer c.Close()

	var testData []*model.UserFollowing
	for _, data := range c.TestDataSlice {
		testData = append(testData, data.(*model.UserFollowing))
	}

	err := c.ICache.(UserFollowingCache).MultiSet(c.Ctx, testData, time.Hour)
	if err != nil {
		t.Fatal(err)
	}

	got, err := c.ICache.(UserFollowingCache).MultiGet(c.Ctx, c.GetIDs())
	if err != nil {
		t.Fatal(err)
	}

	expected := c.GetTestData()
	for k, v := range expected {
		assert.Equal(t, got[utils.StrToUint64(k)], v.(*model.UserFollowing))
	}
}

func Test_userFollowingCache_MultiSet(t *testing.T) {
	c := newUserFollowingCache()
	defer c.Close()

	var testData []*model.UserFollowing
	for _, data := range c.TestDataSlice {
		testData = append(testData, data.(*model.UserFollowing))
	}

	err := c.ICache.(UserFollowingCache).MultiSet(c.Ctx, testData, time.Hour)
	if err != nil {
		t.Fatal(err)
	}
}

func Test_userFollowingCache_Del(t *testing.T) {
	c := newUserFollowingCache()
	defer c.Close()

	record := c.TestDataSlice[0].(*model.UserFollowing)
	err := c.ICache.(UserFollowingCache).Del(c.Ctx, record.ID)
	if err != nil {
		t.Fatal(err)
	}
}

func Test_userFollowingCache_SetCacheWithNotFound(t *testing.T) {
	c := newUserFollowingCache()
	defer c.Close()

	record := c.TestDataSlice[0].(*model.UserFollowing)
	err := c.ICache.(UserFollowingCache).SetCacheWithNotFound(c.Ctx, record.ID)
	if err != nil {
		t.Fatal(err)
	}
}

func TestNewUserFollowingCache(t *testing.T) {
	c := NewUserFollowingCache(&model.CacheType{
		CType: "memory",
	})

	assert.NotNil(t, c)
}
