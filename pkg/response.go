package pkg

type BaseSuccessResponse struct {
	Status string `json:"status"`
	Message string `json:"message"`
	Data any `json:"data"`
}

type BaseErrorResponse struct {
	Status string `json:"status"`
	Message string `json:"message"`
}

func NewBaseSuccessResponse(message string, data any) BaseSuccessResponse {
	return BaseSuccessResponse{
		Status: "success",
		Message: message,
		Data: data,
	}
}

func NewBaseErrorResponse(message string) BaseErrorResponse {
	return BaseErrorResponse{
		Status: "error",
		Message: message,
	}
}