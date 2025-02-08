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

func newAuditLogCache() *gotest.Cache {
	record1 := &model.AuditLog{}
	record1.ID = 1
	record2 := &model.AuditLog{}
	record2.ID = 2
	testData := map[string]interface{}{
		utils.Uint64ToStr(record1.ID): record1,
		utils.Uint64ToStr(record2.ID): record2,
	}

	c := gotest.NewCache(testData)
	c.ICache = NewAuditLogCache(&database.CacheType{
		CType: "redis",
		Rdb:   c.RedisClient,
	})
	return c
}

func Test_auditLogCache_Set(t *testing.T) {
	c := newAuditLogCache()
	defer c.Close()

	record := c.TestDataSlice[0].(*model.AuditLog)
	err := c.ICache.(AuditLogCache).Set(c.Ctx, record.ID, record, time.Hour)
	if err != nil {
		t.Fatal(err)
	}

	// nil data
	err = c.ICache.(AuditLogCache).Set(c.Ctx, 0, nil, time.Hour)
	assert.NoError(t, err)
}

func Test_auditLogCache_Get(t *testing.T) {
	c := newAuditLogCache()
	defer c.Close()

	record := c.TestDataSlice[0].(*model.AuditLog)
	err := c.ICache.(AuditLogCache).Set(c.Ctx, record.ID, record, time.Hour)
	if err != nil {
		t.Fatal(err)
	}

	got, err := c.ICache.(AuditLogCache).Get(c.Ctx, record.ID)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, record, got)

	// zero key error
	_, err = c.ICache.(AuditLogCache).Get(c.Ctx, 0)
	assert.Error(t, err)
}

func Test_auditLogCache_MultiGet(t *testing.T) {
	c := newAuditLogCache()
	defer c.Close()

	var testData []*model.AuditLog
	for _, data := range c.TestDataSlice {
		testData = append(testData, data.(*model.AuditLog))
	}

	err := c.ICache.(AuditLogCache).MultiSet(c.Ctx, testData, time.Hour)
	if err != nil {
		t.Fatal(err)
	}

	got, err := c.ICache.(AuditLogCache).MultiGet(c.Ctx, c.GetIDs())
	if err != nil {
		t.Fatal(err)
	}

	expected := c.GetTestData()
	for k, v := range expected {
		assert.Equal(t, got[utils.StrToUint64(k)], v.(*model.AuditLog))
	}
}

func Test_auditLogCache_MultiSet(t *testing.T) {
	c := newAuditLogCache()
	defer c.Close()

	var testData []*model.AuditLog
	for _, data := range c.TestDataSlice {
		testData = append(testData, data.(*model.AuditLog))
	}

	err := c.ICache.(AuditLogCache).MultiSet(c.Ctx, testData, time.Hour)
	if err != nil {
		t.Fatal(err)
	}
}

func Test_auditLogCache_Del(t *testing.T) {
	c := newAuditLogCache()
	defer c.Close()

	record := c.TestDataSlice[0].(*model.AuditLog)
	err := c.ICache.(AuditLogCache).Del(c.Ctx, record.ID)
	if err != nil {
		t.Fatal(err)
	}
}

func Test_auditLogCache_SetCacheWithNotFound(t *testing.T) {
	c := newAuditLogCache()
	defer c.Close()

	record := c.TestDataSlice[0].(*model.AuditLog)
	err := c.ICache.(AuditLogCache).SetPlaceholder(c.Ctx, record.ID)
	if err != nil {
		t.Fatal(err)
	}
	b := c.ICache.(AuditLogCache).IsPlaceholderErr(err)
	t.Log(b)
}

func TestNewAuditLogCache(t *testing.T) {
	c := NewAuditLogCache(&database.CacheType{
		CType: "",
	})
	assert.Nil(t, c)
	c = NewAuditLogCache(&database.CacheType{
		CType: "memory",
	})
	assert.NotNil(t, c)
	c = NewAuditLogCache(&database.CacheType{
		CType: "redis",
	})
	assert.NotNil(t, c)
}
