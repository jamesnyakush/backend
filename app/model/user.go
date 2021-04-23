package model

import (
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID                   uuid.UUID
	FirstName            string
	LastName             string
	Email                string `gorm:"not null; unique"`
	Password             string `gorm:"not null"`
	Phone                string `gorm:"not null;unique"`
	RoleId               uuid.UUID
	CountyCode           string
	ImageUrl             string
	PushNotificationKey  string
	InviteCode           string `gorm:"not null;unique"`
	Verified             bool   `gorm:"not null"`
	VerificationExpires  string `gorm:"not null"`
	VerificationToken    uuid.UUID
	ResetPasswordExpires time.Time
	ResetPasswordToken   string
	gorm.Model
}

type Role struct {
	RoleId   uuid.UUID
	RoleName string `gorm:"not null; unique"`
	gorm.Model
}

type Permission struct {
	PermissionId   uuid.UUID
	PermissionName string `gorm:"not null; unique"`
	RoleId         uuid.UUID
	gorm.Model
}
