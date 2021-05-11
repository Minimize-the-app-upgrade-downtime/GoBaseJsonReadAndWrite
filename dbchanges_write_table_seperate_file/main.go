package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	
)

// office struct
type offices struct{
	Offices []office `json:"offices"`
}

type office struct{	
	OfficeCode string `json:"officeCode,omitempty"`
	OfficeName string `json:"officeName,omitempty"`
	City string `json:city,omitempty`
	Phone string `json:phone,omitempty`
	AddressLine1  string `json:"addressLine1,omitempty"`
	Address string `json:"address,omitempty"`
	AddressLine2 string `json:"addressLine2,omitempty"`
	PostalCode string `json:"postalCode,omitempty"`
	Street string `json:"street,omitempty"`
	Country string `json:"country,omitempty"`
}

// employee struct
type employees struct{
	Employees []employee `json:"employees,omitempty"`
}

type employee struct{
	EmployeeNumber  int `json:"employeeNumber,omitempty",`
	FirstName  string `json:"firstName,omitempty",`
	LastName  string `json:"lastName,omitempty",`
	Email  string `json:"email,omitempty",`
	OfficeCode  string `json:"officeCode,omitempty",`
	JobTitle  string `json:"jobTitle,omitempty",`
	PhoneNumber  string `json:"phoneNumber,omitempty",`
}

func main(){

	jsonFile, err := os.Open("request.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened users.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue ,_ := ioutil.ReadAll(jsonFile)

	// get office varible 
	var office offices
	json.Unmarshal(byteValue, &office)
	// fmt.Println(office)
	for i := 0; i < len(office.Offices); i++ {
        office.Offices[i].Address =   office.Offices[i].AddressLine1
		office.Offices[i].AddressLine1 = ""
		office.Offices[i].AddressLine2 = ""
		if len(office.Offices[i].PostalCode) >= 30{
			office.Offices[i].PostalCode = "00000"
		}  
		fmt.Println(office)
		file, _ := json.MarshalIndent(office, "", " ")
		 _ = ioutil.WriteFile("office.json", file, 0644)
	}

	// employ obj
	var employee employees
	json.Unmarshal(byteValue, &employee)
	for i := 0; i < len(employee.Employees); i++ {
        employee.Employees[i].PhoneNumber = "021000000"  // add defult value
		//fmt.Println(emp)
		file, _ := json.MarshalIndent(employee, "", " ")
		_ = ioutil.WriteFile("employee.json", file, 0644)
	}

		
}     

