package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var employees []Employee = createTestEmployee()

type emp struct {
	Firstname string
	Lastname  string
	Age       int
}

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

func GetEmployeesDB(w http.ResponseWriter, r *http.Request) {
	session, err := mgo.Dial("mongodb://admin:password@ds027809.mlab.com:27809/go-connect")
	if err == nil {
		c := session.DB("go-connect")
		fmt.Println(c, "connected")
	} else {
		fmt.Println("Error: ", err)
	}
	var e []emp
	c := session.DB("go-connect").C("Employee")

	query := c.Find(bson.M{}).All(&e)
	session.DB("go-connect").C("Employee").Find(query)
	fmt.Println(e)
}

func PostEmployee(w http.ResponseWriter, r *http.Request) {
	var e emp
	// buff, _ := ioutil.ReadAll(r.Body)
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&e)
	if err != nil {
		panic(err)
	}
	// fmt.Print(string(buff))
	fmt.Println(e)
	format, _ := json.Marshal(e)
	fmt.Println(string(format))
	session, err := mgo.Dial("mongodb://admin:password@ds027809.mlab.com:27809/go-connect")
	if err == nil {
		c := session.DB("go-connect")
		fmt.Println(c, "connected")
	} else {
		fmt.Println("Error: ", err)
	}
	errInsert := session.DB("go-connect").C("Employee").Insert(e)
	if errInsert != nil {
		panic(err)
	}
	session.Close()

}

func main() {

	http.HandleFunc("/", hello)
	http.HandleFunc("/getEmployeesj", GetEmployeesj)
	http.HandleFunc("/getEmployeesx", GetEmployeesx)
	http.HandleFunc("/postEmployee", PostEmployee)
	http.HandleFunc("/getEmployeesDb", GetEmployeesDB)
	http.ListenAndServe(":8000", nil)
}
