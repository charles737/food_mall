package main

import (
	"fmt"
	"food_mall/config"
	"food_mall/handler"
	"food_mall/model"
	"food_mall/repository"
	"food_mall/service"
	"log"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

var (
	DB *gorm.DB
	UserHandler handler.UserHandler
	ProductHandler handler.ProductHandler
)

func initViper() {
	if err := config.Init(""); err != nil {
		panic(err)
	}
}

func initDB() {
	fmt.Println("Database init start...")

	conf := &model.DBConf{
		Host: viper.GetString("database.host"),
		User: viper.GetString("database.user"),
		Password: viper.GetString("database.password"),
		DbName: viper.GetString("database.name"),
	}

	config := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=%t&loc=%s",
		conf.User,
		conf.Password,
		conf.Host,
		conf.DbName,
		true,
		"Local",
		)

	var err error
	// 这里不要使用 := ，会和 config 重复声明变量
	DB, err = gorm.Open("mysql", config)
	if err != nil {
		log.Fatalf("connect error: %v\n", err)
	}
	// 数据库表名单数
	DB.SingularTable(true)
	fmt.Println("Database init over...")
}

func initHandler() {
	UserHandler = handler.UserHandler{
		UserSrv: &service.UserService{
			Repo: &repository.UserRepository{
				DB: DB,
			},
		}}

	ProductHandler = handler.ProductHandler{
		ProductSrv: &service.ProductService{
			Repo: &repository.ProductRepository{
				DB: DB,
			},
		}}
}

func init() {
	initViper()
	initDB()
	initHandler()
}