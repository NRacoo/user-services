package error

import "errors"

var (
	ErrUserNotFound        = errors.New("user not found")
	ErrPasswordIncorrect   = errors.New("password incorrect")
	ErrUsernameExist       = errors.New("username is already used")
	ErrEmailExist          = errors.New("email is already used")
	ErrPasswordDoesntMatch = errors.New("password doesn't match")
)

var userError = []error{
	ErrUserNotFound,
	ErrPasswordIncorrect,
	ErrUsernameExist,
	ErrPasswordDoesntMatch,
}
