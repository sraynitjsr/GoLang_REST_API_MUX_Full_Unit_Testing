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

	testService := NewPostService(mockRepo)

	mockRepo.On("FindAll").Return([]entity.Post{post}, nil)

	result, _ := testService.FindAll()

	assert.Equal(t, int64(100), result[0].ID)
	assert.Equal(t, "Some Title", result[0].Title)
	assert.Equal(t, "Some Text", result[0].Text)

	//Mock Assertion => Behavorial Testing
	mockRepo.AssertExpectations(t)
}

func TestCreate(t *testing.T) {
	post := entity.Post{
		ID:    111,
		Title: "Some Dummy Title",
		Text:  "Some Dummy Text",
	}

	mockRepo := new(MockRepository)

	testService := NewPostService(mockRepo)

	mockRepo.On("Save").Return(&post, nil)

	result, err := testService.Create(&post)

	assert.Nil(t, err)
	assert.NotNil(t, result.ID)
	assert.Equal(t, "Some Dummy Title", result.Title)
	assert.Equal(t, "Some Dummy Text", result.Text)

	mockRepo.AssertExpectations(t)
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
