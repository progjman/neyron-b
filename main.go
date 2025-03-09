package main

import (
	"fmt"
	"neyron-b/config"
	"neyron-b/models"
)

func main() {
	config.ConnectDB() //Подключение БД

	config.DB.AutoMigrate(&models.Account{})

	fmt.Println("База данных подключена.")
}
