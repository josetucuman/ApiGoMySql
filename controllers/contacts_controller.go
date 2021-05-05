package controllers
import (
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/chejo343/go_contacts/models"
	"github.com/chejo343/go_contacts/utils"
	"github.com/gorilla/mux"
)

// GetContact obtiene un contacto por su ID

func GetContact(w http.ResponseWriter, r *http.Request){

	contact := models.Contact{}

	//se obtiene primero el parametro id de la URL

	id := mux.Vars(r)["id"]

	//Conexion a DB

	db := utils.GetConnection()

	defer db.Close()

	// Consulta a la DB - SELECT * FROM contacts WHERE ID = ?

	db.Find(&contact, id)

	// Se comprueba que exista el registro

	if contact.ID > 0{
		// se codifican los datos a JSON
		j, _ := json.Marshal(contact)

		//se envian los datos
		
		utils.SendResponse(w, http.StatusOK, j)
	} else {
		utils.SendErr(w, http.StatusNotFound)
	}
}

// GetContacts obtiene todos los contactos

func GetContacts(w http.ResponseWriter, r *http.Request){
	// Slice (array) donde se guardaran los datos

	contacts := []models.Contact{}

	//Conexion a la DB
	db := utils.GetConnection()

	defer db.Close()

	// Consulta a la DB - SELECT * FROM contacts
	db.Find(&contacts)

	// Se codifican los datos a formato JSON
	j, _ := json.Marshal(contacts)

	// Se envian los datos
	utils.SendResponse(w, http.StatusOK, j)
}

// StoreContact guarda un nuevo contacto
func StoreContact(w http.ResponseWriter, r *http.Request){

	//Estructura donde se guardan los datos del body
	contact := models.Contact{}

	//conexion a DB
	db := utils.GetConnection()

	defer db.Close()

	// Se decodifican los datos del body a la estructura contact
	err := json.NewDecoder(r.Body).Decode(&contact)
	if err != nil {
		// Sí hay algun error en los datos se devolvera un error 400
		fmt.Println(err)
		utils.SendErr(w, http.StatusBadRequest)
		return
	}

	// Se guardan los datos en la DB
	err = db.Create(&contact).Error

	if err != nil {
		// Sí hay algun error al guardar los datos se devolvera un error 500
		fmt.Println(err)
		utils.SendErr(w, http.StatusInternalServerError)
		return
	}

	//Se codifica el nuevo dato y se devuelve
	j, _ := json.Marshal(contact)
	utils.SendResponse(w, http.StatusCreated, j)
}

