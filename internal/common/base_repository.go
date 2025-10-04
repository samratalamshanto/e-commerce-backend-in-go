package common

import (
	"gorm.io/gorm"
)

// Define interface
type BaseRepositoryInterface[T any] interface {
	Save(entity *T) (*T, error)
	FindById(id uint) (*T, error)
	FindByIdList(idList []uint) ([]T, error)
	FindAll() ([]T, error)
	FindAllPagination(offset, limit int) ([]T, error)
	DeleteById(id uint) (bool, error)
	DeleteIdList(idList []uint) (bool, error)
}

/*
var _ Interface = &Struct{} -->compile-time assertion that *Struct implements the interface Interface.
doesn’t actually create a variable at runtime — the blank identifier _ discards it.
If Struct does not implement all methods of Interface, the compiler will throw an error.

var x struct{}    // declare variable of type empty struct, default values assigned.
y := struct{}{}   // create a new empty struct instance, zero bytes.const

don’t care about the actual value — we only care that *Struct can be assigned to Interface.
The variable name _ discards it (so no runtime memory is really kept).
*/

var _ BaseRepositoryInterface[any] = &BaseRepository[any]{}

type BaseRepository[T any] struct {
	DB *gorm.DB
}

func GetBaseRepository[T any](db *gorm.DB) *BaseRepository[T] {
	return &BaseRepository[T]{DB: db}
}

func (r *BaseRepository[T]) Save(entity *T) (*T, error) {
	if err := r.DB.Save(entity).Error; err != nil {
		return nil, err
	}
	return entity, nil
}

func (r *BaseRepository[T]) FindById(id uint) (*T, error) {
	/*
		wrong: var entity *T means: You created a variable of type *T
		Its default value is nil. So entity == nil

		GORM tries to write the DB row into entity.
		But entity is nil, so GORM has no memory to put values into
		→ it tries to dereference a nil pointer → panic.

		Right: You need a valid memory location for GORM to fill with DB data.
		var entity T → creates a real T value on the stack (not nil).
		&entity → gives GORM a pointer to write into.
	*/

	var entity T
	if err := r.DB.Take(&entity, id).Error; err != nil {
		return nil, err
	}
	return &entity, nil
}

func (r *BaseRepository[T]) FindByIdList(idList []uint) ([]T, error) {
	/*
		Find works without a model instance but Delete doesn’t.
		GORM inspects the type of T ([]T → T).Table name (users),Fields (ID, Name..)
	*/

	var resultList []T
	if err := r.DB.Find(&resultList, idList).Error; err != nil {
		return nil, err
	}
	return resultList, nil
}

func (r *BaseRepository[T]) FindAll() ([]T, error) {
	var resultList []T
	if err := r.DB.Find(&resultList).Error; err != nil {
		return nil, err
	}
	return resultList, nil
}

func (r *BaseRepository[T]) FindAllPagination(offset, limit int) ([]T, error) {
	var resultList []T
	err := r.DB.Limit(limit).Offset(offset).Find(&resultList).Error
	if err != nil {
		return nil, err
	}
	return resultList, nil
}

func (r *BaseRepository[T]) DeleteById(id uint) (bool, error) {
	var entity T //GORM need model reference to determines table name,primary key
	res := r.DB.Delete(&entity, id)
	if err := res.Error; err != nil {
		return false, err
	}
	return true, nil
}

func (r *BaseRepository[T]) DeleteIdList(idList []uint) (bool, error) {
	var entity T //GORM need model reference to determines table name,primary key
	if err := r.DB.Delete(&entity, idList).Error; err != nil {
		return false, err
	}
	return true, nil
}
