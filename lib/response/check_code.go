package response

func IsApiFailed(code int) bool {
	return clientError(code) || serverError(code)
}

func IsApiClientError(code int) bool {
	return clientError(code)
}

func IsApiServerError(code int) bool {
	return serverError(code)
}

func IsApiSuccess(code int) bool {
	return code >= 200 && code < 300
}

func clientError(code int) bool {
	return code >= 400 && code < 500
}
func serverError(code int) bool {
	return code >= 500
}
