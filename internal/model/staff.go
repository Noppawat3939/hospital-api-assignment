package model

import "time"

type Staff struct {
	ID         uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	HospitalID string    `gorm:"type:uuid;not null;index:idx_staff_hospital_id_username,priority:1" json:"hospital_id"`
	Username   string    `gorm:"not null;index:idx_staff_hospital_id_username,priority:2" json:"username"`
	Password   string    `gorm:"not null" json:"-"`
	CreatedAt  time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
