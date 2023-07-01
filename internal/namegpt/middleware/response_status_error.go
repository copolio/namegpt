package middleware

type HttpStatus int

type ResponseStatusError struct {
	Code     HttpStatus `json:"code"`
	Message  string     `json:"message"`
	MetaData string     `json:"metaData"`
}
