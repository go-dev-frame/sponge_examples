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

var _ TransferDetailDao = (*transferDetailDao)(nil)

// TransferDetailDao defining the dao interface
type TransferDetailDao interface {
	Create(ctx context.Context, table *model.TransferDetail) error
	DeleteByTransferID(ctx context.Context, transferID string) error
	UpdateByTransferID(ctx context.Context, table *model.TransferDetail) error
	GetByTransferID(ctx context.Context, transferID string) (*model.TransferDetail, error)
	GetByColumns(ctx context.Context, params *query.Params) ([]*model.TransferDetail, int64, error)

	CreateByTx(ctx context.Context, tx *gorm.DB, table *model.TransferDetail) (string, error)
	DeleteByTx(ctx context.Context, tx *gorm.DB, transferID string) error
	UpdateByTx(ctx context.Context, tx *gorm.DB, table *model.TransferDetail) error
}

type transferDetailDao struct {
	db    *gorm.DB
	cache cache.TransferDetailCache // if nil, the cache is not used.
	sfg   *singleflight.Group       // if cache is nil, the sfg is not used.
}

// NewTransferDetailDao creating the dao interface
func NewTransferDetailDao(db *gorm.DB, xCache cache.TransferDetailCache) TransferDetailDao {
	if xCache == nil {
		return &transferDetailDao{db: db}
	}
	return &transferDetailDao{
		db:    db,
		cache: xCache,
		sfg:   new(singleflight.Group),
	}
}

func (d *transferDetailDao) deleteCache(ctx context.Context, transferID string) error {
	if d.cache != nil {
		return d.cache.Del(ctx, transferID)
	}
	return nil
}

// Create a record, insert the record and the transferID value is written back to the table
func (d *transferDetailDao) Create(ctx context.Context, table *model.TransferDetail) error {
	return d.db.WithContext(ctx).Create(table).Error
}

// DeleteByTransferID delete a record by transferID
func (d *transferDetailDao) DeleteByTransferID(ctx context.Context, transferID string) error {
	err := d.db.WithContext(ctx).Where("transfer_id = ?", transferID).Delete(&model.TransferDetail{}).Error
	if err != nil {
		return err
	}

	// delete cache
	_ = d.deleteCache(ctx, transferID)

	return nil
}

// UpdateByTransferID update a record by transferID
func (d *transferDetailDao) UpdateByTransferID(ctx context.Context, table *model.TransferDetail) error {
	err := d.updateDataByTransferID(ctx, d.db, table)

	// delete cache
	_ = d.deleteCache(ctx, table.TransferID)

	return err
}

func (d *transferDetailDao) updateDataByTransferID(ctx context.Context, db *gorm.DB, table *model.TransferDetail) error {
	if table.TransferID == "" {
		return errors.New("transferID cannot be empty")
	}

	update := map[string]interface{}{}

	if table.TransferID != "" {
		update["transfer_id"] = table.TransferID
	}
	if table.SkuID != 0 {
		update["sku_id"] = table.SkuID
	}
	if table.Quantity != 0 {
		update["quantity"] = table.Quantity
	}

	return db.WithContext(ctx).Model(table).Updates(update).Error
}

// GetByTransferID get a record by transferID
func (d *transferDetailDao) GetByTransferID(ctx context.Context, transferID string) (*model.TransferDetail, error) {
	// no cache
	if d.cache == nil {
		record := &model.TransferDetail{}
		err := d.db.WithContext(ctx).Where("transfer_id = ?", transferID).First(record).Error
		return record, err
	}

	// get from cache
	record, err := d.cache.Get(ctx, transferID)
	if err == nil {
		return record, nil
	}

	// get from database
	if errors.Is(err, database.ErrCacheNotFound) {
		// for the same transferID, prevent high concurrent simultaneous access to database
		val, err, _ := d.sfg.Do(transferID, func() (interface{}, error) {

			table := &model.TransferDetail{}
			err = d.db.WithContext(ctx).Where("transfer_id = ?", transferID).First(table).Error
			if err != nil {
				// set placeholder cache to prevent cache penetration, default expiration time 10 minutes
				if errors.Is(err, database.ErrRecordNotFound) {
					if err = d.cache.SetPlaceholder(ctx, transferID); err != nil {
						logger.Warn("cache.SetPlaceholder error", logger.Err(err), logger.Any("transferID", transferID))
					}
					return nil, database.ErrRecordNotFound
				}
				return nil, err
			}
			// set cache
			if err = d.cache.Set(ctx, transferID, table, cache.TransferDetailExpireTime); err != nil {
				logger.Warn("cache.Set error", logger.Err(err), logger.Any("transferID", transferID))
			}
			return table, nil
		})
		if err != nil {
			return nil, err
		}
		table, ok := val.(*model.TransferDetail)
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
//	sort: sort fields, default is transferID backwards, you can add - sign before the field to indicate reverse order, no - sign to indicate ascending order, multiple fields separated by comma
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
func (d *transferDetailDao) GetByColumns(ctx context.Context, params *query.Params) ([]*model.TransferDetail, int64, error) {
	if params.Sort == "" {
		params.Sort = "-transfer_id"
	}
	queryStr, args, err := params.ConvertToGormConditions()
	if err != nil {
		return nil, 0, errors.New("query params error: " + err.Error())
	}

	var total int64
	if params.Sort != "ignore count" { // determine if count is required
		err = d.db.WithContext(ctx).Model(&model.TransferDetail{}).Where(queryStr, args...).Count(&total).Error
		if err != nil {
			return nil, 0, err
		}
		if total == 0 {
			return nil, total, nil
		}
	}

	records := []*model.TransferDetail{}
	order, limit, offset := params.ConvertToPage()
	err = d.db.WithContext(ctx).Order(order).Limit(limit).Offset(offset).Where(queryStr, args...).Find(&records).Error
	if err != nil {
		return nil, 0, err
	}

	return records, total, err
}

// CreateByTx create a record in the database using the provided transaction
func (d *transferDetailDao) CreateByTx(ctx context.Context, tx *gorm.DB, table *model.TransferDetail) (string, error) {
	err := tx.WithContext(ctx).Create(table).Error
	return table.TransferID, err
}

// DeleteByTx delete a record by transferID in the database using the provided transaction
func (d *transferDetailDao) DeleteByTx(ctx context.Context, tx *gorm.DB, transferID string) error {
	err := tx.WithContext(ctx).Where("transfer_id = ?", transferID).Delete(&model.TransferDetail{}).Error
	if err != nil {
		return err
	}

	// delete cache
	_ = d.deleteCache(ctx, transferID)

	return nil
}

// UpdateByTx update a record by transferID in the database using the provided transaction
func (d *transferDetailDao) UpdateByTx(ctx context.Context, tx *gorm.DB, table *model.TransferDetail) error {
	err := d.updateDataByTransferID(ctx, tx, table)

	// delete cache
	_ = d.deleteCache(ctx, table.TransferID)

	return err
}
