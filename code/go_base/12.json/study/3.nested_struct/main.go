package main

import (
	"encoding/json"
	"fmt"
)

type Class struct {
	Title    string    `json:"title"`
	Students []Student `json:"studentList"`
}

type Student struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	SNo  string `json:"sno"`
}

func nestedStructToJson() {
	class1 := Class{
		Title:    "3-1班",
		Students: make([]Student, 0, 10),
	}
	for i := 0; i < 10; i++ {
		student := Student{
			Name: fmt.Sprintf("学生%v", i),
			Age:  20,
			SNo:  fmt.Sprintf("学号%v", i),
		}
		class1.Students = append(class1.Students, student)
	}

	jsonStr, err := json.Marshal(class1)
	if err == nil {
		str := string(jsonStr)
		fmt.Printf("str: %v\n", str)
	} else {
		fmt.Printf("err: %v\n", err)
	}
}

func JsonToNestedStruct() {
	i := `{"title":"3-1班","studentList":[{"name":"学生0","age":20,"sno":"学号0"},{"name":"学生1","age":20,"sno":"学号1"},{"name":"学生2","age":20,"sno":"学号2"},{"name":"学生3","age":20,"sno":"学号3"},{"name":"学生4","age":20,"sno":"学号4"},{"name":"学生5","age":20,"sno":"学号5"},{"name":"学生6","age":20,"sno":"学号6"},{"name":"学生7","age":20,"sno":"学号7"},{"name":"学生8","age":20,"sno":"学号8"},{"name":"学生9","age":20,"sno":"学号9"}]}`
	class1 := &Class{}
	err := json.Unmarshal([]byte(i), class1)
	if err == nil {
		fmt.Printf("class1: %#v\n", class1)
	} else {
		fmt.Printf("err: %v\n", err)
	}
}

func main() {
	nestedStructToJson()
	fmt.Printf("\"_____________________\": %v\n", "_____________________")
	JsonToNestedStruct()
}
