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

func newProductSkuHandler() *gotest.Handler {
	testData := &model.ProductSku{}
	testData.ID = 1
	// you can set the other fields of testData here, such as:
	//testData.CreatedAt = time.Now()
	//testData.UpdatedAt = testData.CreatedAt

	// init mock cache
	c := gotest.NewCache(map[string]interface{}{utils.Uint64ToStr(testData.ID): testData})
	c.ICache = cache.NewProductSkuCache(&database.CacheType{
		CType: "redis",
		Rdb:   c.RedisClient,
	})

	// init mock dao
	d := gotest.NewDao(c, testData)
	d.IDao = dao.NewProductSkuDao(d.DB, c.ICache.(cache.ProductSkuCache))

	// init mock handler
	h := gotest.NewHandler(d, testData)
	h.IHandler = &productSkuHandler{productSkuDao: d.IDao.(dao.ProductSkuDao)}
	iHandler := h.IHandler.(storeV1.ProductSkuLogicer)

	testFns := []gotest.RouterInfo{
		{
			FuncName: "Create",
			Method:   http.MethodPost,
			Path:     "/productSku",
			HandlerFunc: func(c *gin.Context) {
				req := &storeV1.CreateProductSkuRequest{}
				_ = c.ShouldBindJSON(req)
				_, err := iHandler.Create(c, req)
				if err != nil {
					response.Error(c, ecode.ErrCreateProductSku)
					return
				}
				response.Success(c)
			},
		},
		{
			FuncName: "DeleteByID",
			Method:   http.MethodDelete,
			Path:     "/productSku/:id",
			HandlerFunc: func(c *gin.Context) {
				req := &storeV1.DeleteProductSkuByIDRequest{
					Id: utils.StrToUint64(c.Param("id")),
				}
				_, err := iHandler.DeleteByID(c, req)
				if err != nil {
					response.Error(c, ecode.ErrDeleteByIDProductSku)
					return
				}
				response.Success(c)
			},
		},
		{
			FuncName: "UpdateByID",
			Method:   http.MethodPut,
			Path:     "/productSku/:id",
			HandlerFunc: func(c *gin.Context) {
				req := &storeV1.UpdateProductSkuByIDRequest{}
				_ = c.ShouldBindJSON(req)
				req.Id = utils.StrToUint64(c.Param("id"))
				_, err := iHandler.UpdateByID(c, req)
				if err != nil {
					response.Error(c, ecode.ErrUpdateByIDProductSku)
					return
				}
				response.Success(c)
			},
		},
		{
			FuncName: "GetByID",
			Method:   http.MethodGet,
			Path:     "/productSku/:id",
			HandlerFunc: func(c *gin.Context) {
				req := &storeV1.GetProductSkuByIDRequest{
					Id: utils.StrToUint64(c.Param("id")),
				}
				_, err := iHandler.GetByID(c, req)
				if err != nil {
					response.Error(c, ecode.ErrGetByIDProductSku)
					return
				}
				response.Success(c)
			},
		},
		{
			FuncName: "List",
			Method:   http.MethodPost,
			Path:     "/productSku/list",
			HandlerFunc: func(c *gin.Context) {
				req := &storeV1.ListProductSkuRequest{}
				_ = c.ShouldBindJSON(req)
				_, err := iHandler.List(c, req)
				if err != nil {
					response.Error(c, ecode.ErrListProductSku)
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

func Test_productSkuHandler_Create(t *testing.T) {
	h := newProductSkuHandler()
	defer h.Close()
	testData := &storeV1.CreateProductSkuRequest{}
	_ = copier.Copy(testData, h.TestData.(*model.ProductSku))

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

func Test_productSkuHandler_DeleteByID(t *testing.T) {
	h := newProductSkuHandler()
	defer h.Close()
	testData := h.TestData.(*model.ProductSku)
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

func Test_productSkuHandler_UpdateByID(t *testing.T) {
	h := newProductSkuHandler()
	defer h.Close()
	testData := &storeV1.UpdateProductSkuByIDRequest{}
	_ = copier.Copy(testData, h.TestData.(*model.ProductSku))
	testData.Id = h.TestData.(*model.ProductSku).ID

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

func Test_productSkuHandler_GetByID(t *testing.T) {
	h := newProductSkuHandler()
	defer h.Close()
	testData := h.TestData.(*model.ProductSku)

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

func Test_productSkuHandler_List(t *testing.T) {
	h := newProductSkuHandler()
	defer h.Close()
	testData := h.TestData.(*model.ProductSku)

	// column names and corresponding data
	rows := sqlmock.NewRows([]string{"id"}).
		AddRow(testData.ID)

	h.MockDao.SQLMock.ExpectQuery("SELECT .*").WillReturnRows(rows)

	result := &httpcli.StdResult{}
	err := httpcli.Post(result, h.GetRequestURL("List"), &storeV1.ListProductSkuRequest{
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
	err = httpcli.Post(result, h.GetRequestURL("List"), &storeV1.ListProductSkuRequest{})
	assert.NoError(t, err)

	// get error test
	err = httpcli.Post(result, h.GetRequestURL("List"), &storeV1.ListProductSkuRequest{Params: &types.Params{
		Page:  0,
		Limit: 10,
	}})
	assert.NoError(t, err)
}

func TestNewProductSkuHandler(t *testing.T) {
	defer func() {
		recover()
	}()
	_ = NewProductSkuHandler()
}
