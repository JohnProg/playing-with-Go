package models

// error response contains everything we need to use http.Error
type HandlerError struct {
	Error   error
	Message string
	Code    int
}

// book model
type Book struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	Id     int    `json:"id"`
}
