package biz

import (
	"context"
	"gorm.io/gorm"
)

// KnowledgeRepoIf is a repo.
type KnowledgeRepoIf interface {
	Create(ctx context.Context, po KnowledgePo) (*KnowledgePo, error)
	Delete(ctx context.Context, id int64) error
	Update(ctx context.Context, po KnowledgePo) (*KnowledgePo, error)
	Find(ctx context.Context, id int64) (*KnowledgePo, error)
	Count(ctx context.Context, wheres ...func(*gorm.DB) *gorm.DB) (int64, error)
	List(ctx context.Context, order string, wheres ...func(*gorm.DB) *gorm.DB) ([]KnowledgePo, error)
	Page(ctx context.Context, page, pageSize int64, order string, wheres ...func(*gorm.DB) *gorm.DB) (int64, []KnowledgePo, error)

	ScopeId(id int64) func(*gorm.DB) *gorm.DB
}
