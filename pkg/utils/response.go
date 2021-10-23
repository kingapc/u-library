package utils

func Success(message string) map[string]interface{} {

	m := make(map[string]interface{})

	m["status"] = "success"
	m["message"] = message

	return m
}

func ErrorX(httpDescription string, showDetail bool, detail string, isWarning bool) map[string]interface{} {

	m := make(map[string]interface{})

	m["status"] = "Fail"
	if isWarning {
		m["status"] = "Warning"
	}

	m["message"] = httpDescription

	if showDetail {
		m["detail"] = detail
	}

	return m
}
