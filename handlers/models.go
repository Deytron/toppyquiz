package handlers

import "gorm.io/gorm"

/* =======================
   QUIZ
======================= */

type Quiz struct {
	gorm.Model
	Name        string
	Description string

	Teams  []Team  `gorm:"constraint:OnDelete:CASCADE;"`
	Themes []Theme `gorm:"constraint:OnDelete:CASCADE;"`
}

/* =======================
   THEMES & QUESTIONS
======================= */

type Theme struct {
	gorm.Model
	Name string
	Icon string

	QuizID uint
	Quiz   Quiz `gorm:"constraint:OnDelete:CASCADE;"`

	Questions []Question `gorm:"constraint:OnDelete:CASCADE;"`
}

type Question struct {
	gorm.Model
	Text                  string
	Answer                string
	ShowMediaWithQuestion bool // Show media when displaying question
	TextSpeed             int

	ThemeID uint
	Theme   Theme `gorm:"constraint:OnDelete:CASCADE;"`
}

/* =======================
   TEAMS & MEMBERS
======================= */

type Team struct {
	gorm.Model
	Name  string
	Score int

	QuizID uint
	Quiz   Quiz `gorm:"constraint:OnDelete:CASCADE;"`

	Members []Member `gorm:"constraint:OnDelete:CASCADE;"`
	Items   []Item   `gorm:"constraint:OnDelete:CASCADE;"`
}

type Member struct {
	gorm.Model
	Name        string
	Description string

	TeamID uint
	Team   Team `gorm:"constraint:OnDelete:CASCADE;"`
}

/* =======================
   ITEMS
======================= */

type ItemType struct {
	gorm.Model
	Name  string
	Price int
}

type Item struct {
	gorm.Model
	Quantity int

	TeamID uint
	Team   Team `gorm:"constraint:OnDelete:CASCADE;"`

	ItemTypeID uint
	ItemType   ItemType `gorm:"constraint:OnDelete:RESTRICT;"`
}

/* =======================
   QUESTION EDIT FORM
======================= */

type QuizEditForm struct {
	Questions    []QuestionEditForm `form:"questions"`
	NewTheme     *NewThemeForm      `form:"new_theme"`
	NewQuestions []NewQuestionForm  `form:"new_questions"`
}

type QuestionEditForm struct {
	ID      uint   `form:"id"`
	ThemeID uint   `form:"theme_id"`
	Text    string `form:"text" binding:"required"`
}

type NewQuestionForm struct {
	ThemeID uint   `form:"theme_id" binding:"required"`
	Text    string `form:"text" binding:"required"`
}

type NewThemeForm struct {
	Name string `form:"name" binding:"required"`
}
