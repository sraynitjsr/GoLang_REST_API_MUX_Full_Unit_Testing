package service

import (
	"testing"

	"github.com/sraynitjsr/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockRepository struct {
	mock.Mock
}

func (mockRepo *MockRepository) Save(post *entity.Post) (*entity.Post, error) {
	args := mockRepo.Called()
	result := args.Get(0)
	return result.(*entity.Post), args.Error(1)
}

func (mockRepo *MockRepository) FindAll() ([]entity.Post, error) {
	args := mockRepo.Called()
	result := args.Get(0)
	return result.([]entity.Post), args.Error(1)
}

func TestFindAll(t *testing.T) {
	post := entity.Post{
		ID:    100,
		Title: "Some Title",
		Text:  "Some Text",
	}
	mockRepo := new(MockRepository)

	mockRepo.On("FindAll").Return([]entity.Post{post}, nil)

	testService := NewPostService(mockRepo)

	result, _ := testService.FindAll()

	assert.Equal(t, int64(100), result[0].ID)
	assert.Equal(t, "Some Title", result[0].Title)
	assert.Equal(t, "Some Text", result[0].Text)
}

func TestValidateEmptyPost(t *testing.T) {
	testService := NewPostService(nil)
	err := testService.Validate(nil)
	assert.NotNil(t, err)
	assert.Equal(t, "empty post", err.Error())
}

func TestValidateEmptyPostTitle(t *testing.T) {
	testService := NewPostService(nil)
	post := entity.Post{
		ID:    1,
		Title: "",
		Text:  "Some Text",
	}
	err := testService.Validate(&post)
	assert.NotNil(t, err)
	assert.Equal(t, "empty title post", err.Error())
}

func TestValidateEmptyPostText(t *testing.T) {
	testService := NewPostService(nil)
	post := entity.Post{
		ID:    1,
		Title: "Some Title",
		Text:  "",
	}
	err := testService.Validate(&post)
	assert.NotNil(t, err)
	assert.Equal(t, "empty text post", err.Error())
}
