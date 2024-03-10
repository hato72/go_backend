package repository

import (
	"fmt"
	"go-rest/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ITaskRepository interface {
	GetAllTasks(tasks *[]model.Task, userId uint) error           //作成したタスクの一覧を取得
	GetTaskById(task *model.Task, userId uint, taskId uint) error //引数のtaskIdに一致するタスクを返す
	CreateTask(task *model.Task) error                            //タスクの新規作成
	UpdateTask(task *model.Task, userId uint, taskId uint) error  //タスクの更新
	DeleteTask(userId uint, taskId uint) error                    //タスクの削除
}

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) ITaskRepository { //コンストラクタ
	return &taskRepository{db}
}

func (tr *taskRepository) GetAllTasks(tasks *[]model.Task, userId uint) error {
	if err := tr.db.Joins("User").Where("user_id=?", userId).Order("created_at").Find(tasks).Error; err != nil { //タスクの一覧から引数のユーザーidに一致するタスクを取得する　その時、作成日時があたらしいものが末尾に来るようにする
		return err
	}
	return nil
}

func (tr *taskRepository) GetTaskById(task *model.Task, userId uint, taskId uint) error {
	if err := tr.db.Joins("User").Where("user_id=?", userId).Find(task, taskId).Error; err != nil { //引数のユーザーidに一致するタスクを取得し、その中でtaskの主キーが引数で受け取ったtaskidに一致するタスクを取得する
		return err
	}
	return nil
}

func (tr *taskRepository) CreateTask(task *model.Task) error {
	if err := tr.db.Create(task).Error; err != nil {
		return err
	}
	return nil
}

func (tr *taskRepository) UpdateTask(task *model.Task, userId uint, taskId uint) error {
	result := tr.db.Model(task).Clauses(clause.Returning{}).Where("id=? AND user_id=?", taskId, userId).Update("title", task.Title)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 1 { //更新されたレコードの数を取得できる
		return fmt.Errorf("object does not exists")
	}
	return nil
}

func (tr *taskRepository) DeleteTask(userId uint, taskId uint) error {
	result := tr.db.Where("id=? AND user_id=?", taskId, userId).Delete(&model.Task{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 1 { //更新されたレコードの数を取得できる
		return fmt.Errorf("object does not exists")
	}
	return nil
}
