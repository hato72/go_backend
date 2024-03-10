package repository

import (
	"go-rest/model"

	"gorm.io/gorm"
)

type IUserRespository interface {
	GetUserByEmail(user *model.User, email string) error //ユーザーオブジェクトのポインタ　検索したいemail
	CreateUser(user *model.User) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) IUserRespository { //コンストラクタ
	return &userRepository{db}
}

func (ur *userRepository) GetUserByEmail(user *model.User, email string) error {
	if err := ur.db.Where("email=?", email).First(user).Error; err != nil { //emailが引数でうけとった値に一致するユーザーを探す
		return err
	}
	return nil
}

func (ur *userRepository) CreateUser(user *model.User) error {
	if err := ur.db.Create(user).Error; err != nil { //ユーザーの作成
		return err
	}
	return nil
}
