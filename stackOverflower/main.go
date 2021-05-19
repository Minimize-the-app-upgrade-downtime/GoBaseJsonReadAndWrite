package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

)

//User defines model for storing account details in database
type db_field struct {

	// employee
	EmployeeNumber  int    `json:"employeeNumber,omitempty",`
	FirstName       string `json:"firstName,omitempty",`
	LastName        string `json:"lastName,omitempty",`
	Email  			string `json:"email,omitempty",`
	Name 	 		string `json:"name,omitempty",`
	Job  		string `json:"job,omitempty",` 
	Address  		string `json:"address,omitempty",`
}

func main(){
	fmt.Println("Schema change apply is up .")
	mux := http.NewServeMux()

	mux.HandleFunc("/", changeRequestForUpdatedVersion)
	
	http.ListenAndServe(":50000", mux)
}

func changeRequestForUpdatedVersion(w http.ResponseWriter, r *http.Request){
	
	log.Println("POST request received from localhost:50000")
	db := db_field{} //initialize 

	//Parse json request body and use it to set fields on db
	err := json.NewDecoder(r.Body).Decode(&db)
	
	if err != nil{
		panic(err)
	}
	log.Println("Apply DB changes Shcema for POST Request")
	
	// change body
	db.Email = ""
	db.Address = "No.xx, xxx,"
	db.Name = db.FirstName + db.LastName
	db.Email = ""
	
	fmt.Println(db)
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