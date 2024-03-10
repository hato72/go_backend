package validator

import (
	"go-rest/model"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type ITaskValidator interface {
	TaskValidate(task model.Task) error
}

type taskValidator struct{}

func NewTaskValidator() ITaskValidator { //taskValidatorのインスタンスを生成するためのコンストラクタ
	return &taskValidator{}
}

func (tv *taskValidator) TaskValidate(task model.Task) error {
	return validation.ValidateStruct(&task,
		validation.Field(
			&task.Title,
			validation.Required.Error("title is required"),            //titleに値が存在するか
			validation.RuneLength(1, 10).Error("limited max 10 char"), //1文字から10文字までの文字数になっているかどうか
		),
	)
}
