package service

import (
	"errors"
	"math/rand"

	"github.com/sraynitjsr/entity"
	"github.com/sraynitjsr/repository"
)

var repo repository.PostRepository

type PostService interface {
	Validate(post *entity.Post) error
	Create(post *entity.Post) (*entity.Post, error)
	FindAll() ([]entity.Post, error)
}

type service struct{}

func NewPostService(repository repository.PostRepository) PostService {
	repo = repository
	return &service{}
}

func (s *service) Validate(post *entity.Post) error {
	if post == nil {
		return errors.New("empty post")
	}
	if post.Title == "" {
		return errors.New("empty title post")
	}
	if post.Text == "" {
		return errors.New("empty text post")
	}
	return nil
}

func (s *service) Create(newPost *entity.Post) (*entity.Post, error) {
	newPost.ID = rand.Int63()
	return repo.Save(newPost)
}

func (s *service) FindAll() ([]entity.Post, error) {
	return repo.FindAll()
}
