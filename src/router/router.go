package router

import (
	config2 "config"
	product2 "controller/product"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	r := gin.Default()

	//用户
	user := r.Group("/users")
	{
		//上传图片
		user.POST("/product/uploadImage", product2.UploadImg)
	}

	_ = r.Run(":" + config2.GetKV("dev", "port"))
}
