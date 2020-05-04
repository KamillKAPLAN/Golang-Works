package apiControllers

import (
	"encoding/json"
	"fmt"
	"github.com/MuxAPI/dbConfig"
	"github.com/MuxAPI/models"
	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
	"net/http"
)

func GetEmployees(w http.ResponseWriter, r *http.Request) {
	var employees []models.Employee

	w.Header().Set("Content-Type", "application/json")

	dbConfig.DB.Order("id").Preload("School").Preload("Certificate").Preload("Education").Find(&employees)
	//dbConfig.DB.Find(&employees)
	_ = json.NewEncoder(w).Encode(&employees)
}

func GetEmployeesById(w http.ResponseWriter, r *http.Request) {
	var employee models.Employee

	w.Header().Set("Content-Type", "application/json")
	parameter := mux.Vars(r)
	err := dbConfig.DB.Where("id =?", parameter["id"]).Find(&employee).Error
	if err != nil {
		w.Header().Set("Content-Type", "application/text")
		_, _ = w.Write([]byte("Kayıt Bulunamadı!\n"))
		return
	}
	_ = json.NewEncoder(w).Encode(&employee).Error()
}

func PostEmployees(w http.ResponseWriter, r *http.Request) {
	var employee = models.Employee{}

	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	decoder := schema.NewDecoder()
	err = decoder.Decode(&employee, r.PostForm)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	} else {
		if dbConfig.DB.NewRecord(&employee) {
			_ = json.NewDecoder(r.Body).Decode(&employee)
			dbConfig.DB.Create(&employee)
			w.Header().Set("Content-Type", "application/text")
			_, _ = w.Write([]byte("Kayıt işlemi başarılı!\n"))
			/*_ = json.NewEncoder(w).Encode(&employee)*/
		}
	}

}

func PutEmployeesById(w http.ResponseWriter, r *http.Request) {
	var employee = models.Employee{}

	w.Header().Set("Content-Type", "application/json")
	parameter := mux.Vars(r)

	err := dbConfig.DB.Where("id =?", parameter["id"]).Find(&employee).Update(&employee).Error
	if err != nil {
		w.Header().Set("Content-Type", "application/text")
		_, _ = w.Write([]byte("Kayıt bulunamadı!\n"))
		fmt.Println(err.Error())
		return
	}
	_ = json.NewDecoder(r.Body).Decode(&employee)
	dbConfig.DB.Save(&employee)
	w.Header().Set("Content-Type", "application/text")
	_, _ = w.Write([]byte("Güncelleme işlemi başarılı!\n"))
}

func DeleteEmployeesById(w http.ResponseWriter, r *http.Request) {
	var employee = models.Employee{}

	parameter := mux.Vars(r)
	err := dbConfig.DB.Where("id =?", parameter["id"]).Find(&employee).Delete(&employee).Error
	if err != nil {
		w.Header().Set("Content-Type", "application/text")
		_, _ = w.Write([]byte("Kayıt bulunamadı!\n"))
		fmt.Println(err.Error())
		return
	}
	w.Header().Set("Content-Type", "application/text")
	_, _ = w.Write([]byte("Silme işlemi Başarılı!\n"))
}
