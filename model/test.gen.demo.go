package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _DemoMgr struct {
	*_BaseMgr
}

// DemoMgr open func
func DemoMgr(db *gorm.DB) *_DemoMgr {
	if db == nil {
		panic(fmt.Errorf("DemoMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_DemoMgr{_BaseMgr: &_BaseMgr{DB: db.Table("demo"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_DemoMgr) GetTableName() string {
	return "demo"
}

// Get 获取
func (obj *_DemoMgr) Get() (result Demo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Demo{}).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_DemoMgr) Gets() (results []*Demo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Demo{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_DemoMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(Demo{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 Primary key
func (obj *_DemoMgr) WithID(id int64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithCreatedAt created_at获取 created time
func (obj *_DemoMgr) WithCreatedAt(createdAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["created_at"] = createdAt })
}

// WithUpdatedAt updated_at获取 updated at
func (obj *_DemoMgr) WithUpdatedAt(updatedAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["updated_at"] = updatedAt })
}

// WithDeletedAt deleted_at获取 deleted time
func (obj *_DemoMgr) WithDeletedAt(deletedAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["deleted_at"] = deletedAt })
}

// WithName name获取
func (obj *_DemoMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// GetByOption 功能选项模式获取
func (obj *_DemoMgr) GetByOption(opts ...Option) (result Demo, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Demo{}).Where(options.query).Find(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_DemoMgr) GetByOptions(opts ...Option) (results []*Demo, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Demo{}).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromID 通过id获取内容 Primary key
func (obj *_DemoMgr) GetFromID(id int64) (result Demo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Demo{}).Where("id = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量查找 Primary key
func (obj *_DemoMgr) GetBatchFromID(ids []int64) (results []*Demo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Demo{}).Where("id IN (?)", ids).Find(&results).Error

	return
}

// GetFromCreatedAt 通过created_at获取内容 created time
func (obj *_DemoMgr) GetFromCreatedAt(createdAt time.Time) (result Demo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Demo{}).Where("created_at = ?", createdAt).Find(&result).Error

	return
}

// GetBatchFromCreatedAt 批量查找 created time
func (obj *_DemoMgr) GetBatchFromCreatedAt(createdAts []time.Time) (results []*Demo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Demo{}).Where("created_at IN (?)", createdAts).Find(&results).Error

	return
}

// GetFromUpdatedAt 通过updated_at获取内容 updated at
func (obj *_DemoMgr) GetFromUpdatedAt(updatedAt time.Time) (result Demo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Demo{}).Where("updated_at = ?", updatedAt).Find(&result).Error

	return
}

// GetBatchFromUpdatedAt 批量查找 updated at
func (obj *_DemoMgr) GetBatchFromUpdatedAt(updatedAts []time.Time) (results []*Demo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Demo{}).Where("updated_at IN (?)", updatedAts).Find(&results).Error

	return
}

// GetFromDeletedAt 通过deleted_at获取内容 deleted time
func (obj *_DemoMgr) GetFromDeletedAt(deletedAt time.Time) (result Demo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Demo{}).Where("deleted_at = ?", deletedAt).Find(&result).Error

	return
}

// GetBatchFromDeletedAt 批量查找 deleted time
func (obj *_DemoMgr) GetBatchFromDeletedAt(deletedAts []time.Time) (results []*Demo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Demo{}).Where("deleted_at IN (?)", deletedAts).Find(&results).Error

	return
}

// GetFromName 通过name获取内容
func (obj *_DemoMgr) GetFromName(name string) (results []*Demo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Demo{}).Where("`name` = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量查找
func (obj *_DemoMgr) GetBatchFromName(names []string) (results []*Demo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Demo{}).Where("`name` IN (?)", names).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_DemoMgr) FetchByPrimaryKey(id int64) (result Demo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Demo{}).Where("`id` = ?", id).Find(&result).Error

	return
}
