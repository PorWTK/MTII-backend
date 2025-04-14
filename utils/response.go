package utils

type Response struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data"`
	Error   any    `json:"error"`
}

type EmptyObj struct{}

func BuildResponseSuccess(message string, data any) Response {
	res := Response{
		Status:  true,
		Message: message,
		Data:    data,
	}
	return res
}

func BuildResponseFailed(message string, error any, data any) Response {
	res := Response{
		Status:  false,
		Message: message,
		Error:   error,
		Data:    data,
	}
	return res
}
