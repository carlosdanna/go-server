package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Employee struct {
	Firstname string
	Lastname  string
	Age       int
}

func DBConnect() *mgo.Session {
	session, err := mgo.Dial("mongodb://admin:password@ds027809.mlab.com:27809/go-connect")
	if err != nil {
		fmt.Println("Error: ", err)
	}
	return session
}

func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello world!, Yeah this is my first api call I am learning this stuff")
}

// Get group of employees from a mongoDB
func GetEmployees(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Content-Type", "application/json")
	session := DBConnect()

	var e []Employee
	c := session.DB("go-connect").C("Employee")
	query := c.Find(bson.M{}).All(&e)
	defer session.DB("go-connect").C("Employee").Find(query)
	session.Close()
	format, _ := json.Marshal(e)
	io.WriteString(w, string(format))
}

// Get group of employees from a mongoDB
func GetEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Content-Type", "application/json")
	session := DBConnect()

	var e Employee
	params := r.URL.Query()

	c := session.DB("go-connect").C("Employee")
	query := c.Find(bson.M{"firstname": params.Get("Firstname")}).One(&e)
	defer session.DB("go-connect").C("Employee").Find(query)
	session.Close()
	format, _ := json.Marshal(e)
	io.WriteString(w, string(format))
}

//Creates a new employee in the Db the employee has to be older that 18 years old
//Otherwise it will send a bad request response
func PostEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Content-Type", "application/json")
	var e Employee
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&e)
	if err != nil {
		panic(err)
	}

	if e.Age < 18 {
		http.Error(w, "The person you entered should be older than 18 years old", http.StatusBadRequest)
		return
	}

	session := DBConnect()
	errInsert := session.DB("go-connect").C("Employee").Insert(e)
	if errInsert != nil {
		panic(err)
	}
	session.Close()

}

func main() {

	http.HandleFunc("/", hello)
	http.HandleFunc("/postEmployee", PostEmployee)
	http.HandleFunc("/getEmployees", GetEmployees)
	http.HandleFunc("/getEmployee", GetEmployee)
	http.ListenAndServe(":8000", nil)
}
