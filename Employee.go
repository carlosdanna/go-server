package main

type Employee struct {
	_id                 string
	Firstname, Lastname string
	Email               string
	Age                 int
}

func createTestEmployee() []Employee {
	employees := []Employee{
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
	return employees
}
