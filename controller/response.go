package controller

import "github.com/kataras/iris/v12"

type Any = interface{}

type Response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    Any
}

var SuccessResponse = Response{
	Status: iris.StatusOK,
}

func BadRequest(err error) Response {
	return Response{
		Status:  iris.StatusBadRequest,
		Message: err.Error(),
		Data:    nil,
	}
}

func InternalError(err error) Response {
	return Response{
		Status:  iris.StatusInternalServerError,
		Message: err.Error(),
	}
}

// 如果 err != nil 就返回 bad request，否则返回 res
func ResponseWithError(res interface{}, err error) Response {
	if err != nil {
		return BadRequest(err)
	}

	return Response{
		Status:  iris.StatusOK,
		Message: "",
		Data:    res,
	}
}
