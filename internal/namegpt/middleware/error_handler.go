package middleware

type HttpStatus int

type HttpError struct {
	Code    HttpStatus `json:"code"`
	Message string     `json:"message"`
}

func (e HttpError) Error() string {
	return e.Message
}
