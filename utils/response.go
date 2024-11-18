package utils

type Response struct {
	Status bool `json:"status"`
	Message string `json:"message"`
	Error any `json:"error,omitempty"`
	Data any `json:"data,omitempty"`
	Meta any `json:"meta,omitempty"`
}

func BuildSuccessResponse(message string, data any) Response {
	res := Response{
		Status: true,
		Message: message,
		Data: data,
	}
	return res
}

func BuildFailedResponse(message string, error any, data any) Response {
	res := Response {
		Status: false,
		Message: message,
		Error: error,
		Data: data,
	}
	return res
}