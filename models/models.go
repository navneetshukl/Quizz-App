package models

import "gorm.io/gorm"

type QUESTIONS struct {
	gorm.Model
	Category string `json:"category"`
	Question string `json:"question"`
	Option1  string `json:"option1"`
	Option2  string `json:"option2"`
	Option3  string `json:"option3"`
	Option4  string `json:"option4"`
	Answer   string `json:"answer"`
}
