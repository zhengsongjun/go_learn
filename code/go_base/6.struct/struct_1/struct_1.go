package main

import "time"

type Employee struct {
	ID        int
	Name      string
	Address   string
	DoB       time.Time
	Position  string
	Salary    int
	ManagerID int
}

func useStruct() {
	var kangjie Employee
	kangjie.ID = 1111
	kangjie.Name = "wang"
	kangjie.Address = "南阳"
}

func main() {
	useStruct()
}
