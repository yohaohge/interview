// app/controllers/question.go
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
		db = db.Where("title LIKE ? OR tags LIKE ? OR related_company LIKE ?", "%"+keyword+"%", "%"+keyword+"%", "%"+keyword+"%")
	}

	// 计算总记录数
	var totalCount int64
	if err := db.Model(&models.Question{}).Count(&totalCount).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 计算总页数
	totalPages := int(totalCount) / pageSize
	if int(totalCount)%pageSize != 0 {
		totalPages++
	}

	result := db.Offset(offset).Limit(pageSize).Find(&questions)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"questions": questions, "total": totalPages})
}

// GetQuestionDetail 获取单个问题的详情
func GetQuestionDetail(c *gin.Context) {
	// 获取 URL 参数中的问题 ID
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid question ID"})
		return
	}

	var question models.Question
	result := utils.DB.First(&question, id)
	if result.Error != nil {
		if result.Error.Error() == "record not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": "Question not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, question)
}
