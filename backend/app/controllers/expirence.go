// app/controllers/expirence.go
package controllers

import (
	"net/http"
	"strconv"

	"www.interview.com/models"
	"www.interview.com/utils"

	"github.com/gin-gonic/gin"
)

// GetInterviewExperienceList 获取面筋列表
func GetInterviewExperienceList(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))
	company := c.Query("company")

	var experiences []models.InterviewExperience
	offset := (page - 1) * pageSize

	db := utils.DB

	if company != "" {
		// 假设面筋表中关联公司信息在 related_question_ids 对应的题库中
		var questionIDs []int
		utils.DB.Where("related_company =?", company).Pluck("id", &questionIDs)
		if len(questionIDs) > 0 {
			for i, id := range questionIDs {
				questionIDs[i] = id
			}
			db = db.Where("FIND_IN_SET(?, related_question_ids) > 0", questionIDs)
		} else {
			experiences = []models.InterviewExperience{}
			c.JSON(http.StatusOK, experiences)
			return
		}
	}

	result := db.Offset(offset).Limit(pageSize).Find(&experiences)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"experiences": experiences, "total": 100})
}

// GetInterviewExperienceDetail 获取单个面经详情
func GetInterviewExperienceDetail(c *gin.Context) {
	// 获取 URL 参数中的面经 ID
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid interview experience ID"})
		return
	}

	var experience models.InterviewExperience
	result := utils.DB.First(&experience, id)
	if result.Error != nil {
		if result.Error.Error() == "record not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": "Interview experience not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, experience)
}
