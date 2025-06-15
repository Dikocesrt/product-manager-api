package pkg

type BaseSuccessResponse struct {
	Status string `json:"status"`
	Message string `json:"message"`
	Data any `json:"data"`
}

type BaseErrorResponse struct {
	Status string `json:"status"`
	Message string `json:"message"`
	Errors []ValidationErrorResponse `json:"errors,omitempty"`
}

type ValidationErrorResponse struct {
	Field   string      `json:"field"`
	Message string      `json:"message"`
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

func NewValidationErrorResponse(message string, validationErrors ValidationErrors) BaseErrorResponse {
	var errors []ValidationErrorResponse
	
	for _, err := range validationErrors {
		errors = append(errors, ValidationErrorResponse{
			Field:   err.Field,
			Message: err.Message,
		})
	}
	
	return BaseErrorResponse{
		Status:  "error",
		Message: message,
		Errors:  errors,
	}
}