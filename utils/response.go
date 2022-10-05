package utils

type Response struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func ResponseSuccess(message string, data any) Response {
	return Response{
		Status:  true,
		Message: message,
		Data:    data,
	}
}

func ResponseFail(message string) Response {
	return Response{
		Status:  true,
		Message: message,
		Data:    nil,
	}
}
