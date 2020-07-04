package model

import "time"

type User struct {
	UserId               uint      `json:"user_id"`
	FirstName            string    `json:"first_name"`
	LastName             string    `json:"last_name"`
	Username             string    `json:"username"`
	Email                string    `json:"email"`
	Password             string    `json:"password"`
	Phone                string    `json:"phone"`
	CountyCode           string    `json:"county_code"`
	ImageUrl             string    `json:"image_url"`
	PushNotificationKey  string    `json:"push_notification_key"`
	InviteCode           string    `json:"invite_code"`
	Verified             bool      `json:"verified"`
	VerificationExpires  string    `json:"verification_expires"`
	VerificationToken    string    `json:"verification_token"`
	ResetPasswordExpires time.Time `json:"reset_password_expires"`
	ResetPasswordToken   string    `json:"reset_password_token"`
	CreatedAt            time.Time `json:"created_at"`
	UpdatedAt            time.Time `json:"updated_at"`
	DeletedAt            time.Time `json:"deleted_at"`
	Roles                Role      `json:"roles"`
}

type Role struct {
	RoleId      string       `json:"role_id"`
	RoleName    string       `json:"role_name"`
	Permissions []Permission `json:"permissions"`
}

type Permission struct {
	PermissionId   string `json:"permission_id"`
	PermissionName string `json:"permission_name"`
}
