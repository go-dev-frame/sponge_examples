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

func newEmployeePerformanceCache() *gotest.Cache {
	record1 := &model.EmployeePerformance{}
	record1.ID = 1
	record2 := &model.EmployeePerformance{}
	record2.ID = 2
	testData := map[string]interface{}{
		utils.Uint64ToStr(record1.ID): record1,
		utils.Uint64ToStr(record2.ID): record2,
	}

	c := gotest.NewCache(testData)
	c.ICache = NewEmployeePerformanceCache(&database.CacheType{
		CType: "redis",
		Rdb:   c.RedisClient,
	})
	return c
}

func Test_employeePerformanceCache_Set(t *testing.T) {
	c := newEmployeePerformanceCache()
	defer c.Close()

	record := c.TestDataSlice[0].(*model.EmployeePerformance)
	err := c.ICache.(EmployeePerformanceCache).Set(c.Ctx, record.ID, record, time.Hour)
	if err != nil {
		t.Fatal(err)
	}

	// nil data
	err = c.ICache.(EmployeePerformanceCache).Set(c.Ctx, 0, nil, time.Hour)
	assert.NoError(t, err)
}

func Test_employeePerformanceCache_Get(t *testing.T) {
	c := newEmployeePerformanceCache()
	defer c.Close()

	record := c.TestDataSlice[0].(*model.EmployeePerformance)
	err := c.ICache.(EmployeePerformanceCache).Set(c.Ctx, record.ID, record, time.Hour)
	if err != nil {
		t.Fatal(err)
	}

	got, err := c.ICache.(EmployeePerformanceCache).Get(c.Ctx, record.ID)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, record, got)

	// zero key error
	_, err = c.ICache.(EmployeePerformanceCache).Get(c.Ctx, 0)
	assert.Error(t, err)
}

func Test_employeePerformanceCache_MultiGet(t *testing.T) {
	c := newEmployeePerformanceCache()
	defer c.Close()

	var testData []*model.EmployeePerformance
	for _, data := range c.TestDataSlice {
		testData = append(testData, data.(*model.EmployeePerformance))
	}

	err := c.ICache.(EmployeePerformanceCache).MultiSet(c.Ctx, testData, time.Hour)
	if err != nil {
		t.Fatal(err)
	}

	got, err := c.ICache.(EmployeePerformanceCache).MultiGet(c.Ctx, c.GetIDs())
	if err != nil {
		t.Fatal(err)
	}

	expected := c.GetTestData()
	for k, v := range expected {
		assert.Equal(t, got[utils.StrToUint64(k)], v.(*model.EmployeePerformance))
	}
}

func Test_employeePerformanceCache_MultiSet(t *testing.T) {
	c := newEmployeePerformanceCache()
	defer c.Close()

	var testData []*model.EmployeePerformance
	for _, data := range c.TestDataSlice {
		testData = append(testData, data.(*model.EmployeePerformance))
	}

	err := c.ICache.(EmployeePerformanceCache).MultiSet(c.Ctx, testData, time.Hour)
	if err != nil {
		t.Fatal(err)
	}
}

func Test_employeePerformanceCache_Del(t *testing.T) {
	c := newEmployeePerformanceCache()
	defer c.Close()

	record := c.TestDataSlice[0].(*model.EmployeePerformance)
	err := c.ICache.(EmployeePerformanceCache).Del(c.Ctx, record.ID)
	if err != nil {
		t.Fatal(err)
	}
}

func Test_employeePerformanceCache_SetCacheWithNotFound(t *testing.T) {
	c := newEmployeePerformanceCache()
	defer c.Close()

	record := c.TestDataSlice[0].(*model.EmployeePerformance)
	err := c.ICache.(EmployeePerformanceCache).SetPlaceholder(c.Ctx, record.ID)
	if err != nil {
		t.Fatal(err)
	}
	b := c.ICache.(EmployeePerformanceCache).IsPlaceholderErr(err)
	t.Log(b)
}

func TestNewEmployeePerformanceCache(t *testing.T) {
	c := NewEmployeePerformanceCache(&database.CacheType{
		CType: "",
	})
	assert.Nil(t, c)
	c = NewEmployeePerformanceCache(&database.CacheType{
		CType: "memory",
	})
	assert.NotNil(t, c)
	c = NewEmployeePerformanceCache(&database.CacheType{
		CType: "redis",
	})
	assert.NotNil(t, c)
}
