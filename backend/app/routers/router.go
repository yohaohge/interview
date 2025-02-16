package routers

import (
	"www.interview.com/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// 题库相关接口
	r.GET("/questions", controllers.GetQuestionList)

	// 面筋相关接口
	r.GET("/interview-experiences", controllers.GetInterviewExperienceList)

	return r
}
