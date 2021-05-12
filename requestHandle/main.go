package main

import (
	"bytes"
	"encoding/json"

	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	
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


func test (rw http.ResponseWriter, req *http.Request){

		//log.Println(req.Body)
		body, err := ioutil.ReadAll(req.Body)
		if err != nil {
			panic(err)
		}
		log.Println(string(body))
		
		
		// create a office_emp object
		var oe office_emp
		json.Unmarshal(body, &oe) // bind the json value
		
		// err := json.NewDecoder(req.Body).Decode()

		log.Println(oe)
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


		payloadBuf := new(bytes.Buffer)
		json.NewEncoder(payloadBuf).Encode(oe)
		req,_ = http.NewRequest("POST","http://localhost:3000",payloadBuf)
		log.Println(req)

		link := "http://localhost:3000"
		// parse url
		url, _ := url.Parse(link)
		// create the reverse proxy
		proxy := httputil.NewSingleHostReverseProxy(url)
		//fmt.Println(proxy)
		// serveHttp is non blocking and uses a go routine under the hood
		proxy.ServeHTTP(rw, req)

	
	
}
func main(){

	http.HandleFunc("/",test)
	log.Fatal(http.ListenAndServe(":50002",nil))
}