package models

import "time"

type Role struct {
	ID         uint   `gorm:"primaryKey;autoIncrement`
	Code       string `gorm:"varchar(15);not null`
	Name       string `gorm;"varchar(15);not null`
	Created_at *time.Time
	Updated_at *time.Time
}
