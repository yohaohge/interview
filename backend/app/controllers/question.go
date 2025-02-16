package controllers

import (
	"net/http"
	"strconv"

	"www.interview.com/models"
	"www.interview.com/utils"

	"github.com/gin-gonic/gin"
)

// GetQuestionList 获取题库列表，支持分页和关键词搜索
func GetQuestionList(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))
	keyword := c.Query("keyword")

	var questions []models.Question
	offset := (page - 1) * pageSize

	db := utils.DB

	if keyword != "" {
		db = db.Where("content LIKE ? OR tags LIKE ? OR related_company LIKE ?", "%"+keyword+"%", "%"+keyword+"%", "%"+keyword+"%")
	}

	result := db.Offset(offset).Limit(pageSize).Find(&questions)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, questions)
}
