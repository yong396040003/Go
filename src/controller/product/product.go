package product

import (
	"config"
	"fmt"
	"github.com/gin-gonic/gin"
)

func UploadImg(c *gin.Context) {
	//取文件 取得文件以流的方式存在
	file, err := c.FormFile("file")
	//取图片id
	imgId := c.PostForm("img_id")

	if err != nil {
		fmt.Print(err)
	}

	imageName := imgId + ".png"
	if err := c.SaveUploadedFile(file,
		config.GetKV("file", "LocalPath")+imageName); err != nil {
		c.String(400, "Fail")
		return
	}
	//上传成功
	c.String(200, "Success")
}
