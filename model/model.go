package model

import (
	"context"

	"github.com/sraynitjsr/db"
	"gopkg.in/mgo.v2/bson"
)

type Student struct {
	Name string `bson:"name"`
	Roll int    `bson:"roll"`
}

type StudentModel struct {
	db db.Database
}

func NewStudentModel(database db.Database) *StudentModel {
	return &StudentModel{db: database}
}

func (sm *StudentModel) Create(student *Student, collectionName string) error {
	collection := sm.db.GetCollection(collectionName)

	_, err := collection.InsertOne(context.Background(), bson.M{
		"name": student.Name,
		"roll": student.Roll,
	})
	if err != nil {
		return err
	}

	return nil
}

func (sm *StudentModel) GetByRoll(roll int, collectionName string) (*Student, error) {
	collection := sm.db.GetCollection(collectionName)

	result := &Student{}
	err := collection.FindOne(context.Background(), bson.M{"roll": roll}).Decode(result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
