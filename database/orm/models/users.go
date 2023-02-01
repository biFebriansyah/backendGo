package models

import "time"

type User struct {
	UserId    string    `gorm:"type:uuid; primaryKey; default:uuid_generate_v4()" json:"id,omitempty" valid:"-"`
	Username  string    `json:"username" valid:"type(string)"`
	Email     string    `json:"email" valid:"email,optional"`
	Role      string    `json:"role,omitempty" valid:"type(string)"`
	Password  string    `json:"password,omitempty" valid:"type(string)"`
	CreatedAt time.Time `json:"created_at" valid:"-"`
	UpdateAt  time.Time `json:"update_at" valid:"-"`
}

type Users []User

func (User) TableName() string {
	return "users"
}
