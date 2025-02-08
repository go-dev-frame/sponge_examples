package handler

import (
	"net/http"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"github.com/stretchr/testify/assert"
	"store/api/types"

	"github.com/go-dev-frame/sponge/pkg/gin/response"
	"github.com/go-dev-frame/sponge/pkg/gotest"
	"github.com/go-dev-frame/sponge/pkg/httpcli"
	"github.com/go-dev-frame/sponge/pkg/utils"

	storeV1 "store/api/store/v1"
	"store/internal/cache"
	"store/internal/dao"
	"store/internal/database"
	"store/internal/ecode"
	"store/internal/model"
)

func newEmployeePerformanceHandler() *gotest.Handler {
	testData := &model.EmployeePerformance{}
	testData.ID = 1
	// you can set the other fields of testData here, such as:
	//testData.CreatedAt = time.Now()
	//testData.UpdatedAt = testData.CreatedAt

	// init mock cache
	c := gotest.NewCache(map[string]interface{}{utils.Uint64ToStr(testData.ID): testData})
	c.ICache = cache.NewEmployeePerformanceCache(&database.CacheType{
		CType: "redis",
		Rdb:   c.RedisClient,
	})

	// init mock dao
	d := gotest.NewDao(c, testData)
	d.IDao = dao.NewEmployeePerformanceDao(d.DB, c.ICache.(cache.EmployeePerformanceCache))

	// init mock handler
	h := gotest.NewHandler(d, testData)
	h.IHandler = &employeePerformanceHandler{employeePerformanceDao: d.IDao.(dao.EmployeePerformanceDao)}
	iHandler := h.IHandler.(storeV1.EmployeePerformanceLogicer)

	testFns := []gotest.RouterInfo{
		{
			FuncName: "Create",
			Method:   http.MethodPost,
			Path:     "/employeePerformance",
			HandlerFunc: func(c *gin.Context) {
				req := &storeV1.CreateEmployeePerformanceRequest{}
				_ = c.ShouldBindJSON(req)
				_, err := iHandler.Create(c, req)
				if err != nil {
					response.Error(c, ecode.ErrCreateEmployeePerformance)
					return
				}
				response.Success(c)
			},
		},
		{
			FuncName: "DeleteByID",
			Method:   http.MethodDelete,
			Path:     "/employeePerformance/:id",
			HandlerFunc: func(c *gin.Context) {
				req := &storeV1.DeleteEmployeePerformanceByIDRequest{
					Id: utils.StrToUint64(c.Param("id")),
				}
				_, err := iHandler.DeleteByID(c, req)
				if err != nil {
					response.Error(c, ecode.ErrDeleteByIDEmployeePerformance)
					return
				}
				response.Success(c)
			},
		},
		{
			FuncName: "UpdateByID",
			Method:   http.MethodPut,
			Path:     "/employeePerformance/:id",
			HandlerFunc: func(c *gin.Context) {
				req := &storeV1.UpdateEmployeePerformanceByIDRequest{}
				_ = c.ShouldBindJSON(req)
				req.Id = utils.StrToUint64(c.Param("id"))
				_, err := iHandler.UpdateByID(c, req)
				if err != nil {
					response.Error(c, ecode.ErrUpdateByIDEmployeePerformance)
					return
				}
				response.Success(c)
			},
		},
		{
			FuncName: "GetByID",
			Method:   http.MethodGet,
			Path:     "/employeePerformance/:id",
			HandlerFunc: func(c *gin.Context) {
				req := &storeV1.GetEmployeePerformanceByIDRequest{
					Id: utils.StrToUint64(c.Param("id")),
				}
				_, err := iHandler.GetByID(c, req)
				if err != nil {
					response.Error(c, ecode.ErrGetByIDEmployeePerformance)
					return
				}
				response.Success(c)
			},
		},
		{
			FuncName: "List",
			Method:   http.MethodPost,
			Path:     "/employeePerformance/list",
			HandlerFunc: func(c *gin.Context) {
				req := &storeV1.ListEmployeePerformanceRequest{}
				_ = c.ShouldBindJSON(req)
				_, err := iHandler.List(c, req)
				if err != nil {
					response.Error(c, ecode.ErrListEmployeePerformance)
					return
				}
				response.Success(c)
			},
		},
	}

	h.GoRunHTTPServer(testFns)

	time.Sleep(time.Millisecond * 200)
	return h
}

func Test_employeePerformanceHandler_Create(t *testing.T) {
	h := newEmployeePerformanceHandler()
	defer h.Close()
	testData := &storeV1.CreateEmployeePerformanceRequest{}
	_ = copier.Copy(testData, h.TestData.(*model.EmployeePerformance))

	h.MockDao.SQLMock.ExpectBegin()
	args := h.MockDao.GetAnyArgs(h.TestData)
	h.MockDao.SQLMock.ExpectExec("INSERT INTO .*").
		WithArgs(args[:len(args)-1]...). // adjusted for the amount of test data
		WillReturnResult(sqlmock.NewResult(1, 1))
	h.MockDao.SQLMock.ExpectCommit()

	result := &httpcli.StdResult{}
	err := httpcli.Post(result, h.GetRequestURL("Create"), testData)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%+v", result)

}

