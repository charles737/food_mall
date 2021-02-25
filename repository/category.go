package repository

import (
	"food_mall/model"
	"food_mall/query"
	"github.com/jinzhu/gorm"
)

type CategoryRepository struct {
	DB *gorm.DB
}

type CategoryRepoInterface interface {
	List(req *query.ListQuery) (Categories []*model.Category, err error)
	GetTotal(req *query.ListQuery) (total int, err error)
	Get(category model.Category) ([]*model.CategoryResult, error)
	Exist(category model.Category) *model.Category
	ExistByCategoryID(id string) *model.Category
	Add(category model.Category) (*model.Category, error)
	Edit(category model.Category) (bool, error)
	Delete(category model.Category) (bool, error)
}