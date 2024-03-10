package model

import "time"

type Task struct {
	ID        uint      `json:"id" gorm:"primaryKey"`  //主キーになる
	Title     string    `json:"title" gorm:"not null"` //空の値を許可しない
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	User      User      `json:"user" gorm:"foreignKey:UserId; constraint:OnDelete:CASCADE"` //userを削除したときにuserに紐づいているタスクも消去される
	UserId    uint      `json:"user_id" gorm:"not null"`
}

type TaskResponse struct {
	ID        uint      `json:"id" gorm:"primaryKey"`  //主キーになる
	Title     string    `json:"title" gorm:"not null"` //空の値を許可しない
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
