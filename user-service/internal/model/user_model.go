package model

import "gorm.io/gorm"




type User struct {
    gorm.Model
    TelegramID int64  `json:"telegram_id" gorm:"unique;not null"`
    Firstname  string `json:"firstname" gorm:"type:varchar(100)"`
    Lastname   string `json:"lastname" gorm:"type:varchar(100)"`
    Role       string `json:"role" gorm:"type:varchar(20)"`
}