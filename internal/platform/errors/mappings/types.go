package mappings

type ErrorDetails struct {
	InternalCode string
	StatusCode   int
	Message      string
}

var InternalServerError = ErrorDetails{
	InternalCode: "common:internal-error",
	StatusCode:   500,
	Message:      "internal server error",
}

var UnauthorizedError = ErrorDetails{
	InternalCode: "common:unauthorized",
	StatusCode:   401,
	Message:      "unauthorized",
}

var ForbiddenError = ErrorDetails{
	InternalCode: "common:forbidden",
	StatusCode:   403,
	Message:      "forbidden",
}
