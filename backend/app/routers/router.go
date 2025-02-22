// app/routers/router.go
package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"www.interview.com/controllers"
	"www.interview.com/middleware"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// 题库相关接口
	r.GET("/questions", controllers.GetQuestionList)
	// 新增：获取单个问题详情的接口
	r.GET("/questions/:id", controllers.GetQuestionDetail)

	// 面筋相关接口
	r.GET("/interview-experiences", controllers.GetInterviewExperienceList)

	// 用户注册和登录接口
	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)

	// 需要身份验证的接口示例
	authenticated := r.Group("/")
	authenticated.Use(middleware.AuthMiddleware())
	authenticated.GET("/protected", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "This is a protected route"})
	})

	return r
}
