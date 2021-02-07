package repository

import (
	"fmt"
	"food_mall/model"
	"food_mall/query"
	"food_mall/utils"
	"github.com/jinzhu/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

type UserRepoInterface interface {
	List(req *query.ListQuery) (users []*model.User, err error)
	GetTotal(req *query.ListQuery) (total int, err error)
	Get(user model.User) (*model.User, error)
	Exist(user model.User) *model.User
	ExistByUserID(id string) *model.User
	ExistByMobile(mobile string) *model.User
	Add(user model.User) (*model.User, error)
	Edit(user model.User) (bool, error)
	Delete(user model.User) (bool, error)
}

func (repo *UserRepository) List(req *query.ListQuery) (users []*model.User, err error) {
	fmt.Println(req)

	// 分页
	limit, offset := utils.Page(req.PageSize, req.Page)

	if err := repo.DB.Order("user_id desc").Limit(limit).Offset(offset).Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (repo *UserRepository) GetTotal(req *query.ListQuery) (total int, err error) {
	var users []model.User

	if err := repo.DB.Find(&users).Count(&total).Error; err != nil {
		return total, err
	}
	return total, nil
}

func (repo *UserRepository) Get(user model.User) (*model.User, error) {
	if err := repo.DB.Where(&user).Find(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (repo *UserRepository) Exist(user model.User) *model.User {
	var count int

	repo.DB.Find(&user).Where("nick_name=?", user.NickName)
	if count > 0 {
		return &user
	}
	return nil
}

func (repo *UserRepository) ExistByUserID(id string) *model.User {
	var user model.User

	repo.DB.Where("user_id=?", id).First(&user)
	return &user
}

func (repo *UserRepository) ExistByMobile(mobile string) *model.User {
	var user model.User
	var count int

	repo.DB.Where("mobile=?", mobile).Find(&user)
	if count > 0 {
		return &user
	}
	return nil
}

func (repo *UserRepository) Add(user model.User) (*model.User, error) {
	if exist := repo.Exist(user); exist != nil {
		return nil, fmt.Errorf("该用户已存在")
	}
	if err := repo.DB.Create(&user).Error; err != nil {
		return nil, fmt.Errorf("用户注册失败")
	}
	return &user, nil
}

func (repo *UserRepository) Edit(user model.User) (bool, error) {
	err := repo.DB.Model(&user).Where("user_id=?", user.UserId).Updates(map[string]interface{}{
		"nick_name": user.NickName,
		"mobile": user.Mobile,
		"address": user.Address,
	}).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

func (repo *UserRepository) Delete(user model.User) (bool, error) {
	err := repo.DB.Model(&user).Where("user_id=?", user.UserId).Update("is_deleted", user.IsDeleted).Error
	if err != nil {
		return false, err
	}
	return true, nil
}
