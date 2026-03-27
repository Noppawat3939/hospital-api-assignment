package model

import "time"

type Staff struct {
	ID         uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	HospitalID string    `gorm:"type:uuid;not null;uniqueIndex:idx_staff_hospital_username,priority:1" json:"hospital_id"`
	Hospital   Hospital  `gorm:"foreignKey:HospitalID;references:ID" json:"hospital"`
	Username   string    `gorm:"not null;uniqueIndex:idx_staff_hospital_username,priority:2" json:"username"`
	Password   string    `gorm:"not null" json:"-"`
	CreatedAt  time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
