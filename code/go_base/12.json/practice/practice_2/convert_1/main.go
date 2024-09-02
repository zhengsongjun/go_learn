package main

import (
	"encoding/json"
	"fmt"
)

type SmallClass struct {
	Title string
}

type Student struct {
	Name string
	Age  int
	Sno  string
}

type Class struct {
	SmallClass
	Name    string
	Student []Student
	Teacher []string
}

func strcutToJson() {
	class := Class{
		SmallClass: SmallClass{
			Title: "第三荣誉班级",
		},
		Name:    "10年级3班",
		Student: make([]Student, 0, 20),
		Teacher: []string{"王", "周", "郑", "吴"},
	}

	for i := 0; i < 20; i++ {
		student := Student{
			Name: fmt.Sprintf("学生%v", i+1),
			Age:  20,
			Sno:  fmt.Sprintf("学号%v", 100+i+1),
		}
		class.Student = append(class.Student, student)
	}

	if classJson, err := json.Marshal(class); err == nil {
		str := string(classJson)
		fmt.Printf("str: %v\n", str)
	}
}

func jsonToStruct() {
	str := `{"Title":"第三荣誉班级","Name":"10年级3班","Student":[{"Name":"学生1","Age":20,"Sno":"学号101"},{"Name":"学生2","Age":20,"Sno":"学号102"},{"Name":"学生3","Age":20,"Sno":"学号103"},{"Name":"学生4","Age":20,"Sno":"学号104"},{"Name":"学生5","Age":20,"Sno":"学号105"},{"Name":"学生6","Age":20,"Sno":"学号106"},{"Name":"学生7","Age":20,"Sno":"学号107"},{"Name":"学生8","Age":20,"Sno":"学号108"},{"Name":"学生9","Age":20,"Sno":"学号109"},{"Name":"学生10","Age":20,"Sno":"学号110"},{"Name":"学生11","Age":20,"Sno":"学号111"},{"Name":"学生12","Age":20,"Sno":"学号112"},{"Name":"学生13","Age":20,"Sno":"学号113"},{"Name":"学生14","Age":20,"Sno":"学号114"},{"Name":"学生15","Age":20,"Sno":"学号115"},{"Name":"学生16","Age":20,"Sno":"学号116"},{"Name":"学生17","Age":20,"Sno":"学号117"},{"Name":"学生18","Age":20,"Sno":"学号118"},{"Name":"学生19","Age":20,"Sno":"学号119"},{"Name":"学生20","Age":20,"Sno":"学号120"}],"Teacher":["王","周","郑","吴"]}`
	class := &Class{}
	if err := json.Unmarshal([]byte(str), class); err == nil {
		fmt.Printf("class: %#v\n", class)
	} else {
		fmt.Printf("err: %v\n", err)
	}

}

func main() {
	strcutToJson()
	fmt.Printf("\"__________________\": %v\n", "__________________")
	jsonToStruct()
}
