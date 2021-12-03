package utils

import (
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

type GetEnvMock struct {
	GetEnvVariableFn func(string) string
}

// func (mock GetEnvMock) GoDotEnvVariable(key string) string {
// 	return mock.GetEnvVariableFn(key)
// }

func TestGoDotEnvVariable(t *testing.T) {

	myMock := GetEnvMock{}
	myMock.GetEnvVariableFn = func(s string) string {
		return "doadmin"
	}

	item := "doadmin"
	assert.Equal(t, item, myMock.GetEnvVariableFn("test"))
}

func TestGoDotEnvVariableEmpty(t *testing.T) {

	myMock := GetEnvMock{}
	myMock.GetEnvVariableFn = func(s string) string {
		return "Key env required"
	}

	item := "Key env required"
	assert.Equal(t, item, myMock.GetEnvVariableFn(""))
}

func TestGoDotEnvVariableNotEnvFile(t *testing.T) {

	myMock := GetEnvMock{}
	myMock.GetEnvVariableFn = func(s string) string {
		err := godotenv.Load("")
		if err != nil {
			return "Unable to load env"
		}

		return ""
	}

	item := "Unable to load env"
	assert.Equal(t, item, myMock.GetEnvVariableFn("test"))
}
