package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
	Sno  string `json:"xxx"`
}

func convertStructTagToJson() {
	student := Student{
		ID:   "9527",
		Name: "test",
		Age:  20,
		Sno:  "so9527",
	}

	jsonStr, err := json.Marshal(student)
	if err == nil {
		i := string(jsonStr)
		fmt.Printf("i: %v\n", i)
	} else {
		fmt.Printf("err: %v\n", err)
	}
}

func convertJsonToStructTag() {
	i := `{"id":"9527","name":"test","age":20,"xxx":"so9527"}`
	var student Student
	err := json.Unmarshal([]byte(i), &student)
	if err == nil {
		fmt.Printf("student: %#v\n", student)
	} else {
		fmt.Printf("err: %#v\n", err)
	}
}

func main() {
	convertStructTagToJson()
	convertJsonToStructTag()
}
