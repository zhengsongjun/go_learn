package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	Name string
	Age  int
	SNo  int
}

func convertStudentToJson() {
	student := Student{
		Name: "test",
		Age:  20,
		SNo:  1203,
	}
	value, error := json.Marshal(student)
	jsonStr := string(value)
	if error == nil {
		fmt.Printf("%#v", jsonStr)
	}
}

func convertJsonToStudent() {
	str := `{"Name":"test","Age":20,"SNo":1203}`
	var student Student
	err := json.Unmarshal([]byte(str), &student)
	if err == nil {
		fmt.Printf("%#v", student)
	} else {
		fmt.Printf("%v", err)
	}
}

func main() {
	convertStudentToJson()
	fmt.Printf("\"________________\": %v\n", "________________")
	convertJsonToStudent()
}
