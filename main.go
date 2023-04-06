package main

import (
	"fmt"

	"github.com/sraynitjsr/model"
)

func main() {
	fmt.Println("Unit Tesing in GoLang Using Ginkgo and Gomega")

	repo := model.NewInMemoryStudentRepository()

	service := model.NewStudentService(repo)

	student := &model.Student{
		Name: "Subhradeep Ray",
		Roll: 2020,
	}

	err := service.Add(student)

	if err != nil {
		fmt.Printf("Error adding student: %s", err.Error())
		return
	}

	addedStudent, err := service.Get(student.Roll)

	if err != nil {
		fmt.Printf("Error retrieving student: %s", err.Error())
		return
	}

	fmt.Printf("Added student: name=%s, roll=%d\n", addedStudent.Name, addedStudent.Roll)
}
