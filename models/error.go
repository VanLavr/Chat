package models

import "errors"

var (
	// ErrInternalServerError will throw if any the Internal Server Error happen
	ErrInternalServerError = errors.New("internal Server Error")
	// ErrNotFound will throw if the requested item is not exists
	ErrNotFound = errors.New("your requested Item is not found")
	// ErrConflict will throw if the current action already exists
	ErrConflict = errors.New("your Item already exist")
	// ErrBadParamInput will throw if the given request-body or params is not valid
	ErrBadParamInput = errors.New("given Param is not valid")
	// ErrUnauthorized will throw if the user is unauthorized
	ErrUnauthorized = errors.New("unauthorized")
	// ErrUserAlreadyInChat will throw if the user already entered this chatroom
	ErrUserAlreadyInChat = errors.New("user already entered this chatroom")
	// ErrEmptyFields will throw when trying ti create a user without name or password
	ErrEmptyFields = errors.New("can not create user without name or password")
	// ErrAlreadyExists will throw if entity with such params already exists
	ErrAlreadyExists = errors.New("such entity already exists")
)
