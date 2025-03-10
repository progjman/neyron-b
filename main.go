package main

import (
	"fmt"
	"neyron-b/config"
	"neyron-b/models"
	"neyron-b/routes"
)

func main() {
	config.ConnectDB()                       //Подключение БД
	config.DB.AutoMigrate(&models.Account{}) //  Авто создание таблиц
	fmt.Println("База данных подключена, таблицы созданы.")

	r := routes.SetupRouter()
	fmt.Println("Сервер запущен на порту 8080")
	r.Run(":8080") //Запуск сервера
}
