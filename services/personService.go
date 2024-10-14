package services

import (
	"Mereb-V2/constants"
	"Mereb-V2/database"
	"Mereb-V2/models"
	"context"
	"errors"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var PersonCollection = database.OpenCollection(database.Client, "person")

func CreatePersonService(person models.Person) (interface{}, error) {
	ctx, cancel := context.WithTimeout(context.Background(), constants.TIME_OUT)
	defer cancel()

	person.PersonID = primitive.NewObjectID().Hex()
	person.CreatedAt = time.Now()
	person.UpdatedAt = time.Now()

	personCreate, err := PersonCollection.InsertOne(ctx, person)
	return personCreate.InsertedID, err
}

func GetAllPersonsService(perPage int, limit int) ([]models.Person, error) {
	ctx, cancel := context.WithTimeout(context.Background(), constants.TIME_OUT)
	defer cancel()

	fmt.Println(limit, perPage)

	opt := options.Find().SetLimit(int64(perPage)).SetSkip(int64(limit))

	filter := bson.M{}

	cursor, err := PersonCollection.Find(ctx, filter, opt)
	if err != nil {
		// fmt.Println("Two", err)
		return nil, err
	}

	var persons []models.Person
	if err = cursor.All(ctx, &persons); err != nil {
		// fmt.Println("One", err)
		return nil, err
	}
	return persons, nil
}

func GetPersonService(id string) (models.Person, error) {
	ctx, cancel := context.WithTimeout(context.Background(), constants.TIME_OUT)
	defer cancel()

	filter := bson.M{"person_id": id}
	// fmt.Println("PersonID : ", id)

	var person models.Person
	err := PersonCollection.FindOne(ctx, filter).Decode(&person)
	fmt.Println(err)
	if err == mongo.ErrNoDocuments {
		return person, errors.New(constants.PERSON_DOES_NOT_EXIST)
	}
	return person, err
}

func UpdatePersonService(id string, updatedPerson models.Person) (*models.Person, error) {
	ctx, cancel := context.WithTimeout(context.Background(), constants.TIME_OUT)
	defer cancel()

	update := bson.M{"$set": bson.M{

		"name":       updatedPerson.Name,
		"age":        updatedPerson.Age,
		"hobbies":    updatedPerson.Hobbies,
		"updated_at": time.Now(),
	}}

	filter := bson.M{"person_id": id}
	fmt.Println("Person Id : ", id)

	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)

	var newPerson models.Person

	if updatedError := PersonCollection.FindOneAndUpdate(ctx, filter, update, opts).Decode(&newPerson); updatedError != nil {
		return nil, updatedError
	}

	return &newPerson, nil
}

func DeletePersonService(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), constants.TIME_OUT)
	defer cancel()

	_, err := PersonCollection.DeleteOne(ctx, bson.M{"person_id": id})
	return err
}
