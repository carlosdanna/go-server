package main

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type MyServer struct {
	mongo *mgo.Session
}

type Employee struct {
	Id        bson.ObjectId `bson:"_id,omitempty" json:"Id"`
	Firstname string        `json:"Firstname"`
	Lastname  string        `json:"Lastname"`
	Username  string        `json:"Username"`
	Password  string        `json:"Password"`
	Age       int           `json:"Age"`
}

type Error struct {
	Code    int    `json: "code"`
	Message string `json: "messsage"`
}

func (s *MyServer) DBConnect() {
	m, err := mgo.Dial("mongodb://admin:password@ds027809.mlab.com:27809/go-connect")

	if err != nil {
		var w http.ResponseWriter
		error := Error{1000, "Problems Initializing the database"}
		format, _ := json.Marshal(error)
		http.Error(w, string(format), http.StatusInternalServerError)
	}
	s.mongo = m
	return
}

//set up the headers for the api calls
//TODO: Check alice plugin for chain handlers
func settingHeaders(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Access-Control-Allow-Origin", "*")
		w.Header().Add("Content-Type", "application/json")
		fn(w, r)
	}
}

func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello world!, Yeah this is my first api call I am learning this stuff")
}

// Get group of employees from a mongoDB
func (s MyServer) GetEmployees(w http.ResponseWriter, r *http.Request) {
	var e []Employee

	c := s.mongo.DB("go-connect").C("Employee")
	err := c.Find(bson.M{}).All(&e)
	if err != nil {
		error := Error{1000, err.Error()}
		format, _ := json.Marshal(error)
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, string(format))
		return
	}
	format, _ := json.Marshal(e)
	io.WriteString(w, string(format))
}

// Get a single employee by firstname from the database
func (s MyServer) GetEmployee(w http.ResponseWriter, r *http.Request) {
	var e Employee
	params := r.URL.Query()

	c := s.mongo.DB("go-connect").C("Employee")
	err := c.Find(bson.M{"firstname": params.Get("Firstname")}).One(&e)
	if err != nil {
		error := Error{1000, err.Error()}
		format, _ := json.Marshal(error)
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, string(format))
		return
	}
	format, _ := json.Marshal(e)
	io.WriteString(w, string(format))
}

// Creates a new employee in the Db the employee has to be older that 18 years old
// Otherwise it will send a bad request response
func (s MyServer) PostEmployee(w http.ResponseWriter, r *http.Request) {
	var e Employee
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&e)
	if err != nil {
		error := Error{1000, "Problems Decoding the information"}
		format, _ := json.Marshal(error)
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, string(format))
		return
	}

	if e.Age < 18 {
		error := Error{1000, "The person you entered should be older than 18 years old"}
		format, _ := json.Marshal(error)
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, string(format))
		return
	}

	errInsert := s.mongo.DB("go-colnnect").C("Employee").Insert(e)
	if errInsert != nil {
		error := Error{1000, "Problem inserting data to the database"}
		format, _ := json.Marshal(error)
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, string(format))
		return
	}
}

func (s MyServer) UpdateEmployee(w http.ResponseWriter, r *http.Request) {
	var e Employee
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&e)
	if err != nil {
		error := Error{1000, "Problems Decoding the information"}
		format, _ := json.Marshal(error)
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, string(format))
		return
	}

	if e.Age < 18 {
		error := Error{1000, "The person you entered should be older than 18 years old"}
		format, _ := json.Marshal(error)
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, string(format))
		return
	}

	errUpdate := s.mongo.DB("go-connect").C("Employee").Update(bson.M{"_id": e.Id}, e)
	if errUpdate != nil {
		error := Error{1000, "Problem updating data to the database"}
		format, _ := json.Marshal(error)
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, string(format))
		return
	}

}

func main() {
	r := mux.NewRouter()

	var s MyServer
	s.DBConnect()

	r.HandleFunc("/", hello)
	r.HandleFunc("/getEmployees", settingHeaders(s.GetEmployees)).Methods("GET")
	r.HandleFunc("/getEmployee", settingHeaders(s.GetEmployee)).Methods("GET")
	r.HandleFunc("/postEmployee", settingHeaders(s.PostEmployee)).Methods("POST")
	r.HandleFunc("/updateEmployee", settingHeaders(s.UpdateEmployee)).Methods("POST")
	http.ListenAndServe(":8000", r)
}
