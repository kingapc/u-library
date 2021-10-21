package utils

func Success() string {
	return "U-LIB - Operation was successful."
}

func ErrorX(code int) string {

	switch code {
	case 400:
		return "ULIB-000 - Bad request check body/params"

	case 401:
		return "ULIB-001 - Authentication required"

	case 402:
		return "ULIB-002 - Payment required"

	case 403:
		return "ULIB-003 - Forbidden Action"

	case 404:
		return "ULIB-004 - Not found"

	case 409:
		return "ULIB-009 - Resource is already in use"

	case 422:
		return "ULIB-022 - Missing Access Header"
	}

	return "ULIB-999 - Server Error"
}
