package utils

import (
	"fmt"
	"github.com/chejo343/go_contacts/models"
)

// MigrateDB migra la base de datos

func MigrateDB(){

	db:= GetConnection()
	defer db.Close()

	fmt.Println("Migrando los modelos ........................")
	
	// Automigrate se encarga de migrar la base de datos sí no se ha migrado, y lo hace a partir del modelo

	db.AutoMigrate(&models.Contact{})
}