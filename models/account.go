package models

import "gorm.io/gorm"

// Account  таблица для хранения счетов пользователей.

type Account struct {
	gorm.Model
	Owner    string  `json:"owner"`
	Name     string  `json:"name"`
	Balance  float64 `json: "balance"`
	Currency string  `json: "currency"`
}
