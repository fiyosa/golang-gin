package model

import "time"

type Address struct {
	Id         uint `json:"id" gorm:"primaryKey;autoIncrement"`
	ContactId  uint
	Street     string    `json:"street" gorm:"type:varchar(255)"`
	City       string    `json:"city" gorm:"type:varchar(255)"`
	Province   string    `json:"province" gorm:"type:varchar(255)"`
	Country    string    `json:"country" gorm:"type:varchar(255);not null"`
	PostalCode string    `json:"postal_code" gorm:"type:varchar(255);not null"`
	CreatedAt  time.Time `json:"created_at" gorm:"type:timestamp(0);not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt  time.Time `json:"updated_at" gorm:"type:timestamp(0);not null;default:CURRENT_TIMESTAMP"`

	Contact Contact `json:"contact_id" gorm:"foreignKey:ContactId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}
