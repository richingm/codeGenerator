package infrastructure

import (
	"context"
	"errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/richingm/knowledge/internal/biz"
	"gorm.io/gorm"
)

type {{.RepoNameToLower}} struct {
	data *Data
	log  *log.Helper
}

// NewKnowledgeRepo .
func NewKnowledgeRepo(data *Data, logger log.Logger) biz.{{.RepoIfName}} {
	return &knowledgeRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *{{.RepoNameToLower}}) Create(ctx context.Context, po biz.KnowledgePo) (*biz.KnowledgePo, error) {
	err := r.data.DB(ctx).Create(&po).Error
	if err != nil {
		return nil, err
	}
	return &po, nil
}

func (r *{{.RepoNameToLower}}) Update(ctx context.Context, po biz.KnowledgePo) (*biz.KnowledgePo, error) {
	err := r.data.DB(ctx).Save(&po).Error
	if err != nil {
		return nil, err
	}
	return &po, nil
}

func (r *{{.RepoNameToLower}}) Delete(ctx context.Context, id int64) error {
	err := r.data.DB(ctx).Delete(&biz.KnowledgePo{}, id).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *{{.RepoNameToLower}}) Find(ctx context.Context, id int64) (*biz.KnowledgePo, error) {
	var res biz.KnowledgePo
	err := r.data.DB(ctx).Model(&biz.KnowledgePo{}).Where("id = ?", id).First(&res).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &res, nil
}

func (r *{{.RepoNameToLower}}) ScopeKeyWord(keyWord string) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if len(keyWord) == 0 {
			return db
		}
		return db.Where("name like ?", "%"+keyWord+"%")
	}
}

func (r *{{.RepoNameToLower}}) ScopeId(id int64) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if id <= 0 {

			return db
		}
		return db.Where("id = ?", id)
	}
}

func (r *{{.RepoNameToLower}}) ScopePid(pid int64) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if pid <= 0 {

			return db
		}
		return db.Where("pid = ?", pid)
	}
}

func (r *{{.RepoNameToLower}}) ScopeImportLevel(importLevel string) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if len(importLevel) == 0 {
			return db
		}
		return db.Where("import_level = ?", importLevel)
	}
}

func (r *{{.RepoNameToLower}}) Count(ctx context.Context, wheres ...func(*gorm.DB) *gorm.DB) (int64, error) {
	var total int64
	err := r.data.DB(ctx).Model(&biz.KnowledgePo{}).Scopes(wheres...).Count(&total).Error
	if err != nil {
		return 0, err
	}
	return total, nil
}

func (r *{{.RepoNameToLower}}) List(ctx context.Context, order string, wheres ...func(*gorm.DB) *gorm.DB) ([]biz.KnowledgePo, error) {
	var res []biz.KnowledgePo
	db := r.data.DB(ctx).Model(&biz.KnowledgePo{}).Scopes(wheres...)
	if len(order) > 0 {
		db = db.Order(order)
	}
	err := db.Find(&res).Error
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (r *{{.RepoNameToLower}}) Page(ctx context.Context, page, pageSize int64, order string, wheres ...func(*gorm.DB) *gorm.DB) (int64, []biz.KnowledgePo, error) {
	// count
	count, err := r.Count(ctx, wheres...)
	if err != nil {
		return 0, nil, err
	}

	// data
	var list []biz.KnowledgePo
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 20
	}
	offset := (page - 1) * pageSize
	db := r.data.DB(ctx).Model(&biz.KnowledgePo{}).Scopes(wheres...).Offset(int(offset)).Limit(int(pageSize))
	if len(order) > 0 {
		db = db.Order(order)
	}
	err = db.Find(&list).Error
	if err != nil {
		return 0, nil, nil
	}
	return count, list, nil
}
