package models

import (
	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserLink struct {
	gorm.Model
	ID       uuid.UUID `json:"uuid"`
	Name     string    `json:"name"`
	Username string    `json:"username"`
	Url      string    `json:"url"`
}

type GetUserRequest struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type User struct {
	gorm.Model
	ID uuid.UUID `gorm:"type:char(36);primary_key"`
	// ID       int64     `json:"id"`
	// UUID     uuid.UUID `json:"uuid" gorm:"type:uuid;primaryKey"`
	Name       string `json:"name,omitempty"`
	Email      string `json:"email,omitempty"`
	Phone      string `json:"phone,omitempty"`
	Psword     string `json:"psword,omitempty"`
	Verified   bool   `json:"verified,omitempty"`
	Role       string `json:"role,omitempty" gorm:"role"`
	ProfilePic string `json:"profilePic,omitempty" gorm:"profile_pic"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New()
	return
}

func (u *User) BeforeDelete(tx *gorm.DB) (err error) {
	if u.Role == "admin" {
		return errors.New("you do not have permission to delete user")
	}
	return
}

func (u *User) BeforeUpdate(tx *gorm.DB) (err error) {
	if u.Role == "admin" {
		return errors.New("you do not have permission to udpate user")
	}
	return
}

type UploadFileRequest struct {
	ID   string `json:"id"`
	File string `json:"file"`
}
