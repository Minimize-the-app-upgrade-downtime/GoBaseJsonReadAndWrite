package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	// "net/http/httputil"
	// "net/url"
)

//User defines model for storing account details in database
type db_field struct {
	
	// office
	OfficeCode 		string `json:"officeCode,omitempty"`
	OfficeName 		string `json:"officeName,omitempty"`
	City 			string `json:"city,omitempty"`
	AddressLine1  	string `json:"addressLine1,omitempty"`
	Address 		string `json:"address,omitempty"`
	AddressLine2 	string `json:"addressLine2,omitempty"`
	PostalCode 		string `json:"postalCode,omitempty"`
	Street 			string `json:"street,omitempty"`
	Country 		string `json:"country,omitempty"`

	// employee
	EmployeeNumber  int    `json:"employeeNumber,omitempty",`
	FirstName       string `json:"firstName,omitempty",`
	LastName        string `json:"lastName,omitempty",`
	Email  			string `json:"email,omitempty"`
	JobTitle  		string `json:"job,omitempty",`
	PhoneNumber     string `json:"phonenumber,omitempty",`
}

func main(){
	mux := http.NewServeMux()

	mux.HandleFunc("/", changeRequestForUpdatedVersion)

	http.ListenAndServe(":50002", mux)
}

func changeRequestForUpdatedVersion(w http.ResponseWriter, r *http.Request){
	db := db_field{} //initialize 

	// log.Println(r)
	//Parse json request body and use it to set fields on db
	err := json.NewDecoder(r.Body).Decode(&db)
	
	if err != nil{
		panic(err)
	}
	// apply schema changes 
	db.Address = db.AddressLine1
	db.AddressLine1 = ""
	db.AddressLine2 = ""
	
	db.PhoneNumber = "0710000000"
	pc := db.PostalCode  // change shema bind the value
	db.PostalCode = pc
	
	//Marshal or convert user object back to json and write to response 
	dbJson, err := json.Marshal(db)
	if err != nil{
		panic(err)
	}
	
	responseBody := bytes.NewBuffer(dbJson)

	// create a post request
	url := "http://localhost:3000" + r.URL.String()
	//log.Println(url)
	resp, err := http.Post(url, "application/json", responseBody)

	if err != nil {
		log.Fatalf("An Error Occured %v", err)
	 }
	 defer resp.Body.Close()

	

	
}