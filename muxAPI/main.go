package main

import (
	"fmt"
	"github.com/MuxAPI/apiControllers"
	"net/http"

	"github.com/MuxAPI/dbConfig"
	"github.com/MuxAPI/models"
	"github.com/gorilla/mux"
)

var employees = []models.Employee{}
var schools = []models.School{}

func main() {
	r := mux.NewRouter()

	dbConfig.InitDB()

	if !dbConfig.DB.HasTable(models.Employee{}) {
		dbConfig.DB.CreateTable(models.Employee{})
		fmt.Println("Employee table created")
	}
	if !dbConfig.DB.HasTable(models.School{}) {
		dbConfig.DB.CreateTable(models.School{})
		fmt.Println("School table created")
	}
	if !dbConfig.DB.HasTable(models.Education{}) {
		dbConfig.DB.CreateTable(models.Education{})
		fmt.Println("Education table created")
	}
	if !dbConfig.DB.HasTable(models.Certificate{}) {
		dbConfig.DB.CreateTable(models.Certificate{})
		fmt.Println("Certificate table created")
	}

	s := r.PathPrefix("/api").Subrouter()

	// Employee GET POST PUT DELETE Transactions
	s.HandleFunc("/employees", apiControllers.GetEmployees).Methods("GET")
	s.HandleFunc("/employees/{id}", apiControllers.GetEmployeesById).Methods("GET")
	s.HandleFunc("/employees", apiControllers.PostEmployees).Methods("POST")
	s.HandleFunc("/employees/{id}", apiControllers.PutEmployeesById).Methods("PUT")
	s.HandleFunc("/employees/{id}", apiControllers.DeleteEmployeesById).Methods("DELETE")

	// School GET POST PUT DELETE Transactions
	s.HandleFunc("/schools", apiControllers.GetSchools).Methods("GET")
	s.HandleFunc("/schools/{id}", apiControllers.GetSchoolsById).Methods("GET")
	s.HandleFunc("/schools", apiControllers.PostSchools).Methods("POST")
	s.HandleFunc("/schools/{id}", apiControllers.PutSchoolsById).Methods("PUT")
	s.HandleFunc("/schools/{id}", apiControllers.DeleteSchoolsById).Methods("DELETE")

	// Certificate GET POST PUT DELETE Transactions
	s.HandleFunc("/certificates", apiControllers.GetCertificates).Methods("GET")
	s.HandleFunc("/certificates/{id}", apiControllers.GetCertificatesById).Methods("GET")
	s.HandleFunc("/certificates", apiControllers.PostCertificates).Methods("POST")
	s.HandleFunc("/certificates/{id}", apiControllers.PutCertificatesById).Methods("PUT")
	s.HandleFunc("/certificates/{id}", apiControllers.DeleteCertificatesById).Methods("DELETE")

	// Education GET POST PUT DELETE Transactions
	s.HandleFunc("/educations", apiControllers.GetEducations).Methods("GET")
	s.HandleFunc("/educations/{id}", apiControllers.GetEducationsById).Methods("GET")
	s.HandleFunc("/educations", apiControllers.PostEducations).Methods("POST")
	s.HandleFunc("/educations/{id}", apiControllers.PutEducationsById).Methods("PUT")
	s.HandleFunc("/educations/{id}", apiControllers.DeleteEducationsById).Methods("DELETE")

	http.ListenAndServe(":8000", r)
}
