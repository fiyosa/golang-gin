package model

import "time"

type Contact struct {
	Id        uint `json:"id" gorm:"primaryKey;autoIncrement"`
	UserId    uint
	FirstName string    `json:"first_name" gorm:"type:varchar(255);not null"`
	LastName  string    `json:"last_name" gorm:"type:varchar(255)"`
	Email     string    `json:"password" gorm:"type:varchar(255);not null"`
	Phone     string    `json:"phone" gorm:"type:varchar(255)"`
	CreatedAt time.Time `json:"created_at" gorm:"type:timestamp(0);not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `json:"updated_at" gorm:"type:timestamp(0);not null;default:CURRENT_TIMESTAMP"`

	User      User      `json:"user_id" gorm:"foreignKey:UserId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Addresses []Address `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}
