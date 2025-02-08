package dao

import (
	"context"
	"errors"

	"golang.org/x/sync/singleflight"
	"gorm.io/gorm"

	"github.com/go-dev-frame/sponge/pkg/logger"
	"github.com/go-dev-frame/sponge/pkg/sgorm/query"
	"github.com/go-dev-frame/sponge/pkg/utils"

	"store/internal/cache"
	"store/internal/database"
	"store/internal/model"
)

var _ SalesOrderItemDao = (*salesOrderItemDao)(nil)

// SalesOrderItemDao defining the dao interface
type SalesOrderItemDao interface {
	Create(ctx context.Context, table *model.SalesOrderItem) error
	DeleteByID(ctx context.Context, id uint64) error
	UpdateByID(ctx context.Context, table *model.SalesOrderItem) error
	GetByID(ctx context.Context, id uint64) (*model.SalesOrderItem, error)
	GetByColumns(ctx context.Context, params *query.Params) ([]*model.SalesOrderItem, int64, error)

	CreateByTx(ctx context.Context, tx *gorm.DB, table *model.SalesOrderItem) (uint64, error)
	DeleteByTx(ctx context.Context, tx *gorm.DB, id uint64) error
	UpdateByTx(ctx context.Context, tx *gorm.DB, table *model.SalesOrderItem) error
}

type salesOrderItemDao struct {
	db    *gorm.DB
	cache cache.SalesOrderItemCache // if nil, the cache is not used.
	sfg   *singleflight.Group       // if cache is nil, the sfg is not used.
}

// NewSalesOrderItemDao creating the dao interface
func NewSalesOrderItemDao(db *gorm.DB, xCache cache.SalesOrderItemCache) SalesOrderItemDao {
	if xCache == nil {
		return &salesOrderItemDao{db: db}
	}
	return &salesOrderItemDao{
		db:    db,
		cache: xCache,
		sfg:   new(singleflight.Group),
	}
}

func (d *salesOrderItemDao) deleteCache(ctx context.Context, id uint64) error {
	if d.cache != nil {
		return d.cache.Del(ctx, id)
	}
	return nil
}

// Create a record, insert the record and the id value is written back to the table
func (d *salesOrderItemDao) Create(ctx context.Context, table *model.SalesOrderItem) error {
	return d.db.WithContext(ctx).Create(table).Error
}

// DeleteByID delete a record by id
func (d *salesOrderItemDao) DeleteByID(ctx context.Context, id uint64) error {
	err := d.db.WithContext(ctx).Where("id = ?", id).Delete(&model.SalesOrderItem{}).Error
	if err != nil {
		return err
	}

	// delete cache
	_ = d.deleteCache(ctx, id)

	return nil
}

// UpdateByID update a record by id
func (d *salesOrderItemDao) UpdateByID(ctx context.Context, table *model.SalesOrderItem) error {
	err := d.updateDataByID(ctx, d.db, table)

	// delete cache
	_ = d.deleteCache(ctx, table.ID)

	return err
}

func (d *salesOrderItemDao) updateDataByID(ctx context.Context, db *gorm.DB, table *model.SalesOrderItem) error {
	if table.ID < 1 {
		return errors.New("id cannot be 0")
	}

	update := map[string]interface{}{}

	if table.OrderID != "" {
		update["order_id"] = table.OrderID
	}
	if table.SkuID != 0 {
		update["sku_id"] = table.SkuID
	}
	if table.Quantity != 0 {
		update["quantity"] = table.Quantity
	}
	if table.Price.IsZero() == false {
		update["price"] = table.Price
	}
	if table.UnitPrice.IsZero() == false {
		update["unit_price"] = table.UnitPrice
	}

	return db.WithContext(ctx).Model(table).Updates(update).Error
}

