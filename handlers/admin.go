package handlers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AdminHandler(c *gin.Context) {
	// GET quiz list from sql db with GoRm
	var quizzes []Quiz
	Db.Find(&quizzes)

	// Render page with data
	c.HTML(http.StatusOK, "admin.tmpl", gin.H{
		"title":   "Tzatziquiz Administration",
		"quizzes": quizzes,
	})
}

func AdminTeamHandler(c *gin.Context) {
	// get quiz id from param
	quizId := c.Param("id")
	quizInt, err := strconv.Atoi(quizId)
	if err != nil {
		c.String(http.StatusBadRequest, "Invalid quiz ID")
		return
	}

	// Get to create team, POST to add to DB
	switch c.Request.Method {
	case "POST":
		teamName := c.PostForm("team-name")
		newTeam := Team{
			Name:   teamName,
			Score:  0,
			QuizID: uint(quizInt),
		}
		Db.Create(&newTeam)
	case "GET":
		// Get quizzes for team creation form
		var quizzes []Quiz
		Db.Find(&quizzes)
	}

	// Render page with data
	c.HTML(http.StatusOK, "admin-team.tmpl", gin.H{
		"title": "Tzatziquiz Administration",
	})
}

func AdminQuizHandler(c *gin.Context) {
	// Get to create team, POST to add to DB
	switch c.Request.Method {
	case "POST":
		quiz := Quiz{
			Name:        c.PostForm("quiz-name"),
			Description: c.PostForm("quiz-desc"),
		}
		Db.Create(&quiz)
		c.Redirect(http.StatusFound, "/admin")
		return
	}

	// Render page with data
	c.HTML(http.StatusOK, "admin-quiz.tmpl", gin.H{
		"title": "Tzatziquiz Administration",
	})
}

func AdminQuizEditHandler(c *gin.Context) {
	// get quiz id from param
	quizString := c.Param("id")
	action := c.PostForm("action")
	quizId, _ := strconv.Atoi(quizString)
	var err error

	var quiz Quiz
	if err := Db.
		Preload("Themes.Questions").
		First(&quiz, quizId).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	if c.Request.Method == http.MethodPost {

		err = Db.Transaction(func(tx *gorm.DB) error {

			switch action {

			case "update_questions":
				var form struct {
					Questions []QuestionEditForm `form:"questions" binding:"required"`
				}

				if err := c.ShouldBind(&form); err != nil {
					return err
				}

				for _, q := range form.Questions {
					if err := tx.Model(&Question{}).
						Where("id = ?", q.ID).
						Update("text", q.Text).Error; err != nil {
						return err
					}
				}

			case "add_theme":
				var form struct {
					NewTheme *NewThemeForm `form:"new_theme" binding:"required"`
				}

				if err := c.ShouldBind(&form); err != nil {
					return err
				}

				return tx.Create(&Theme{
					Name:   form.NewTheme.Name,
					QuizID: quiz.ID,
				}).Error

			case "add_question":
				var form struct {
					NewQuestions []NewQuestionForm `form:"new_questions" binding:"required"`
				}

				if err := c.ShouldBind(&form); err != nil {
					return err
				}

				for _, q := range form.NewQuestions {
					var count int64
					if err := tx.Model(&Theme{}).
						Where("id = ? AND quiz_id = ?", q.ThemeID, quiz.ID).
						Count(&count).Error; err != nil {
						return err
					}

					if count == 0 {
						return errors.New("invalid theme")
					}

					if err := tx.Create(&Question{
						Text:    q.Text,
						ThemeID: q.ThemeID,
					}).Error; err != nil {
						return err
					}
				}

			default:
				return errors.New("unknown action")
			}

			return nil
		})

		if err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		if err := Db.
			Preload("Themes.Questions").
			First(&quiz, quizId).Error; err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
	}

	var questions []Question
	for _, t := range quiz.Themes {
		questions = append(questions, t.Questions...)
	}

	c.HTML(http.StatusOK, "edit-quiz.tmpl", gin.H{
		"title":     "Édition du quiz",
		"QID":       quizId,
		"Quiz":      quiz,
		"Questions": questions,
	})
}
