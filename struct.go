// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
)

type Employee struct {
	_id                 string
	Firstname, Lastname string
	Email               string
	Age                 int
}

func (e Employee) Fullname() string {
	return e.Firstname + " " + e.Lastname
}
func Fullname(e Employee) string {
	return e.Firstname + " " + e.Lastname
}

func (e *Employee) ChangeName(firstname string, lastname string) {
	e.Firstname = firstname
	e.Lastname = lastname
}
func ChangeName(e *Employee, firstname string, lastname string) {
	e.Firstname = firstname
	e.Lastname = lastname
}

func main() {
	var employee Employee = Employee{
		_id:       "577bcde9da1918bcab54613f",
		Age:       23,
		Firstname: "Lesa",
		Lastname:  "Barlow",
		Email:     "lesabarlow@bluplanet.com",
	}
	fmt.Println(employee)
	fmt.Println(employee.Fullname())
	fmt.Println(Fullname(employee))
	employee.ChangeName("firstname", "lastname")
	fmt.Println(employee.Fullname())
	ChangeName(&employee, "Carlos", "Danna")
	fmt.Println(employee.Fullname())

	// http.ListenAndServe(":3000", nil)

	var employees []Employee = []Employee{{
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
	for i := 0; i < 5; i++ {
		format, _ := json.Marshal(employees[i])
		fmt.Println(string(format))
	}
	xmlTransform, _ := xml.Marshal(employees)

	jsonTransform, _ := json.Marshal(employees)

	fmt.Println(string(jsonTransform))
	fmt.Println(string(xmlTransform))
}