// GetByID get a record by id
func (d *salesOrderItemDao) GetByID(ctx context.Context, id uint64) (*model.SalesOrderItem, error) {
	// no cache
	if d.cache == nil {
		record := &model.SalesOrderItem{}
		err := d.db.WithContext(ctx).Where("id = ?", id).First(record).Error
		return record, err
	}

	// get from cache
	record, err := d.cache.Get(ctx, id)
	if err == nil {
		return record, nil
	}

	// get from database
	if errors.Is(err, database.ErrCacheNotFound) {
		// for the same id, prevent high concurrent simultaneous access to database
		val, err, _ := d.sfg.Do(utils.Uint64ToStr(id), func() (interface{}, error) { //nolint
			table := &model.SalesOrderItem{}
			err = d.db.WithContext(ctx).Where("id = ?", id).First(table).Error
			if err != nil {
				if errors.Is(err, database.ErrRecordNotFound) {
					// set placeholder cache to prevent cache penetration, default expiration time 10 minutes
					if err = d.cache.SetPlaceholder(ctx, id); err != nil {
						logger.Warn("cache.SetPlaceholder error", logger.Err(err), logger.Any("id", id))
					}
					return nil, database.ErrRecordNotFound
				}
				return nil, err
			}
			// set cache
			if err = d.cache.Set(ctx, id, table, cache.SalesOrderItemExpireTime); err != nil {
				logger.Warn("cache.Set error", logger.Err(err), logger.Any("id", id))
			}
			return table, nil
		})
		if err != nil {
			return nil, err
		}
		table, ok := val.(*model.SalesOrderItem)
		if !ok {
			return nil, database.ErrRecordNotFound
		}
		return table, nil
	}

	if d.cache.IsPlaceholderErr(err) {
		return nil, database.ErrRecordNotFound
	}

	return nil, err
}

// GetByColumns get paging records by column information,
// Note: query performance degrades when table rows are very large because of the use of offset.
//
// params includes paging parameters and query parameters
// paging parameters (required):
//
//	page: page number, starting from 0
//	limit: lines per page
//	sort: sort fields, default is id backwards, you can add - sign before the field to indicate reverse order, no - sign to indicate ascending order, multiple fields separated by comma
//
// query parameters (not required):
//
//	name: column name
//	exp: expressions, which default is "=",  support =, !=, >, >=, <, <=, like, in, notin, isnull, isnotnull
//	value: column value, if exp=in, multiple values are separated by commas
//	logic: logical type, default value is "and", support &, and, ||, or
//
// example: search for a male over 20 years of age
//
//	params = &query.Params{
//	    Page: 0,
//	    Limit: 20,
//	    Columns: []query.Column{
//		{
//			Name:    "age",
//			Exp: ">",
//			Value:   20,
//		},
//		{
//			Name:  "gender",
//			Value: "male",
//		},
//	}
func (d *salesOrderItemDao) GetByColumns(ctx context.Context, params *query.Params) ([]*model.SalesOrderItem, int64, error) {
	queryStr, args, err := params.ConvertToGormConditions()
	if err != nil {
		return nil, 0, errors.New("query params error: " + err.Error())
	}

	var total int64
	if params.Sort != "ignore count" { // determine if count is required
		err = d.db.WithContext(ctx).Model(&model.SalesOrderItem{}).Where(queryStr, args...).Count(&total).Error
		if err != nil {
			return nil, 0, err
		}
		if total == 0 {
			return nil, total, nil
		}
	}

	records := []*model.SalesOrderItem{}
	order, limit, offset := params.ConvertToPage()
	err = d.db.WithContext(ctx).Order(order).Limit(limit).Offset(offset).Where(queryStr, args...).Find(&records).Error
	if err != nil {
		return nil, 0, err
	}

	return records, total, err
}

// CreateByTx create a record in the database using the provided transaction
func (d *salesOrderItemDao) CreateByTx(ctx context.Context, tx *gorm.DB, table *model.SalesOrderItem) (uint64, error) {
	err := tx.WithContext(ctx).Create(table).Error
	return table.ID, err
}

// DeleteByTx delete a record by id in the database using the provided transaction
func (d *salesOrderItemDao) DeleteByTx(ctx context.Context, tx *gorm.DB, id uint64) error {
	err := tx.WithContext(ctx).Where("id = ?", id).Delete(&model.SalesOrderItem{}).Error
	if err != nil {
		return err
	}

	// delete cache
	_ = d.deleteCache(ctx, id)

	return nil
}

// UpdateByTx update a record by id in the database using the provided transaction
func (d *salesOrderItemDao) UpdateByTx(ctx context.Context, tx *gorm.DB, table *model.SalesOrderItem) error {
	err := d.updateDataByID(ctx, tx, table)

	// delete cache
	_ = d.deleteCache(ctx, table.ID)

	return err
}
