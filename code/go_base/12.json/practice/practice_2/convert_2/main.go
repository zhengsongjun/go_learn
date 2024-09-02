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

func strctToJson() {
	class := Class{
		SmallClass: SmallClass{
			Title: "第三荣誉班级",
		},
		Name:    "10-2班",
		Student: make([]Student, 0, 20),
		Teacher: []string{"周", "吴", "郑", "王"},
	}

	for i := 0; i < 20; i++ {
		student := Student{
			Name: fmt.Sprintf("学生%v", i+1),
			Age:  20,
			Sno:  fmt.Sprintf("学号100%v", i+1),
		}

		class.Student = append(class.Student, student)
	}

	if jsonStr, err := json.Marshal(class); err == nil {
		str := string(jsonStr)
		fmt.Printf("str: %v\n", str)
	} else {
		fmt.Printf("err: %v\n", err)
	}
}

func jsonToStruct() {
	jsonStr := `{"Title":"第三荣誉班级","Name":"10-2班","Student":[{"Name":"学生1","Age":20,"Sno":"学号1001"},{"Name":"学生2","Age":20,"Sno":"学号1002"},{"Name":"学生3","Age":20,"Sno":"学号1003"},{"Name":"学生4","Age":20,"Sno":"学号1004"},{"Name":"学生5","Age":20,"Sno":"学号1005"},{"Name":"学生6","Age":20,"Sno":"学号1006"},{"Name":"学生7","Age":20,"Sno":"学号1007"},{"Name":"学生8","Age":20,"Sno":"学号1008"},{"Name":"学生9","Age":20,"Sno":"学号1009"},{"Name":"学生10","Age":20,"Sno":"学号10010"},{"Name":"学生11","Age":20,"Sno":"学号10011"},{"Name":"学生12","Age":20,"Sno":"学号10012"},{"Name":"学生13","Age":20,"Sno":"学号10013"},{"Name":"学生14","Age":20,"Sno":"学号10014"},{"Name":"学生15","Age":20,"Sno":"学号10015"},{"Name":"学生16","Age":20,"Sno":"学号10016"},{"Name":"学生17","Age":20,"Sno":"学号10017"},{"Name":"学生18","Age":20,"Sno":"学号10018"},{"Name":"学生19","Age":20,"Sno":"学号10019"},{"Name":"学生20","Age":20,"Sno":"学号10020"}],"Teacher":["周","吴","郑","王"]}`
	class := &Class{}
	if err := json.Unmarshal([]byte(jsonStr), class); err == nil {
		fmt.Printf("class: %#v\n", class)
	} else {
		fmt.Printf("err: %v\n", err)
	}
}

func main() {
	strctToJson()
	fmt.Printf("\"__________________________\": %v\n", "__________________________")
	jsonToStruct()
}
