package models

import "grom.io/gorm"

type Feedbacks struct {
	gorm.Model
	RatingId    int    `gorm:"primary, autoIncrement" json:"id"`
	value       int    `json:"value"`
	Description string `json:"descripntion"`
	CustomerID  int    `json:"lastName"`
}
