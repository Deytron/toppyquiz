package handlers

import "gorm.io/gorm"

type Question struct {
	gorm.Model
	Text                  string
	Answer                string
	ShowMediaWithQuestion bool  // Whether to show media when displaying question
	TextSpeed             int   // Speed of text display in ms per character
	ThemeID               uint  // Foreign key // Media link, can be empty, or link to image/video or sound
	Theme                 Theme `gorm:"foreignKey:ThemeID"` // Foreign key relation, 0 = any theme, any other number = specific theme
}

type Theme struct {
	gorm.Model
	Name string
	Icon string // Media icon link
}

type Team struct {
	gorm.Model
	Name   string
	Score  int
	QuizID uint
	Quiz   Quiz   `gorm:"foreignKey:QuizID"`
	Items  []Item `gorm:"foreignKey:TeamID"`
}

type ItemType struct {
	gorm.Model
	Name  string
	Price int
}

type Item struct {
	gorm.Model
	TeamID     uint
	ItemTypeID uint
	ItemType   ItemType `gorm:"foreignKey:ItemTypeID"`
	Quantity   int
}

type Quiz struct {
	gorm.Model
	Name        string
	Description string
	Questions   []Question `gorm:"many2many:quiz_questions;"`
	Teams       []Team     `gorm:"many2many:leaderboard_teams;"`
	Themes      []Theme    `gorm:"many2many:leaderboard_themes;"`
}
