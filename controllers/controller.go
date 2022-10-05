package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sort"

	"github.com/ShivenduTyagii/GO-MTA/models"
)

func GetMtaData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	mtaobj := models.GetMtaobj()
	allmtas := mtaobj.GetAllmtas()
	sort.Strings(allmtas)
	fmt.Println(allmtas)
	json.NewEncoder(w).Encode(allmtas)
}
