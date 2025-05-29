package utils

func HttpMethodCheck(method string) bool {
	httpMethods := []string{
		"GET", "POST", "PATCH", "DELETE",
	}

	valid := false

	for _, m := range httpMethods {
		if m == method {
			valid = true
			break
		}
	}

	return valid
}
