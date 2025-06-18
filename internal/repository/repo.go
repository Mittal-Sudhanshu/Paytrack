package repository

import (
	"context"

	"gorm.io/gorm"
)

type Repository[T any] struct {
	DB *gorm.DB
}

func NewRepository[T any](db *gorm.DB) *Repository[T] {
	return &Repository[T]{DB: db}
}
func (r *Repository[T]) Create(ctx context.Context, entity *T) (*T, error) {
	err := r.DB.WithContext(ctx).Create(entity).Error
	return entity, err
}

func (r *Repository[T]) GetByID(ctx context.Context, id string) (*T, error) {
	var entity T
	err := r.DB.WithContext(ctx).First(&entity, id).Error
	return &entity, err
}

func (r *Repository[T]) Update(ctx context.Context, id any, updates map[string]interface{}) error {
	var entity T
	return r.DB.WithContext(ctx).Model(&entity).Where("id = ?", id).Updates(updates).Error
}

func (r *Repository[T]) Delete(ctx context.Context, id any) error {
	var entity T
	return r.DB.WithContext(ctx).Delete(&entity, id).Error
}

func (r *Repository[T]) List(ctx context.Context) ([]T, error) {
	var entities []T
	err := r.DB.WithContext(ctx).Find(&entities).Error
	return entities, err
}
func (r *Repository[T]) Query(ctx context.Context, params map[string]interface{}) ([]T, error) {
	var entities []T
	query := r.DB.WithContext(ctx)
	for key, value := range params {
		query = query.Where(key+" = ?", value)
	}
	err := query.Find(&entities).Error
	return entities, err
}
func (r *Repository[T]) WithTransaction(ctx context.Context, fn func(txRepo *Repository[T]) error) error {
	return r.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		txRepo := &Repository[T]{DB: tx}
		return fn(txRepo)
	})
}
