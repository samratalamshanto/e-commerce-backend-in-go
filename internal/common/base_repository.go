package common

import (
	"gorm.io/gorm"
)

type BaseRepository[T any] struct {
	db *gorm.DB
}

func GetRepository[T any]() *BaseRepository[T] {
	return &BaseRepository[T]{db: DBInstance}
}

func (r *BaseRepository[T]) Save(entity *T) (*T, error) {
	if err := r.db.Save(entity).Error; err != nil {
		return nil, err
	}
	return entity, nil
}

func (r *BaseRepository[T]) FindById(id uint) (*T, error) {
	/*
		var entity *T means: You created a variable of type *T
		Its default value is nil. So entity == nil

		GORM tries to write the DB row into entity.
		But entity is nil,
		so GORM has no memory to put values into
		→ it tries to dereference a nil pointer → panic.

		You need a valid memory location for GORM to fill with DB data.
		var entity T → creates a real T value on the stack (not nil).
		&entity → gives GORM a pointer to write into.
	*/

	var entity T
	if err := r.db.Take(&entity, id).Error; err != nil {
		return nil, err
	}
	return &entity, nil
}

func (r *BaseRepository[T]) FindByIdList(idList []uint) ([]T, error) {
	/*
		Find works without a model instance but Delete doesn’t.
		GORM inspects the type of T ([]T → T).Table name (users),Fields (ID, Name..)
	*/

	var resutlList []T
	if err := r.db.Take(&resutlList, idList).Error; err != nil {
		return nil, err
	}
	return resutlList, nil
}

func (r *BaseRepository[T]) FindAll() ([]T, error) {
	var resultList []T
	if err := r.db.Find(&resultList).Error; err != nil {
		return nil, err
	}
	return resultList, nil
}

func (r *BaseRepository[T]) FindAllPagination(offset, limit int) ([]T, error) {
	var resultList []T
	err := r.db.Limit(limit).Offset(offset).Find(&resultList).Error
	if err != nil {
		return nil, err
	}
	return resultList, nil
}

func (r *BaseRepository[T]) DeleteById(id uint) (bool, error) {
	var entity T //GORM need model reference to determines table name,primary key
	res := r.db.Delete(&entity, id)
	if err := res.Error; err != nil {
		return false, err
	}
	return true, nil
}

func (r *BaseRepository[T]) DeleteIdList(idList []uint) (bool, error) {
	var entity T //GORM need model reference to determines table name,primary key
	if err := r.db.Delete(&entity, idList).Error; err != nil {
		return false, err
	}
	return true, nil
}
