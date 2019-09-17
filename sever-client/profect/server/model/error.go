package model
//Customize errors is based on business logic
import (
	"errors"
)

var (
	ERROR_USER_NOTEXISTS = errors.New("user dose not exsist")
	ERROR_USER_EXISTS = errors.New("user already exsist")
	ERROR_USER_PWD = errors.New("user's PWD is error ")
)