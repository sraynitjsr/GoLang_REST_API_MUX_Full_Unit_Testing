package main

import (
	"fmt"
	"testing"

	. "github.com/sraynitjsr/model"
)

type mockStudentRepository struct {
	students map[int]*Student
}

func (r *mockStudentRepository) Add(student *Student) error {
	r.students[student.Roll] = student
	return nil
}

func (r *mockStudentRepository) Get(roll int) (*Student, error) {
	student, ok := r.students[roll]
	if !ok {
		return nil, fmt.Errorf("student not found")
	}
	return student, nil
}

func TestAddStudent(t *testing.T) {
	// Create a new mock student repository
	repo := &mockStudentRepository{students: make(map[int]*Student)}

	// Create a new student service with the mock student repository
	service := NewStudentService(repo)

	// Create a new student
	student := &Student{
		Name: "John Doe",
		Roll: 123,
	}

	// Add the student using the student service
	err := service.Add(student)
	if err != nil {
		t.Errorf("Error adding student: %s", err.Error())
	}

	// Retrieve the added student from the mock repository
	addedStudent, err := repo.Get(student.Roll)
	if err != nil {
		t.Errorf("Error retrieving student: %s", err.Error())
	}

	// Check if the retrieved student is the same as the added student
	if addedStudent != student {
		t.Errorf("Added student and retrieved student do not match")
	}
}
