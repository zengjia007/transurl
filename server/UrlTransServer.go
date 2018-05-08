package main

import (
	"github.com/gin-gonic/gin"
	"transurl/api"
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin/binding"
)

var msg string = "转换成功！"

type LongUrl struct {
	LongUrl string `form:"long_url" binding:"required"`
}

type ShortUrl struct {
	ShortUrl string `form:"short_url" binding:"required"`
}

func main()  {
	router := gin.Default()
	router.POST("/longToShortUrl", func(context *gin.Context) {
		//originUrl := context.PostForm("long_url")
		var longUrl LongUrl
		err := context.MustBindWith(&longUrl, binding.FormMultipart)
		if err == nil {
			// 调用转换为短URL的逻辑
			shortUrl, e := api.LongToShort(longUrl.LongUrl)
			if e != nil {
				fmt.Println("long url trans to short url fail, fail info: ", err.Error())
				msg = "转换失败！"
			}
			// for test
			context.JSON(200, gin.H{
				"short_url": shortUrl,
				"msg": msg,
			})
		} else {
			context.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}
	})

	router.POST("/shortToLongUrl", func(context *gin.Context) {
		//shortUrl := context.PostForm("short_url")
		var shortUrl ShortUrl
		err := context.MustBindWith(&shortUrl, binding.FormMultipart)
		if err == nil {
			// 调用转换为长URL的逻辑
			originUrl, e := api.ShortToLongUrl(shortUrl.ShortUrl)

			if e != nil {
				fmt.Println("short url trans to long url fail, fail info: ", err.Error())
				msg = "转换失败！"
			}
			context.JSON(200, gin.H{
				"short_url": originUrl,
				"meg": msg,
			})
		} else {
			context.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}
	})
	// 监听本机的8088端口
	router.Run(":8088")
}
