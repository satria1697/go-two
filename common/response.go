package common

type Response struct {
	Data    any    `json:"data"`
	Message string `json:"message"`
	Error   string `json:"error"`
}

func SuccessResponse(data any) Response {
	return Response{
		Data:    data,
		Message: "success",
		Error:   "",
	}
}

func ErrorResponse(error error) Response {
	return Response{
		Data:    nil,
		Message: "failed",
		Error:   error.Error(),
	}
}
