package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type office_emp struct{
	Offices 	[]office `json:"offices"`
	Employees 	[]employee `json:"employees,omitempty"`
}


type office struct{	
	OfficeCode 		string `json:"officeCode,omitempty"`
	OfficeName 		string `json:"officeName,omitempty"`
	City 			string `json:city,omitempty`
	Phone 			string `json:phone,omitempty`
	AddressLine1  	string `json:"addressLine1,omitempty"`
	Address 		string `json:"address,omitempty"`
	AddressLine2 	string `json:"addressLine2,omitempty"`
	PostalCode 		string `json:"postalCode,omitempty"`
	Street 			string `json:"street,omitempty"`
	Country 		string `json:"country,omitempty"`
}



type employee struct{
	EmployeeNumber  int    `json:"employeeNumber,omitempty",`
	FirstName       string `json:"firstName,omitempty",`
	LastName        string `json:"lastName,omitempty",`
	Email  			string `json:"email,omitempty"`
	OfficeCode 		string `json:"officeCode,omitempty",`
	JobTitle  		string `json:"jobTitle,omitempty",`
	PhoneNumber     string `json:"phoneNumber,omitempty",`
}

func main(){

	jsonFile, err := os.Open("request.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened request.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue ,_ := ioutil.ReadAll(jsonFile)

	// create a office_emp object
	var oe office_emp
	json.Unmarshal(byteValue, &oe) // bind the json value

	// update the employee json
	for i := 0; i < len(oe.Employees); i++ {
        oe.Employees[i].PhoneNumber = "021000000"  // add defult value
	}


	// update the office json
	for i := 0; i < len(oe.Offices); i++ {
		oe.Offices[i].Address =   oe.Offices[i].AddressLine1
		oe.Offices[i].AddressLine1 = ""
		oe.Offices[i].AddressLine2 = ""
		if len(oe.Offices[i].PostalCode) >= 30{
			oe.Offices[i].PostalCode = "00000"
		}  
	}

	// write file 
	file, _ := json.MarshalIndent(oe, "", " ")
 	_ = ioutil.WriteFile("test.json", file, 0644)
	 fmt.Println("Successfully Write test.json")
}     

