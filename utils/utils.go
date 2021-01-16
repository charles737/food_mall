package utils

import (
	"crypto/md5"
	"fmt"
	"io"
	"time"
)

const TimeLayout = "2006-01-02 15:04:05"

var (
	Local = time.FixedZone("CST", 8*3600)
)

/*
	Page 分页
	Limit 要查询的记录数
	Page 页码
	offset 偏移量
 */
func Page(Limit, Page int) (limit, offset int) {
	if Limit > 0 {
		limit = Limit
	} else {
		limit = 10
	}
	if Page > 0 {
		offset = (Page -1) * limit
	} else {
		offset = -1
	}
	return limit, offset
}

/*
	Sort 排序
	默认 created_at desc
 */
func Sort(Sort string) (sort string) {
	if Sort != "" {
		sort = Sort
	} else {
		sort = "create_at desc"
	}
	return sort
}

/*
	GetNow 返回当前时间
 */
func GetNow() string {
	now := time.Now().In(Local).Format(TimeLayout)
	return now
}

/*
	时间格式化
 */
func TimeFormat(s string) string {
	result, err := time.ParseInLocation(TimeLayout, s, time.Local)
	if err != nil {
		panic(err)
	}
	return result.In(Local).Format(TimeLayout)
}

/*
	Md5 md5加密
 */
func Md5(str string) string {
	w := md5.New()
	io.WriteString(w, str)
	md5str := fmt.Sprintf("%x", w.Sum(nil))
	return md5str
}