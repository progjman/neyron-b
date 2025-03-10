package handlers

import (
	"net/http"
	"neyron-b/config"
	"neyron-b/models"

	"github.com/gin-gonic/gin"
)

func CreateAccount(c *gin.Context) { // Создание аккаунта
	var account models.Account
	if err := c.ShouldBindJSON(&account); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Create(&account)
	c.JSON(http.StatusCreated, account)
}

func GetAccounts(c *gin.Context) { //Получить все аккаунты
	var accounts []models.Account
	config.DB.Find(&accounts)
	c.JSON(http.StatusOK, accounts)
}

func GetAccountByID(c *gin.Context) { //Получить аккаунд по id
	var account models.Account
	id := c.Param("id")

	if err := config.DB.First(&account, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Аккаунд не найден"})
		return
	}

	c.JSON(http.StatusOK, account)
}

func UpdateAccount(c *gin.Context) { //Обновить данные аккаунта
	var account models.Account
	id := c.Param("id")

	if err := config.DB.First(&account, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Аккаунд не найден"})
		return
	}

	if err := c.ShouldBindJSON(&account); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Save(&account)
	c.JSON(http.StatusOK, account)
}

func DeleteAccount(c *gin.Context) { //Удалить аккаунт
	var account models.Account
	id := c.Param("id")

	if err := config.DB.First(&account, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Аккаунд не найден"})
		return
	}

	config.DB.Delete(&account)
	c.JSON(http.StatusOK, gin.H{"message": "Аккаунт удалён"})
}
