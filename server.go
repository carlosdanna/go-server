package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
)

var employees []Employee = createTestEmployee()

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
