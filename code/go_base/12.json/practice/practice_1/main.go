package main

import (
	"encoding/json"
	"fmt"
)

type SmallClass struct {
	Name string
}

type Student struct {
	Name string
	Age  int
	Sno  string
}

type Class struct {
	SmallClass
	Teacher []string
	Student []Student
}

func structToJson() {
	class1 := Class{
		SmallClass: SmallClass{
			Name: "3-1",
		},
		Teacher: []string{"wang", "zheng", "hu"},
		Student: make([]Student, 0, 10),
	}

	for i := 0; i < 10; i++ {
		student := Student{
			Name: fmt.Sprintf("学生名%v", i+1),
			Age:  20,
			Sno:  fmt.Sprintf("%v", i+1),
		}
		class1.Student = append(class1.Student, student)
	}

	jsonStr, err := json.Marshal(class1)
	if err == nil {
		str := string(jsonStr)
		fmt.Printf("str: %v\n", str)
	} else {
		fmt.Printf("err: %v\n", err)
	}
}

func JsonToStrcut() {
	str := `{"Name":"3-1","Teacher":["wang","zheng","hu"],"Student":[{"Name":"学生名1","Age":20,"Sno":"1"},{"Name":"学生名2","Age":20,"Sno":"2"},{"Name":"学生名3","Age":20,"Sno":"3"},{"Name":"学生名4","Age":20,"Sno":"4"},{"Name":"学生名5","Age":20,"Sno":"5"},{"Name":"学生名6","Age":20,"Sno":"6"},{"Name":"学生名7","Age":20,"Sno":"7"},{"Name":"学生名8","Age":20,"Sno":"8"},{"Name":"学生名9","Age":20,"Sno":"9"},{"Name":"学生名10","Age":20,"Sno":"10"}]}`
	class := &Class{}
	err := json.Unmarshal([]byte(str), class)
	if err == nil {
		fmt.Printf("class: %#v\n", class)
	} else {
		fmt.Printf("err: %v\n", err)
	}

}

func main() {
	structToJson()
	fmt.Printf("\"______________\": %v\n", "______________")
	JsonToStrcut()
}
