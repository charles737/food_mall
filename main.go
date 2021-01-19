package main

import (
	"food_mall/handler"
	"net/http"
	"github.com/spf13/viper"
	"github.com/gin-gonic/gin"
)

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type, AccessToken, X-CSRF-Token, Authorization, Token")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		// 处理请求
		c.Next()
	}
}

func main() {
	r := gin.Default()
	r.Use(Cors())
	gin.SetMode(viper.GetString("mode"))

	user := r.Group("/api/user")
	{
		user.GET("/list", UserHandler.UserListHandler)
		user.GET("/info/:id", UserHandler.UserInfoHandler)
		user.POST("/add", UserHandler.AddUserHandler)
		user.POST("/edit", UserHandler.EditUserHandler)
		user.POST("/delete/:id", UserHandler.DeleteUserHandler)
	}

	product := r.Group("/api/product")
	{
		product.GET("/list", ProductHandler.ProductListHandler)
		product.GET("/info/:id", ProductHandler.ProductInfoHandler)
		product.POST("/add", ProductHandler.AddProductHandler)
		product.POST("/edit", ProductHandler.EditProductHandler)
		product.POST("/delete", ProductHandler.DeleteProductHandler)
	}

	port := viper.GetString("port")

	r.Run(port)
}