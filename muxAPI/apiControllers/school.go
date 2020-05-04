package apiControllers

import (
	"encoding/json"
	"fmt"
	"github.com/MuxAPI/dbConfig"
	"github.com/MuxAPI/models"
	"github.com/gorilla/mux"
	"net/http"
)

func GetSchools(w http.ResponseWriter, r *http.Request) {
	var schools []models.School

	w.Header().Set("Content-Type", "application/json")
	dbConfig.DB.Order("id").Find(&schools)
	_ = json.NewEncoder(w).Encode(schools)
}

func GetSchoolsById(w http.ResponseWriter, r *http.Request) {
	var school models.School

	w.Header().Set("Content-Type", "application/json")
	param := mux.Vars(r)
	err := dbConfig.DB.Where("id =?", param["id"]).Find(&school).Error
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	_ = json.NewEncoder(w).Encode(&school)
}

func PostSchools(w http.ResponseWriter, r *http.Request) {
	var school models.School

	if dbConfig.DB.NewRecord(&school) {
		_ = json.NewDecoder(r.Body).Decode(&school)
		dbConfig.DB.Create(&school)
		w.Header().Set("Content-Type", "application/text")
		_, _ = w.Write([]byte("Kayıt işlemi başarılı!\n"))
		/*_ = json.NewEncoder(w).Encode(&school)*/
	}
}

func PutSchoolsById(w http.ResponseWriter, r *http.Request) {
	var school models.School

	w.Header().Set("Content-Type", "application/json")
	parameter := mux.Vars(r)

	err := dbConfig.DB.Where("id =?", parameter["id"]).Find(&school).Update(&school).Error
	if err != nil {
		w.Header().Set("Content-Type", "application/text")
		_, _ = w.Write([]byte("Kayıt bulunamadı!\n"))
		fmt.Println(err.Error())
		return
	}
	_ = json.NewDecoder(r.Body).Decode(&school)
	dbConfig.DB.Save(&school)
	w.Header().Set("Content-Type", "application/text")
	_, _ = w.Write([]byte("Güncelleme işlemi başarılı!\n"))
}

func DeleteSchoolsById(w http.ResponseWriter, r *http.Request) {
	var school models.School

	param := mux.Vars(r)
	err := dbConfig.DB.Where("id =?", param["id"]).Find(&school).Delete(&school).Error
	if err != nil {
		w.Header().Set("Content-Type", "application/text")
		_, _ = w.Write([]byte("Kayıt bulunamadı!\n"))
		fmt.Println(err.Error())
		return
	}
	w.Header().Set("Content-Type", "application/text")
	_, _ = w.Write([]byte("Silme işlemi Başarılı!\n"))
}
