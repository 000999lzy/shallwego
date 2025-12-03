package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	Username string `form:"username" binding:"required" validate:"min=2" json:"user"`
	Password string `form:"password" binding:"required" json:"password"`
	Gender   string `form:"gender" json:"gen"`
}

func main() {
	//自定义日志格式
	/*router := gin.New()

	router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("%s|%s|%s|%s|%d|%d|%s|%s\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))

	router.Use(gin.Recovery())

	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	router.Run(":8888")*/

	gin.ForceConsoleColor()

	router := gin.Default()

	//get请求参数获取
	router.GET("/user/:bid", func(c *gin.Context) {
		uid := c.Query("uid")
		username := c.Query("name")
		page := c.DefaultQuery("page", "10")

		bid := c.Param("bid")

		c.JSON(http.StatusOK, gin.H{
			"uid":      uid,
			"username": username,
			"page":     page,
			"bid":      bid,
		})
	})

	//get请求参数绑定结构体
	router.GET("/userinfo", func(c *gin.Context) {
		fmt.Println(c.GetHeader("Content-Type"))

		var userinfo User
		if err := c.ShouldBind(&userinfo); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, userinfo)
	})

	//post请求参数获取
	router.POST("/adduser", func(c *gin.Context) {
		username := c.PostForm("username")
		gender := c.DefaultPostForm("gender", "male")
		hobby := c.PostFormArray("hobby")

		c.JSON(http.StatusOK, gin.H{
			"username": username,
			"gender":   gender,
			"hobby":    hobby,
		})
	})

	//post请求参数绑定结构体
	router.POST("/login", func(c *gin.Context) {
		fmt.Println(c.GetHeader("Content-Type"))

		var userinfo User
		if err := c.ShouldBind(&userinfo); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, userinfo)
	})

	//使用模板渲染
	router.LoadHTMLGlob("../templates/**/*")
	router.GET("/get/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "get/index.tmpl", gin.H{
			"title": "get",
		})
	})

	router.GET("/post/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "post/index.tmpl", gin.H{
			"title": "post",
		})
	})

	router.Run(":8888")
}
