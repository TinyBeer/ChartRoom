package router

import (
	"ChatRoom/Web/handlers"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var (
	key_cookie string = "gin_cookie"
	val_cookie string = "test"
)

// 函数方式返回中间件   也可使用中间件函数
func AuthMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取客户端cookie并校验
		if cookie, err := c.Cookie(key_cookie); err == nil {
			if cookie == val_cookie {
				c.Next()
				return
			}
			fmt.Println(cookie)
		}

		// 返回错误代码
		c.JSON(304, gin.H{
			"error": "err",
		})

		// 禁止后续访问
		c.Abort()
		return

	}
}

func SetupRouter() (r *gin.Engine) {
	r = gin.Default()

	r.LoadHTMLGlob("templates/*")

	r.Static("/xxx", "statics")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", nil)
	})

	r.GET("/hall", AuthMiddleWare(), func(c *gin.Context) {
		c.HTML(http.StatusOK, "hall.html", nil)
	})

	r.GET("/cookie", func(c *gin.Context) {
		// 获取客户端携带的cookie
		// c.Request.Cookie()
		cookie, err := c.Cookie("key_cookie")
		// 设置domain  注意区分本地测试和外网访问
		domain := "localhost"
		if c.ClientIP() == "::1" {
			domain = "http://3843359ku3.wicp.vip"
		}
		fmt.Println()
		if err != nil {
			cookie = "NotSet"
			// 设置cookie
			c.SetCookie("key_cookie", "value_cookie", 120, "/", domain, false, true)
		}
		fmt.Println("cookie=", cookie)
		c.JSON(http.StatusOK, gin.H{
			"cookie": cookie,
		})
	})

	r.GET("/content", AuthMiddleWare(), func(c *gin.Context) {
		userID := 100
		handlers.GetContentHandler(c, userID)
	})

	r.POST("/content", AuthMiddleWare(), func(c *gin.Context) {
		cookie, err := c.Cookie("gin_cookie")
		fmt.Println(cookie)
		if err != nil {
			c.Abort()
			return
		}
		userID, _ := strconv.Atoi(cookie)
		fmt.Println(userID)
		handlers.PostContentHandler(c, userID)
	})

	r.POST("/login", func(c *gin.Context) {
		userID, err := handlers.LoginHandler(c)
		if err != nil {
			return
		}
		fmt.Println(userID)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"res": "fail",
				"err": "",
			})
		} else {

			// 设置domain  注意区分本地测试和外网访问
			domain := "localhost"
			if c.ClientIP() != "::1" {
				domain = "http://3843359ku3.wicp.vip"
			}

			c.SetCookie("gin_cookie", val_cookie, 100, "/", domain, false, true)
			c.JSON(http.StatusOK, gin.H{
				"res": "ok",
			})
		}
	})

	r.POST("/register", func(c *gin.Context) {
		err := handlers.RegisterHandler(c)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"res": "fail",
				"err": err.Error(),
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"res": "ok",
			})
		}
	})

	return
}
