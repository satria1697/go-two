package common

type Response struct {
	Data    any       `json:"data"`
	Message string    `json:"message"`
	Error   *[]string `json:"error"`
}

func SuccessResponse(data any) Response {
	return Response{
		Data:    data,
		Message: "success",
		Error:   nil,
	}
}

func ErrorResponse(errors []error) Response {
	return Response{
		Data:    nil,
		Message: "failed",
		Error:   handleError(errors),
	}
}

func handleError(errors []error) *[]string {
	var errs []string
	for _, err := range errors {
		errs = append(errs, err.Error())
	}
	return &errs
}
