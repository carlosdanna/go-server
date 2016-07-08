package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
)

type Employee struct {
	_id                 string
	Firstname, Lastname string
	Email               string
	Age                 int
}

var employees []Employee = []Employee{
	{
		_id:       "577bcde9da1918bcab54613f",
		Age:       23,
		Firstname: "Lesa",
		Lastname:  "Barlow",
		Email:     "lesabarlow@bluplanet.com",
	},
	{
		_id:       "577bcde9da1918bdab54613f",
		Age:       30,
		Firstname: "Lesa",
		Lastname:  "Barlow",
		Email:     "lesabarlow@bluplanet.com",
	},
	{
		_id:       "577bcde9da1918beab54613f",
		Age:       29,
		Firstname: "Lesa",
		Lastname:  "Barlow",
		Email:     "lesabarlow@bluplanet.com",
	}, {
		_id:       "577bcde9da1918bfab54613f",
		Age:       28,
		Firstname: "Lesa",
		Lastname:  "Barlow",
		Email:     "lesabarlow@bluplanet.com",
	}, {
		_id:       "577bcde9da1918b0ab54613f",
		Age:       27,
		Firstname: "Lesa",
		Lastname:  "Barlow",
		Email:     "lesabarlow@bluplanet.com",
	}}

func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello world!")
}

func GetEmployeesj(w http.ResponseWriter, r *http.Request) {
	format, _ := json.Marshal(employees)
	fmt.Println("Getting employees on JSON")
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Content-Type", "application/json")
	io.WriteString(w, string(format))
}
func GetEmployeesx(w http.ResponseWriter, r *http.Request) {
	format, _ := xml.Marshal(employees)
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Content-Type", "application/xml")
	fmt.Println("Getting employees on XML")
	io.WriteString(w, string(format))
}

func main() {
	http.HandleFunc("/", hello)
	http.HandleFunc("/getEmployeesj", GetEmployeesj)
	http.HandleFunc("/getEmployeesx", GetEmployeesx)
	http.ListenAndServe(":8000", nil)
}