func Test_employeePerformanceHandler_DeleteByID(t *testing.T) {
	h := newEmployeePerformanceHandler()
	defer h.Close()
	testData := h.TestData.(*model.EmployeePerformance)
	expectedSQLForDeletion := "DELETE .*"

	h.MockDao.SQLMock.ExpectBegin()
	h.MockDao.SQLMock.ExpectExec(expectedSQLForDeletion).
		WithArgs(testData.ID). // adjusted for the amount of test data
		WillReturnResult(sqlmock.NewResult(int64(testData.ID), 1))
	h.MockDao.SQLMock.ExpectCommit()

	result := &httpcli.StdResult{}
	err := httpcli.Delete(result, h.GetRequestURL("DeleteByID", testData.ID))
	if err != nil {
		t.Fatal(err)
	}
	if result.Code != 0 {
		t.Fatalf("%+v", result)
	}

	// zero id error test
	err = httpcli.Delete(result, h.GetRequestURL("DeleteByID", 0))
	assert.NoError(t, err)

	// delete error test
	err = httpcli.Delete(result, h.GetRequestURL("DeleteByID", 111))
	assert.NoError(t, err)
}

func Test_employeePerformanceHandler_UpdateByID(t *testing.T) {
	h := newEmployeePerformanceHandler()
	defer h.Close()
	testData := &storeV1.UpdateEmployeePerformanceByIDRequest{}
	_ = copier.Copy(testData, h.TestData.(*model.EmployeePerformance))
	testData.Id = h.TestData.(*model.EmployeePerformance).ID

	h.MockDao.SQLMock.ExpectBegin()
	h.MockDao.SQLMock.ExpectExec("UPDATE .*").
		WithArgs(h.MockDao.AnyTime, testData.Id). // adjusted for the amount of test data
		WillReturnResult(sqlmock.NewResult(int64(testData.Id), 1))
	h.MockDao.SQLMock.ExpectCommit()

	result := &httpcli.StdResult{}
	err := httpcli.Put(result, h.GetRequestURL("UpdateByID", testData.Id), testData)
	if err != nil {
		t.Fatal(err)
	}
	if result.Code != 0 {
		t.Fatalf("%+v", result)
	}

	// zero id error test
	err = httpcli.Put(result, h.GetRequestURL("UpdateByID", 0), testData)
	assert.NoError(t, err)

	// update error test
	err = httpcli.Put(result, h.GetRequestURL("UpdateByID", 111), testData)
	assert.NoError(t, err)
}

func Test_employeePerformanceHandler_GetByID(t *testing.T) {
	h := newEmployeePerformanceHandler()
	defer h.Close()
	testData := h.TestData.(*model.EmployeePerformance)

	// column names and corresponding data
	rows := sqlmock.NewRows([]string{"id"}).
		AddRow(testData.ID)

	h.MockDao.SQLMock.ExpectQuery("SELECT .*").
		WithArgs(testData.ID).
		WillReturnRows(rows)

	result := &httpcli.StdResult{}
	err := httpcli.Get(result, h.GetRequestURL("GetByID", testData.ID))
	if err != nil {
		t.Fatal(err)
	}
	if result.Code != 0 {
		t.Fatalf("%+v", result)
	}

	// zero id error test
	err = httpcli.Get(result, h.GetRequestURL("GetByID", 0))
	assert.NoError(t, err)

	// get error test
	err = httpcli.Get(result, h.GetRequestURL("GetByID", 111))
	assert.NoError(t, err)
}

func Test_employeePerformanceHandler_List(t *testing.T) {
	h := newEmployeePerformanceHandler()
	defer h.Close()
	testData := h.TestData.(*model.EmployeePerformance)

	// column names and corresponding data
	rows := sqlmock.NewRows([]string{"id"}).
		AddRow(testData.ID)

	h.MockDao.SQLMock.ExpectQuery("SELECT .*").WillReturnRows(rows)

	result := &httpcli.StdResult{}
	err := httpcli.Post(result, h.GetRequestURL("List"), &storeV1.ListEmployeePerformanceRequest{
		Params: &types.Params{
			Page:  0,
			Limit: 10,
			Sort:  "ignore count", // ignore test count
		}})
	if err != nil {
		t.Fatal(err)
	}
	if result.Code != 0 {
		t.Fatalf("%+v", result)
	}

	// nil params error test
	err = httpcli.Post(result, h.GetRequestURL("List"), &storeV1.ListEmployeePerformanceRequest{})
	assert.NoError(t, err)

	// get error test
	err = httpcli.Post(result, h.GetRequestURL("List"), &storeV1.ListEmployeePerformanceRequest{Params: &types.Params{
		Page:  0,
		Limit: 10,
	}})
	assert.NoError(t, err)
}

func TestNewEmployeePerformanceHandler(t *testing.T) {
	defer func() {
		recover()
	}()
	_ = NewEmployeePerformanceHandler()
}
