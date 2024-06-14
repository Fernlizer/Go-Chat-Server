// db/models/user.go
package models

import "gorm.io/gorm"

type UserBlock struct {
	gorm.Model
	UserBlockByID int  `gorm:"not null" json:"user_block_by_id"`
	UserBlockedID int  `gorm:"not null" json:"user_blocked_id"`
	Status        bool `gorm:"not null" json:"status"`

	UserBlock   User `gorm:"foreignKey:UserBlockByID"`
	UserBlocked User `gorm:"foreignKey:UserBlockedID"`
}
