package model

import (
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	UserId               uuid.UUID
	FirstName            string
	LastName             string
	Email                string `gorm:"not null; unique"`
	Password             string `gorm:"not null"`
	Phone                string `gorm:"not null;unique"`
	CountyCode           string
	ImageUrl             string
	PushNotificationKey  string
	InviteCode           string
	Verified             bool   `gorm:"not null"`
	VerificationExpires  string `gorm:"not null"`
	VerificationToken    uuid.UUID
	ResetPasswordExpires time.Time
	ResetPasswordToken   string
}

type Role struct {
	RoleId   uuid.UUID
	RoleName string `gorm:"not null; unique"`
}

type Permission struct {
	PermissionId   uuid.UUID
	PermissionName string `gorm:"not null; unique"`
}
