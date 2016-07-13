package main

import (
	"encoding/json"
	"io"
	"net/http"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var session *mgo.Session = DBConnect()

type Employee struct {
	Id 			bson.ObjectId `bson:"_id,omitempty" json:"id"`
	Firstname 	string	`json:"firstname"`
	Lastname  	string	`json:"lastname"`
	Age       	int		`json:"age"`
}

func DBConnect() *mgo.Session {
	session, err := mgo.Dial("mongodb://admin:password@ds027809.mlab.com:27809/go-connect")

	if err != nil {
		var w http.ResponseWriter
		http.Error(w, "Problems Initializing the database", http.StatusInternalServerError)
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

	var e []Employee
	c := session.DB("go-connect").C("Employee")
	query := c.Find(bson.M{}).All(&e)
	defer session.DB("go-connect").C("Employee").Find(query)
	format, _ := json.Marshal(e)
	io.WriteString(w, string(format))
}

// Get a single employee by firstname from the database
func GetEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Content-Type", "application/json")

	var e Employee
	params := r.URL.Query()

	c := session.DB("go-connect").C("Employee")
	query := c.Find(bson.M{"firstname": params.Get("Firstname")}).One(&e)
	defer session.DB("go-connect").C("Employee").Find(query)
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
		http.Error(w, "Problems Decoding the information", http.StatusInternalServerError)
	}

	if e.Age < 18 {
		http.Error(w, "The person you entered should be older than 18 years old", http.StatusBadRequest)
		return
	}

	errInsert := session.DB("go-connect").C("Employee").Insert(e)
	if errInsert != nil {
		http.Error(w, "Problem inserting data to the database", http.StatusInternalServerError)
	}
}

func UpdateEmployee(w http.ResponseWriter, r *http.Request){
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Content-Type", "application/json")
	var e Employee
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&e)
	if err != nil {
		http.Error(w, "Problems Decoding the information", http.StatusInternalServerError)
	}

	if e.Age < 18 {
		http.Error(w, "The person you entered should be older than 18 years old", http.StatusBadRequest)
		return
	}

	errUpdate := session.DB("go-connect").C("Employee").Update(bson.M{"_id": e.Id}, e)
	if errUpdate != nil {
		http.Error(w, "Problem updating data to the database", http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/", hello)
	http.HandleFunc("/postEmployee", PostEmployee)
	http.HandleFunc("/getEmployees", GetEmployees)
	http.HandleFunc("/getEmployee", GetEmployee)
	http.HandleFunc("/updateEmployee", UpdateEmployee)
	http.ListenAndServe(":8000", nil)
}
