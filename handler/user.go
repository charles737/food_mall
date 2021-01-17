package handler

import (
	"food_mall/enum"
	"food_mall/model"
	"food_mall/query"
	"food_mall/resp"
	"food_mall/service"
	"net/http"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	UserSrv service.UserSrv
}

func (h *UserHandler) GetEntity(result model.User) resp.User {
	return resp.User{
		Id: result.UserId,
		Key: result.UserId,
		UserId: result.UserId,
		NickName: result.NickName,
		Mobile: result.Mobile,
		Address: result.Address,
		IsDeleted: result.IsDeleted,
		IsLocked: result.IsLocked,
	}
}

func (h *UserHandler) UserInfoHandler(c *gin.Context) {
	entity := resp.Entity{
		Code: int(enum.OperateFail),
		Msg: enum.OperateFail.String(),
		Total: 0,
		TotalPage: 1,
		Data: nil,
	}
	userId := c.Param("id")
	if userId == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"entity": entity})
		return
	}
	u := model.User{
		UserId: userId,
	}
	result, err := h.UserSrv.Get(u)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"entity": entity})
		return
	}

	r := h.GetEntity(*result)

	entity = resp.Entity{
		Code: http.StatusOK,
		Msg: "OK",
		Total: 0,
		TotalPage: 0,
		Data: r,
	}
	c.JSON(http.StatusOK, gin.H{"entity": entity})
}

func (h *UserHandler) UserListHandler(c *gin.Context) {
	var q query.ListQuery

	entity := resp.Entity{
		Code: int(enum.OperateFail),
		Msg: enum.OperateFail.String(),
		Total: 0,
		TotalPage: 1,
		Data: nil,
	}

	err := c.ShouldBindQuery(&q)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"entity": entity})
		return
	}

	list, err := h.UserSrv.List(&q)
	total, err := h.UserSrv.GetTotal(&q)
	if err != nil {
		panic(err)
	}

	if q.PageSize == 0 {
		q.PageSize = 5
	}
	ret := int(total % q.PageSize)
	ret2 := int(total / q.PageSize)
	totalPage := 0
	if ret == 0 {
		totalPage = ret2
	} else {
		totalPage = ret2 + 1
	}

	var newList []*resp.User
	for _, item := range list {
		r := h.GetEntity(*item)
		newList = append(newList, &r)
	}

	entity = resp.Entity{
		Code: http.StatusOK,
		Msg: "OK",
		Total: total,
		TotalPage: totalPage,
		Data: newList,
	}
	c.JSON(http.StatusOK, gin.H{"entity": entity})
}

func (h *UserHandler) AddUserHandler(c *gin.Context)  {
	entity := resp.Entity{
		Code: int(enum.OperateFail),
		Msg: enum.OperateFail.String(),
		Total: 0,
		Data: nil,
	}
	u := model.User{}
	err := c.ShouldBindJSON(&u)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"entity": entity})
		return
	}

	r, err := h.UserSrv.Add(u)
	if err != nil {
		entity.Msg = err.Error()
		return
	}
	if r.UserId == "" {
		c.JSON(http.StatusOK, gin.H{"entity": entity})
		return
	}

	entity.Code = int(enum.OperateOk)
	entity.Msg = enum.OperateOk.String()
	c.JSON(http.StatusOK, gin.H{"entity": entity})
}

func (h *UserHandler) EditUserHandler(c *gin.Context) {
	
}