package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type Employees struct {
	Employee []Employee `json:"employee"`
}

type Employee struct {
	Name      string `json:"name,omitempty"`
	Age       int    `json:"age"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

func main() {

	jsonFile, err := os.Open("name_request.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened users.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var emp Employees
	json.Unmarshal(byteValue, &emp)

	for i := 0; i < len(emp.Employee); i++ {

		name := strings.Fields(emp.Employee[i].Name)
		emp.Employee[i].Name = ""           // remove name
		emp.Employee[i].FirstName = name[0] // add first name
		emp.Employee[i].LastName = name[1]  // add last name

		fmt.Println(emp)
		file, _ := json.MarshalIndent(emp, "", " ")

		_ = ioutil.WriteFile("test.json", file, 0644)
	}

}
