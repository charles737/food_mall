package service

import (
	"errors"
	"food_mall/model"
	"food_mall/query"
	"food_mall/repository"
)

type OrderService struct {
	Repo repository.OrderRepoInterface
}

type OrderSrv interface {
	List(req *query.ListQuery) (Orders []*model.Order, err error)
	GetTotal(req *query.ListQuery) (total int, err error)
	Get(Order model.Order) (*model.Order, error)
	Exist(Order model.Order) *model.Order
	ExistByOrderID(id string) *model.Order
	Add(Order model.Order) (*model.Order, error)
	Edit(Order model.Order) (bool, error)
	Delete(Order model.Order) (bool, error)
}

func (srv *OrderService) List(req *query.ListQuery) (orders []*model.Order, err error) {
	return srv.Repo.List(req)
}

func (srv *OrderService) GetTotal(req *query.ListQuery) (total int, err error) {
	return srv.GetTotal(req)
}

func (srv *OrderService) Get(order model.Order) (*model.Order, error) {
	return srv.Repo.Get(order)
}

func (srv *OrderService) Exist(order model.Order) *model.Order {
	return srv.Repo.Exist(order)
}

func (srv *OrderService) ExistByOrderID(id string) *model.Order {
	return srv.Repo.ExistByOrderID(id)
}

func (srv *OrderService) Add(order model.Order) (*model.Order, error) {
	return srv.Repo.Add(order)
}

func (srv *OrderService) Edit(order model.Order) (bool, error) {
	o := srv.ExistByOrderID(order.OrderId)
	if o == nil || o.Mobile == "" {
		return false, errors.New("订单号不存在")
	}
	return srv.Repo.Edit(order)
}

func (srv *OrderService) Delete(order model.Order) (bool, error) {
	order.IsDeleted = !order.IsDeleted
	return srv.Repo.Delete(order)
}