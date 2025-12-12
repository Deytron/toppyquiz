package main

import "gorm.io/gorm"

type Question struct {
	gorm.Model
	Text   string
	Answer string
	Media  string // Media link, can be empty, or link to image/video or sound
	Theme  Theme  // Foreign key relation, 0 = any theme, any other number = specific theme
}

type Theme struct {
	gorm.Model
	Name string
	Icon string // Media icon link
}

type Team struct {
	gorm.Model
	Name  string
	Score int
}

type ItemType struct {
	gorm.Model
	Name  string
	Price int
}

type Item struct {
	gorm.Model
	TeamID   uint
	ItemType ItemType // Foreign key relation
	Quantity int
}

type Quiz struct {
	gorm.Model
	Name string
	Questions []Question `gorm:"many2many:quiz_questions;"`
}