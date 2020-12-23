package cats

import (
	cati "MeowGoWithDB/services/cat/interface"
	"log"

	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

//CatRepository struct
type CatRepository struct {
	Connection *gorm.DB
}

//Repo for call Repository
var Repo CatRepository

//InitCatRepository func
func InitCatRepository(connection *gorm.DB) {
	Repo = CatRepository{
		Connection: connection,
	}
}

//GetAll Func
func GetAll(w http.ResponseWriter, r *http.Request) {

	var result []cati.Cat

	Repo.Connection.Table("CattoHouse").Where("isDeleted = ?", false).Find(&result)

	// This for relation fetch
	// for index, data := range result {
	// 	Repo.Connect.Table("relationTable").Where("id = ?",data.relationID).Find(&result[index].relation)
	// }

	json.NewEncoder(w).Encode(result)
}

//GetAllWithDeleted Func
func GetAllWithDeleted(w http.ResponseWriter, r *http.Request) {

	var result []cati.Cat

	Repo.Connection.Table("CattoHouse").Find(&result)

	// This for relation fetch
	// for index, data := range result {
	// 	Repo.Connect.Table("relationTable").Where("id = ?",data.relationID).Find(&result[index].relation)
	// }

	json.NewEncoder(w).Encode(result)
}

//GetByID Func
func GetByID(w http.ResponseWriter, r *http.Request) {

	id := mux.Vars(r)["id"]

	var result cati.Cat

	Repo.Connection.Table("CattoHouse").Where("isDeleted = ? AND id = ?", false, id).Find(&result)

	// This for relation fetch
	// 	Repo.Connect.Table("relationTable").Where("id = ?",result.relationID).Find(&result.relation)

	if result.ID != "" {
		json.NewEncoder(w).Encode(result)
	} else {
		json.NewEncoder(w).Encode(nil)
	}
}

//Create Func
func Create(w http.ResponseWriter, r *http.Request) {

	getBody := json.NewDecoder(r.Body)

	var body cati.Cat

	err := getBody.Decode(&body)
	if err != nil {
		log.Print(err.Error())
		return
	}

	Repo.Connection.Table("CattoHouse").Create(&body)

	// result := mckcat.Create(body)

	json.NewEncoder(w).Encode(body)
}

//Update Func
func Update(w http.ResponseWriter, r *http.Request) {

	id := mux.Vars(r)["id"]

	getBody := json.NewDecoder(r.Body)

	var body cati.Cat

	err := getBody.Decode(&body)
	if err != nil {
		log.Print(err.Error())
		return
	}

	Repo.Connection.Table("CattoHouse").Where("id = ?", id).Updates(body)

	json.NewEncoder(w).Encode(body)
}

//Delete Func
func Delete(w http.ResponseWriter, r *http.Request) {

	id := mux.Vars(r)["id"]

	Repo.Connection.Table("CattoHouse").Where("id = ?", id).Updates(map[string]interface{}{
		"isDeleted": true,
	})

	w.Write([]byte(id))
}
