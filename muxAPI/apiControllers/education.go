package apiControllers

import (
	"encoding/json"
	"fmt"
	"github.com/MuxAPI/dbConfig"
	"github.com/MuxAPI/models"
	"github.com/gorilla/mux"
	"net/http"
)

func GetEducations(w http.ResponseWriter, r *http.Request) {
	var educations []models.Education

	w.Header().Set("Content-Type", "application/json")
	dbConfig.DB.Order("id").Find(&educations)
	_ = json.NewEncoder(w).Encode(educations)
}

func GetEducationsById(w http.ResponseWriter, r *http.Request) {
	var education models.Education

	w.Header().Set("Content-Type", "application/json")
	param := mux.Vars(r)
	err := dbConfig.DB.Where("id =?", param["id"]).Find(&education).Error
	if err != nil {
		w.Header().Set("Content-Type", "application/text")
		_, _ = w.Write([]byte("Kayıt bulunamadı!\n"))
		fmt.Println(err.Error())
		return
	}
	_ = json.NewEncoder(w).Encode(&education)
}

func PostEducations(w http.ResponseWriter, r *http.Request) {
	var education models.Education

	if dbConfig.DB.NewRecord(&education) {
		_ = json.NewDecoder(r.Body).Decode(&education)
		dbConfig.DB.Create(&education)
		w.Header().Set("Content-Type", "application/text")
		_, _ = w.Write([]byte("Kayıt işlemi başarılı!\n"))
		/*_ = json.NewEncoder(w).Encode(&school)*/
	}
}

func PutEducationsById(w http.ResponseWriter, r *http.Request) {
	var education models.Education

	w.Header().Set("Content-Type", "application/json")
	parameter := mux.Vars(r)

	err := dbConfig.DB.Where("id =?", parameter["id"]).Find(&education).Update(&education).Error
	if err != nil {
		w.Header().Set("Content-Type", "application/text")
		_, _ = w.Write([]byte("Kayıt bulunamadı!\n"))
		fmt.Println(err.Error())
		return
	}
	_ = json.NewDecoder(r.Body).Decode(&education)
	dbConfig.DB.Save(&education)
	w.Header().Set("Content-Type", "application/text")
	_, _ = w.Write([]byte("Güncelleme işlemi başarılı!\n"))
}

func DeleteEducationsById(w http.ResponseWriter, r *http.Request) {
	var education models.Education

	param := mux.Vars(r)
	err := dbConfig.DB.Where("id =?", param["id"]).Find(&education).Delete(&education).Error
	if err != nil {
		w.Header().Set("Content-Type", "application/text")
		_, _ = w.Write([]byte("Kayıt bulunamadı!\n"))
		fmt.Println(err.Error())
		return
	}
	w.Header().Set("Content-Type", "application/text")
	_, _ = w.Write([]byte("Silme işlemi Başarılı!\n"))
}
