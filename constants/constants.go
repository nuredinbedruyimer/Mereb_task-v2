package constants

import "time"

var (
	TIME_OUT               = time.Second * 25
	PERSON_DOES_NOT_EXIST  = "Person Does Not Exist !!!!"
	DATA_BINDING_ERROR     = "Invalid input data, Binding Faild."
	INTERNAL_SERVER_ERROR  = "Internal Server Error. Please Try Again Later."
	RESOURCE_ERROR         = "Conflict: Person Already Exists."
	VALIDATION_FAILED      = "Input Validation Error, Check Required Fields."
	FAILED_TO_GET_PERSON   = "Failed To Get Persons"
	PERSON_CREATED         = "Person Created Successfully !!"
	PERSON_FETCHED_SUCCESS = "Person Fetched Successfully !!"
	PERSON_UPDATED_SUCCESS = "Person Updated Successfully !!"
	PERSON_DELETED_SUCCESS = "Person Deleted Successfully"
	NOT_FOUND              = "Not Found"
)
