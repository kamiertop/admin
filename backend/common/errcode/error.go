package errcode

import "github.com/gofiber/fiber/v3"

type AppError struct {
	Err  error
	Msg  string
	Code int
}

func (e *AppError) Error() string {
	return e.Err.Error()
}

func NewError(code int, msg string, err error) *AppError {
	return &AppError{
		Err:  err,
		Msg:  msg,
		Code: code,
	}
}

func NewInternalError(err error) *AppError {
	return &AppError{
		Err:  err,
		Msg:  "服务器开小差啦",
		Code: fiber.StatusInternalServerError,
	}
}

func NewBadRequestError(err error) *AppError {
	return &AppError{
		Err:  err,
		Msg:  "参数有误",
		Code: fiber.StatusBadRequest,
	}
}
