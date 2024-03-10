package main

import (
	"go-rest/controller"
	"go-rest/db"
	"go-rest/repository"
	"go-rest/router"
	"go-rest/usecase"
	"go-rest/validator"
)

func main() {
	db := db.NewDB()
	userValidator := validator.NewUserValidator()
	taskValidator := validator.NewTaskValidator()
	userRepository := repository.NewUserRepository(db) //各コンストラクタ起動
	taskRepository := repository.NewTaskRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository, userValidator)
	taskUsecase := usecase.NewTaskUsecase(taskRepository, taskValidator)
	userController := controller.NewUserController(userUsecase)
	taskController := controller.NewTaskController(taskUsecase)
	e := router.NewRouter(userController, taskController)
	e.Logger.Fatal(e.Start(":8080")) //サーバー起動
	//docker,pgAdminを起動->docker compose up -d -> bashでGO_ENV=dev go run main.go
}
