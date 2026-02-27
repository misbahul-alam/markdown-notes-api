package model

import (
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID         uint           `gorm:"primaryKey" json:"id"`
	FirstName  string         `gorm:"size:20;not null" json:"first_name"`
	LastName   string         `gorm:"size:20;not null" json:"last_name"`
	Email      string         `gorm:"size:255;uniqueIndex;not null" json:"email"`
	Password   string         `gorm:"size:255;not null" json:"-"`
	IsVerified bool           `gorm:"default:true" json:"is_verified"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}

func (u *User) BeforeSave(tx *gorm.DB) (err error) {
	if u.Password != "" {
		hashed, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		u.Password = string(hashed)
	}
	return nil
}

func (u *User) CheckPassword(plain string) bool {
	if plain == "" {
		return false
	}
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(plain))
	return err == nil
}
