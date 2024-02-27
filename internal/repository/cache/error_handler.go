package cache

type ErrorHandler struct {
	Err        error `json:"err"`
	StatusCode int   `json:"statusCode"`
}

func NewErrorHandler(err error, statusCode int) *ErrorHandler {
	return &ErrorHandler{Err: err, StatusCode: statusCode}
}
func (e ErrorHandler) Error() string {
	return e.Err.Error()
}
