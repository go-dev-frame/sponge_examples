package dao

import (
	"context"
	"errors"

	"golang.org/x/sync/singleflight"
	"gorm.io/gorm"

	"github.com/go-dev-frame/sponge/pkg/logger"
	"github.com/go-dev-frame/sponge/pkg/sgorm/query"

	"store/internal/cache"
	"store/internal/database"
	"store/internal/model"
)

var _ CheckDetailDao = (*checkDetailDao)(nil)

// CheckDetailDao defining the dao interface
type CheckDetailDao interface {
	Create(ctx context.Context, table *model.CheckDetail) error
	DeleteByCheckID(ctx context.Context, checkID string) error
	UpdateByCheckID(ctx context.Context, table *model.CheckDetail) error
	GetByCheckID(ctx context.Context, checkID string) (*model.CheckDetail, error)
	GetByColumns(ctx context.Context, params *query.Params) ([]*model.CheckDetail, int64, error)

	CreateByTx(ctx context.Context, tx *gorm.DB, table *model.CheckDetail) (string, error)
	DeleteByTx(ctx context.Context, tx *gorm.DB, checkID string) error
	UpdateByTx(ctx context.Context, tx *gorm.DB, table *model.CheckDetail) error
}

type checkDetailDao struct {
	db    *gorm.DB
	cache cache.CheckDetailCache // if nil, the cache is not used.
	sfg   *singleflight.Group    // if cache is nil, the sfg is not used.
}

// NewCheckDetailDao creating the dao interface
func NewCheckDetailDao(db *gorm.DB, xCache cache.CheckDetailCache) CheckDetailDao {
	if xCache == nil {
		return &checkDetailDao{db: db}
	}
	return &checkDetailDao{
		db:    db,
		cache: xCache,
		sfg:   new(singleflight.Group),
	}
}

func (d *checkDetailDao) deleteCache(ctx context.Context, checkID string) error {
	if d.cache != nil {
		return d.cache.Del(ctx, checkID)
	}
	return nil
}

// Create a record, insert the record and the checkID value is written back to the table
func (d *checkDetailDao) Create(ctx context.Context, table *model.CheckDetail) error {
	return d.db.WithContext(ctx).Create(table).Error
}

// DeleteByCheckID delete a record by checkID
func (d *checkDetailDao) DeleteByCheckID(ctx context.Context, checkID string) error {
	err := d.db.WithContext(ctx).Where("check_id = ?", checkID).Delete(&model.CheckDetail{}).Error
	if err != nil {
		return err
	}

	// delete cache
	_ = d.deleteCache(ctx, checkID)

	return nil
}

// UpdateByCheckID update a record by checkID
func (d *checkDetailDao) UpdateByCheckID(ctx context.Context, table *model.CheckDetail) error {
	err := d.updateDataByCheckID(ctx, d.db, table)

	// delete cache
	_ = d.deleteCache(ctx, table.CheckID)

	return err
}

func (d *checkDetailDao) updateDataByCheckID(ctx context.Context, db *gorm.DB, table *model.CheckDetail) error {
	if table.CheckID == "" {
		return errors.New("checkID cannot be empty")
	}

	update := map[string]interface{}{}

	if table.CheckID != "" {
		update["check_id"] = table.CheckID
	}
	if table.SkuID != 0 {
		update["sku_id"] = table.SkuID
	}
	if table.SystemQty != 0 {
		update["system_qty"] = table.SystemQty
	}
	if table.ActualQty != 0 {
		update["actual_qty"] = table.ActualQty
	}

	return db.WithContext(ctx).Model(table).Updates(update).Error
}

// GetByCheckID get a record by checkID
func (d *checkDetailDao) GetByCheckID(ctx context.Context, checkID string) (*model.CheckDetail, error) {
	// no cache
	if d.cache == nil {
		record := &model.CheckDetail{}
		err := d.db.WithContext(ctx).Where("check_id = ?", checkID).First(record).Error
		return record, err
	}

	// get from cache
	record, err := d.cache.Get(ctx, checkID)
	if err == nil {
		return record, nil
	}

	// get from database
	if errors.Is(err, database.ErrCacheNotFound) {
		// for the same checkID, prevent high concurrent simultaneous access to database
		val, err, _ := d.sfg.Do(checkID, func() (interface{}, error) {

			table := &model.CheckDetail{}
			err = d.db.WithContext(ctx).Where("check_id = ?", checkID).First(table).Error
			if err != nil {
				// set placeholder cache to prevent cache penetration, default expiration time 10 minutes
				if errors.Is(err, database.ErrRecordNotFound) {
					if err = d.cache.SetPlaceholder(ctx, checkID); err != nil {
						logger.Warn("cache.SetPlaceholder error", logger.Err(err), logger.Any("checkID", checkID))
					}
					return nil, database.ErrRecordNotFound
				}
				return nil, err
			}
			// set cache
			if err = d.cache.Set(ctx, checkID, table, cache.CheckDetailExpireTime); err != nil {
				logger.Warn("cache.Set error", logger.Err(err), logger.Any("checkID", checkID))
			}
			return table, nil
		})
		if err != nil {
			return nil, err
		}
		table, ok := val.(*model.CheckDetail)
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
//	sort: sort fields, default is checkID backwards, you can add - sign before the field to indicate reverse order, no - sign to indicate ascending order, multiple fields separated by comma
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
func (d *checkDetailDao) GetByColumns(ctx context.Context, params *query.Params) ([]*model.CheckDetail, int64, error) {
	if params.Sort == "" {
		params.Sort = "-check_id"
	}
	queryStr, args, err := params.ConvertToGormConditions()
	if err != nil {
		return nil, 0, errors.New("query params error: " + err.Error())
	}

	var total int64
	if params.Sort != "ignore count" { // determine if count is required
		err = d.db.WithContext(ctx).Model(&model.CheckDetail{}).Where(queryStr, args...).Count(&total).Error
		if err != nil {
			return nil, 0, err
		}
		if total == 0 {
			return nil, total, nil
		}
	}

	records := []*model.CheckDetail{}
	order, limit, offset := params.ConvertToPage()
	err = d.db.WithContext(ctx).Order(order).Limit(limit).Offset(offset).Where(queryStr, args...).Find(&records).Error
	if err != nil {
		return nil, 0, err
	}

	return records, total, err
}

// CreateByTx create a record in the database using the provided transaction
func (d *checkDetailDao) CreateByTx(ctx context.Context, tx *gorm.DB, table *model.CheckDetail) (string, error) {
	err := tx.WithContext(ctx).Create(table).Error
	return table.CheckID, err
}

// DeleteByTx delete a record by checkID in the database using the provided transaction
func (d *checkDetailDao) DeleteByTx(ctx context.Context, tx *gorm.DB, checkID string) error {
	err := tx.WithContext(ctx).Where("check_id = ?", checkID).Delete(&model.CheckDetail{}).Error
	if err != nil {
		return err
	}

	// delete cache
	_ = d.deleteCache(ctx, checkID)

	return nil
}

// UpdateByTx update a record by checkID in the database using the provided transaction
func (d *checkDetailDao) UpdateByTx(ctx context.Context, tx *gorm.DB, table *model.CheckDetail) error {
	err := d.updateDataByCheckID(ctx, tx, table)

	// delete cache
	_ = d.deleteCache(ctx, table.CheckID)

	return err
}
