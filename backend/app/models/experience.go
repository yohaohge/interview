package models

type InterviewExperience struct {
	ID                 int    `gorm:"primaryKey" json:"id"`
	Title              string `json:"title"`
	Content            string `json:"content"`
	RelatedQuestionIDs string `json:"related_question_ids"`
}

func (InterviewExperience) TableName() string {
	return "interview_experiences"
}
