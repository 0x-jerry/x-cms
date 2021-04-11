package controller

import "github.com/kataras/iris/v12"

type Response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type Any = interface{}

var SuccessResponse = Response{
	Status: iris.StatusOK,
}

func BadRequest(err error) Response {
	return Response{
		Status:  iris.StatusBadRequest,
		Message: err.Error(),
	}
}

// 如果 err != nil 就返回 bad request，否则返回 res
func ResponseWithError(res interface{}, err error) Any {
	if err != nil {
		return BadRequest(err)
	}

	return res
}
