package utils

import (
	"log"
	"github.com/jinzhu/gorm"
	// mysql
_ 	"github.com/jinzhu/gorm/dialects/mysql"
)

// GetConnection obtiene una conexión a la base de datos

func GetConnection() *gorm.DB{
	db, err := gorm.Open("mysql", "root:root@/db_pruebas_apiGo?charset=utf-8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal(err)
	}

	return db
}