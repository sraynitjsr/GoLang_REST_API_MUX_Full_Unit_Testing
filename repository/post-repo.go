package repository

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
	"github.com/sraynitjsr/entity"
)

const (
	projectId  string = "sraynitjsr-2020"
	collection string = "posts"
)

type PostRepository interface {
	Save(post *entity.Post) (*entity.Post, error)
	FindAll() ([]entity.Post, error)
}

type RepositoryImpl struct{}

func NewPostRepository() PostRepository {
	return &RepositoryImpl{}
}

func (repo *RepositoryImpl) Save(post *entity.Post) (*entity.Post, error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectId)
	if err != nil {
		log.Fatalf("Failed to create the firestore client: %v", err)
		return nil, err
	}

	defer client.Close()

	_, _, addError := client.Collection(collection).Add(ctx, map[string]interface{}{
		"ID":    post.ID,
		"Title": post.Title,
		"Text":  post.Text,
	})

	if err != nil {
		log.Fatalf("Failed to add a new post: %v", addError)
		return nil, addError
	}

	return post, nil
}

func (repo *RepositoryImpl) FindAll() ([]entity.Post, error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectId)
	if err != nil {
		log.Fatalf("Failed to create the firestore client: %v", err)
		return nil, err
	}

	defer client.Close()

	var allPosts []entity.Post

	allDocumentsIterator := client.Collection(collection).Documents(ctx)

	for {
		newDoc, docError := allDocumentsIterator.Next()
		if docError != nil {
			err = docError
			break
		}

		currentPost := entity.Post{
			ID:    newDoc.Data()["ID"].(int64),
			Title: newDoc.Data()["Title"].(string),
			Text:  newDoc.Data()["Text"].(string),
		}
		allPosts = append(allPosts, currentPost)
	}
	if len(allPosts) == 0 {
		return nil, err
	} else {
		return allPosts, nil
	}
}
