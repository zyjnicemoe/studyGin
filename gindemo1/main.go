package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", nil)
	})
	r.POST("/login", func(c *gin.Context) {
		username, _ := c.GetPostForm("username")
		password, _ := c.GetPostForm("password")
		c.HTML(http.StatusOK, "success.html", gin.H{
			"username": username,
			"password": password,
		})
	})
	r.GET("/param/:name/:age", func(c *gin.Context) {
		name := c.Param("name")
		age := c.Param("age")
		c.HTML(http.StatusOK, "param.html", gin.H{
			"name": name,
			"age":  age,
		})
	})
	r.GET("/blog/:year/:month", func(c *gin.Context) {
		year := c.Param("year")
		month := c.Param("month")
		c.HTML(http.StatusOK, "blog.html", gin.H{
			"year":  year,
			"month": month,
		})
	})

	type UserInfo struct {
		Username string `form:"username" json:"username"`
		Password string `form:"password" json:"password"`
	}
	r.GET("/user", func(c *gin.Context) {
		//username := c.Query("username")
		//password := c.Query("password")
		//user := &UserInfo{
		//	Username: username,
		//	Password: password,
		//}
		user := &UserInfo{}
		//将传入的参数和user的属性绑定起来
		err := c.ShouldBind(&user)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"msg": "参数错误",
			})
		} else {
			c.HTML(http.StatusOK, "user.html", gin.H{
				"user": user,
			})
		}
	})

	r.POST("/form", func(c *gin.Context) {
		user := &UserInfo{}
		//将传入的参数和user的属性绑定起来
		err := c.ShouldBind(&user)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"msg": "参数错误",
			})
		} else {
			c.HTML(http.StatusOK, "user.html", gin.H{
				"user": user,
			})
		}
	})

	r.GET("/upload", func(c *gin.Context) {
		c.HTML(http.StatusOK, "upload.html", nil)
	})
	r.MaxMultipartMemory = 8 << 20 //8M内存限制
	r.POST("/upload", func(c *gin.Context) {
		//请求中读取文件
		file, err := c.FormFile("file")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"msg": err.Error(),
			})
		} else {
			//将文件保存到本地
			c.SaveUploadedFile(file, "upload/"+file.Filename)
			c.JSON(http.StatusOK, gin.H{
				"status": "OK",
				"url":    "/upload",
			})
		}
	})
	r.POST("/uploads", func(c *gin.Context) {
		//请求中读取文件
		form, _ := c.MultipartForm()
		files := form.File["file"]
		for i, file := range files {
			log.Println(i, file.Filename)
			c.SaveUploadedFile(file, "upload/"+file.Filename)
		}
		c.JSON(http.StatusOK, gin.H{
			"status": "OK",
			"msg":    "多文件上传成功",
		})
	})

	//http重定向
	r.GET("/test", func(c *gin.Context) {
		//跳转
		c.Redirect(http.StatusMovedPermanently, "https://www.bilibili.com")
	})
	//
	r.GET("/a", func(c *gin.Context) {
		//转跳到/b 对于的路由处理函数
		//将请求URI修改
		c.Request.URL.Path = "/b"
		//继续后续处理
		r.HandleContext(c)
	})
	r.GET("/b", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "OK b",
		})
	})

	//路由与路由组
	r.GET("/index", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": "GET",
		})
	})
	r.POST("/index", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": "POST",
		})
	})
	r.PUT("/index", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": "PUT",
		})
	})
	r.DELETE("/index", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": "DELETE",
		})
	})
	//r.HEAD("/index", func(c *gin.Context) {
	//	c.JSON(http.StatusOK, gin.H{
	//		"method": "HEAD",
	//	})
	//})
	r.Any("/any", func(c *gin.Context) {
		switch c.Request.Method {
		case http.MethodGet:
			c.JSON(http.StatusOK, gin.H{
				"method": "GET",
			})
		case http.MethodPost:
			c.JSON(http.StatusOK, gin.H{
				"method": "POST",
			})
		case http.MethodPut:
			c.JSON(http.StatusOK, gin.H{
				"method": "PUT",
			})
		case http.MethodDelete:
			c.JSON(http.StatusOK, gin.H{
				"method": "DELETE",
			})
		case http.MethodHead:
			c.JSON(http.StatusOK, gin.H{
				"method": "HEAD",
			})
		case http.MethodPatch:
			c.JSON(http.StatusOK, gin.H{
				"method": "Patch",
			})
		}
	})
	//访问路由中不存在的转跳到404.html
	r.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusNotFound, "404.html", gin.H{})
	})

	//视频的首页和详情页
	//r.GET("/video/index", func(c *gin.Context) {
	//	c.JSON(http.StatusOK, gin.H{
	//		"msg": "/video/index",
	//	})
	//})
	//r.GET("/video/main", func(c *gin.Context) {
	//	c.JSON(http.StatusOK, gin.H{
	//		"msg": "/video/main",
	//	})
	//})
	//商城的首页和详情页
	//r.GET("/shop/index", func(c *gin.Context) {
	//	c.JSON(http.StatusOK, gin.H{
	//		"msg": "/shop/index",
	//	})
	//})
	//路由组
	//将共用前缀提取出来
	vedioGroup := r.Group("/video")
	{
		vedioGroup.GET("/index", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"msg": "/video/index",
			})
		})
		vedioGroup.GET("/main", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"msg": "/video/main",
			})
		})
	}

	shopGroup := r.Group("/shop")
	{
		shopGroup.GET("/index", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"msg": "/shop/index",
			})
		})
		shopGroup.GET("/main", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"msg": "/shop/main",
			})
		})
		xx := shopGroup.Group("/xx")
		xx.GET("/oo", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"msg": "/shop/xx/oo",
			})
		})
	}
	r.Run(":9029")
}
