package main

import (
	"log"

	"github.com/sraynitjsr/db"
	"github.com/sraynitjsr/model"
)

func main() {
	// Initialize the database connection
	database := &db.MongoDB{
		URI:      "mongodb://localhost:27017",
		Database: "mydatabase",
	}
	err := database.Connect()
	if err != nil {
		log.Fatal("Error connecting to the database: ", err)
	}
	defer database.Disconnect()

	// Create a new student
	newStudent := &model.Student{
		Name: "Subhradeep Ray",
		Roll: 2020,
	}
	studentModel := model.NewStudentModel(database)
	err = studentModel.Create(newStudent, "students")
	if err != nil {
		log.Fatal("Error in creating student")
	}

	// Get a student by roll number
	roll := 2020
	result, err := studentModel.GetByRoll(roll, "students")
	if err != nil {
		log.Fatal("Error getting student: ", err)
	}
	log.Printf("Student with roll number %d: %+v\n", roll, result)
}
