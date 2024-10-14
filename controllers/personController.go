package controllers

import (
	"Mereb-V2/constants"
	"Mereb-V2/helpers"
	"Mereb-V2/models"
	"Mereb-V2/services"
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func CreatePersonController(c *gin.Context) {

	//  Bind The Request Body To Go Struct with Person Type
	var person models.Person
	if err := c.ShouldBindJSON(&person); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Message": constants.DATA_BINDING_ERROR,
			"Error":   err.Error()})
		return
	}

	//  Validate Request Body

	personValidator := helpers.NewValidatorService()

	if validationError := personValidator.ValidateData(person); validationError != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Message": constants.VALIDATION_FAILED,
			"Error":   validationError.Error()})
		return
	}

	ctx, cancell := context.WithTimeout(context.Background(), constants.TIME_OUT)
	defer cancell()

	//  Handle No Same Person Created Twice

	filter := bson.M{"name": person.Name}

	personCount, _ := services.PersonCollection.CountDocuments(ctx, filter)

	if personCount > 0 {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Message": constants.RESOURCE_ERROR})
		return

	}

	//  Do The creation Task On The Service Part and return The Created Person else Handle The Case
	insertionID, err := services.CreatePersonService(person)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Message": constants.INTERNAL_SERVER_ERROR,
			"Error":   err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"Message":     constants.PERSON_CREATED,
		"InsertionID": insertionID,
	})
}
func GetPersonsController(c *gin.Context) {
	limitStr := c.Query("limit")
	offsetStr := c.Query("offset")
	fmt.Println(limitStr)
	personPerPage := 3
	personOffset := 0

	if currValue, err := strconv.Atoi(limitStr); err == nil && currValue >= 1 {
		personPerPage = currValue
	}

	if currValue, err := strconv.Atoi(offsetStr); err == nil && currValue >= 1 {
		personOffset = currValue
	}

	// fmt.Println("Error In Here")

	persons, err := services.GetAllPersonsService(personPerPage, personOffset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Message": constants.FAILED_TO_GET_PERSON,
			"Error":   err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"Message": constants.PERSON_FETCHED_SUCCESS,
		"PerPage": personPerPage,
		"Page":    personOffset,
		"Data":    persons,
	})
}
func GetPersonController(c *gin.Context) {
	personID := c.Param("id")

	person, err := services.GetPersonService(personID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"Message": constants.PERSON_DOES_NOT_EXIST,
			"Error":   err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"Message": constants.PERSON_FETCHED_SUCCESS,
		"Person":  person,
	})
}

func UpdatePersonController(c *gin.Context) {
	personID := c.Param("id")
	var updatedPerson models.Person
	if err := c.ShouldBindJSON(&updatedPerson); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Message": constants.DATA_BINDING_ERROR,
			"Error":   err.Error()})
		return
	}
	newPerson, err := services.UpdatePersonService(personID, updatedPerson)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Message": constants.INTERNAL_SERVER_ERROR,
			"Error":   err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Message": constants.PERSON_UPDATED_SUCCESS,
		"Data":    *newPerson,
	})
}

func DeletePersonController(c *gin.Context) {
	id := c.Param("id")
	if err := services.DeletePersonService(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Message": constants.INTERNAL_SERVER_ERROR,
			"Error":   err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Message": constants.PERSON_DELETED_SUCCESS,
	})
}
func NotFoundHandler(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{"Message": constants.NOT_FOUND})
}
