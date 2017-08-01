package models

import "_models"

type Question struct {
	ID      int
	Text    string
	Options []_models.Option
}
