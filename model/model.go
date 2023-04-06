package model

import "fmt"

type Student struct {
	Name string
	Roll int
}

type StudentRepository interface {
	Add(student *Student) error
	Get(roll int) (*Student, error)
}

type StudentService struct {
	repo StudentRepository
}

func NewStudentService(repo StudentRepository) *StudentService {
	return &StudentService{repo: repo}
}

func (s *StudentService) Add(student *Student) error {
	return s.repo.Add(student)
}

func (s *StudentService) Get(roll int) (*Student, error) {
	return s.repo.Get(roll)
}

type InMemoryStudentRepository struct {
	students map[int]*Student
}

func NewInMemoryStudentRepository() *InMemoryStudentRepository {
	return &InMemoryStudentRepository{students: make(map[int]*Student)}
}

func (r *InMemoryStudentRepository) Add(student *Student) error {
	r.students[student.Roll] = student
	return nil
}

func (r *InMemoryStudentRepository) Get(roll int) (*Student, error) {
	student, ok := r.students[roll]
	if !ok {
		return nil, fmt.Errorf("student not found")
	}
	return student, nil
}
