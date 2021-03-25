package orm

import (
	"github.com/raja/argon2pw"
	"gorm.io/gorm"
)

// User is the user
type User struct {
	ID            uint   `gorm:"primaryKey;autoIncrement" json:"-"`
	Username      string `gorm:"uniqueIndex" json:"username"`
	Email         string `gorm:"unique" json:"email"`
	PlainPassword string `gorm:"-" json:"-"`
	Active        bool   `json:"active"`
	AfkRating     uint8  `json:"afkRating"`
	IsAdmin       bool   `json:"isAdmin"`
	Password      string `json:"-"`
}

// Auth checks a password against the user's one
func (u *User) Auth(password string) (bool, error) {
	if !u.Active {
		return false, nil
	}

	return argon2pw.CompareHashWithPassword(u.Password, password)
}

func (u *User) updatePassword(tx *gorm.DB) (err error) {
	if u.PlainPassword != "" {
		u.Password, err = argon2pw.GenerateSaltedHash(u.PlainPassword)
	}
	return
}

// BeforeCreate is a gorm hook
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	return u.updatePassword(tx)
}

// BeforeSave is a gorm hook
func (u *User) BeforeSave(tx *gorm.DB) (err error) {
	return u.updatePassword(tx)
}
