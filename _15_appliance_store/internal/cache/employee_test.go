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

func newEmployeeCache() *gotest.Cache {
	record1 := &model.Employee{}
	record1.ID = 1
	record2 := &model.Employee{}
	record2.ID = 2
	testData := map[string]interface{}{
		utils.Uint64ToStr(record1.ID): record1,
		utils.Uint64ToStr(record2.ID): record2,
	}

	c := gotest.NewCache(testData)
	c.ICache = NewEmployeeCache(&database.CacheType{
		CType: "redis",
		Rdb:   c.RedisClient,
	})
	return c
}

func Test_employeeCache_Set(t *testing.T) {
	c := newEmployeeCache()
	defer c.Close()

	record := c.TestDataSlice[0].(*model.Employee)
	err := c.ICache.(EmployeeCache).Set(c.Ctx, record.ID, record, time.Hour)
	if err != nil {
		t.Fatal(err)
	}

	// nil data
	err = c.ICache.(EmployeeCache).Set(c.Ctx, 0, nil, time.Hour)
	assert.NoError(t, err)
}

func Test_employeeCache_Get(t *testing.T) {
	c := newEmployeeCache()
	defer c.Close()

	record := c.TestDataSlice[0].(*model.Employee)
	err := c.ICache.(EmployeeCache).Set(c.Ctx, record.ID, record, time.Hour)
	if err != nil {
		t.Fatal(err)
	}

	got, err := c.ICache.(EmployeeCache).Get(c.Ctx, record.ID)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, record, got)

	// zero key error
	_, err = c.ICache.(EmployeeCache).Get(c.Ctx, 0)
	assert.Error(t, err)
}

func Test_employeeCache_MultiGet(t *testing.T) {
	c := newEmployeeCache()
	defer c.Close()

	var testData []*model.Employee
	for _, data := range c.TestDataSlice {
		testData = append(testData, data.(*model.Employee))
	}

	err := c.ICache.(EmployeeCache).MultiSet(c.Ctx, testData, time.Hour)
	if err != nil {
		t.Fatal(err)
	}

	got, err := c.ICache.(EmployeeCache).MultiGet(c.Ctx, c.GetIDs())
	if err != nil {
		t.Fatal(err)
	}

	expected := c.GetTestData()
	for k, v := range expected {
		assert.Equal(t, got[utils.StrToUint64(k)], v.(*model.Employee))
	}
}

func Test_employeeCache_MultiSet(t *testing.T) {
	c := newEmployeeCache()
	defer c.Close()

	var testData []*model.Employee
	for _, data := range c.TestDataSlice {
		testData = append(testData, data.(*model.Employee))
	}

	err := c.ICache.(EmployeeCache).MultiSet(c.Ctx, testData, time.Hour)
	if err != nil {
		t.Fatal(err)
	}
}

func Test_employeeCache_Del(t *testing.T) {
	c := newEmployeeCache()
	defer c.Close()

	record := c.TestDataSlice[0].(*model.Employee)
	err := c.ICache.(EmployeeCache).Del(c.Ctx, record.ID)
	if err != nil {
		t.Fatal(err)
	}
}

func Test_employeeCache_SetCacheWithNotFound(t *testing.T) {
	c := newEmployeeCache()
	defer c.Close()

	record := c.TestDataSlice[0].(*model.Employee)
	err := c.ICache.(EmployeeCache).SetPlaceholder(c.Ctx, record.ID)
	if err != nil {
		t.Fatal(err)
	}
	b := c.ICache.(EmployeeCache).IsPlaceholderErr(err)
	t.Log(b)
}

func TestNewEmployeeCache(t *testing.T) {
	c := NewEmployeeCache(&database.CacheType{
		CType: "",
	})
	assert.Nil(t, c)
	c = NewEmployeeCache(&database.CacheType{
		CType: "memory",
	})
	assert.NotNil(t, c)
	c = NewEmployeeCache(&database.CacheType{
		CType: "redis",
	})
	assert.NotNil(t, c)
}
