package common

import (
	"net/http"
)

type AppError struct {
	StatusCode  int    `json:"status_code"`
	RootErr     error  `json:"-"`
	Message     string `json:"message"`
	Log	  		string `json:"log"`
	Key	  		string `json:"error_key"`
}

func NewFullErrorResponse(statusCode int, rootErr error, message, log, key string) *AppError {
	return &AppError{
		StatusCode: statusCode,
		RootErr: rootErr,
		Message: message,
		Log: log,
		Key: key,
	}
}

func NewErrorResponse(root error, msg, log, key string) *AppError {
	return &AppError{
		StatusCode: http.StatusBadRequest,
		RootErr: root,
		Message: msg,
		Log: log,
		Key: key,
	}
}

func (e *AppError) RootError() error {
	if err, ok := e.RootErr.(*AppError); ok{
		return err.RootError()
	}
	return e.RootErr
}

func (e *AppError) Error() string {
	return e.RootError().Error()
}



func NewCustomErr(root error, msg, key string) *AppError {
	if root != nil {
		return NewErrorResponse(root, msg, root.Error(), key)
	}
	return NewErrorResponse(nil, msg, msg, key)
}

func ErrDB(err error) *AppError {
	return NewFullErrorResponse(http.StatusInternalServerError, err, "something went wrong with the database", err.Error(), "DB_ERROR")
}

func ErrBadRequest(msg string) *AppError {
	return NewCustomErr(nil, msg, "ErrBadRequest")
}

func ErrInvalidRequest(err error) *AppError {
	return NewErrorResponse(err, "invalid request", err.Error(), "ErrInvalidRequest")
}

func ErrInternalServer(err error) *AppError {
	return NewFullErrorResponse(http.StatusInternalServerError, err, "internal server error", err.Error(), "ErrInternalServer")
}

func ErrCannotListEntity(entity string, err error) *AppError {
	return NewCustomErr(
		err, 
		"Cannot list "+entity, 
		"ErrCannotListEntity",
	)
}

func ErrCannotCreateEntity(entity string, err error) *AppError {
	return NewCustomErr(
		err, 
		"Cannot create "+entity, 
		"ErrCannotCreateEntity",
	)
}

func ErrCannotUpdateEntity(entity string, err error) *AppError {
	return NewCustomErr(
		err,
		"Cannot update "+entity,
		"ErrCannotUpdateEntity",
	)
}

func ErrCannotDeleteEntity(entity string, err error) *AppError {
	return NewCustomErr(
		err,
		"Cannot delete "+entity,
		"ErrCannotDeleteEntity",
	)
}