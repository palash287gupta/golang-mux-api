package errors

//ServiceError should be used toreturn business error messages
type ServiceError struct {
	Message string `json:"message"`
}
