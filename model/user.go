package model

import "time"

type User struct {
	ID        uint      `json:"id" gorm:"primaryKey"` //主キーになる
	Email     string    `json:"email" gorm:"unique"`  //重複を許さない
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserResponse struct {
	ID    uint   `json:"id" gorm:"primaryKey"`
	Email string `json:"email" gorm:"unique"`
}
