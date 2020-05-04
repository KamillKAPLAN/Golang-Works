package apiControllers

import (
	"encoding/json"
	"fmt"
	"github.com/MuxAPI/dbConfig"
	"github.com/MuxAPI/models"
	"github.com/gorilla/mux"
	"net/http"
)

func GetCertificates(w http.ResponseWriter, r *http.Request) {
	var certificates []models.Certificate

	w.Header().Set("Content-Type", "application/json")
	dbConfig.DB.Order("id").Find(&certificates)
	_ = json.NewEncoder(w).Encode(certificates)
}

func GetCertificatesById(w http.ResponseWriter, r *http.Request) {
	var certificate models.Certificate

	w.Header().Set("Content-Type", "application/json")
	param := mux.Vars(r)
	err := dbConfig.DB.Where("id =?", param["id"]).Find(&certificate).Error
	if err != nil {
		w.Header().Set("Content-Type", "application/text")
		_, _ = w.Write([]byte("Kayıt bulunamadı!\n"))
		fmt.Println(err.Error())
		return
	}
	_ = json.NewEncoder(w).Encode(&certificate)
}

func PostCertificates(w http.ResponseWriter, r *http.Request) {
	var certificate models.Certificate

	if dbConfig.DB.NewRecord(&certificate) {
		_ = json.NewDecoder(r.Body).Decode(&certificate)
		dbConfig.DB.Create(&certificate)
		w.Header().Set("Content-Type", "application/text")
		_, _ = w.Write([]byte("Kayıt işlemi başarılı!\n"))
		/*_ = json.NewEncoder(w).Encode(&school)*/
	}
}

func PutCertificatesById(w http.ResponseWriter, r *http.Request) {
	var certificate models.Certificate

	w.Header().Set("Content-Type", "application/json")
	parameter := mux.Vars(r)

	err := dbConfig.DB.Where("id =?", parameter["id"]).Find(&certificate).Update(&certificate).Error
	if err != nil {
		w.Header().Set("Content-Type", "application/text")
		_, _ = w.Write([]byte("Kayıt bulunamadı!\n"))
		fmt.Println(err.Error())
		return
	}
	_ = json.NewDecoder(r.Body).Decode(&certificate)
	dbConfig.DB.Save(&certificate)
	w.Header().Set("Content-Type", "application/text")
	_, _ = w.Write([]byte("Güncelleme işlemi başarılı!\n"))
}

func DeleteCertificatesById(w http.ResponseWriter, r *http.Request) {
	var certificate models.Certificate

	param := mux.Vars(r)
	err := dbConfig.DB.Where("id =?", param["id"]).Find(&certificate).Delete(&certificate).Error
	if err != nil {
		w.Header().Set("Content-Type", "application/text")
		_, _ = w.Write([]byte("Kayıt bulunamadı!\n"))
		fmt.Println(err.Error())
		return
	}
	w.Header().Set("Content-Type", "application/text")
	_, _ = w.Write([]byte("Silme işlemi Başarılı!\n"))
}
