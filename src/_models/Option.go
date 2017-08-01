package models

type Option struct {
	ID         int
	QuestionID int
	Text       string
	IsCorrect  bool
}
