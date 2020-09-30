package cats

import (
	cati "MeowGoWithDB/services/cat/interface"
	mckcat "MeowGoWithDB/services/db/justmck"
	"log"
	"strconv"

	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

//GetAll Func
func GetAll(w http.ResponseWriter, r *http.Request) {

	cats := mckcat.GetAll()

	json.NewEncoder(w).Encode(cats)
}

//GetByID Func
func GetByID(w http.ResponseWriter, r *http.Request) {

	id := mux.Vars(r)["id"]

	getID, err := strconv.Atoi(id)

	if err != nil {
		log.Print(err.Error())
		return
	}

	cats := mckcat.GetByID(getID)

	json.NewEncoder(w).Encode(cats)
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

	result := mckcat.Create(body)

	json.NewEncoder(w).Encode(result)
}

//Update Func
func Update(w http.ResponseWriter, r *http.Request) {

	id := mux.Vars(r)["id"]

	getID, err := strconv.Atoi(id)

	if err != nil {
		log.Print(err.Error())
		return
	}

	getBody := json.NewDecoder(r.Body)

	var body cati.Cat

	err = getBody.Decode(&body)
	if err != nil {
		log.Print(err.Error())
		return
	}

	body.ID = getID

	result := mckcat.Update(body, getID)

	json.NewEncoder(w).Encode(result)
}

//Delete Func
func Delete(w http.ResponseWriter, r *http.Request) {

	id := mux.Vars(r)["id"]

	getID, err := strconv.Atoi(id)

	if err != nil {
		log.Print(err.Error())
		return
	}

	cats := mckcat.Delete(getID)

	json.NewEncoder(w).Encode(cats)
}
