package errors

import (
	"reflect"
	"unicode/utf8"
)

func trimFirstRune(s string) string {
	_, i := utf8.DecodeRuneInString(s)
	return s[i:]
}

func IsErrorOfType(err error, i interface{}) bool {
	errorType := trimFirstRune(reflect.TypeOf(err).String())
	expectedType := reflect.TypeOf(i).String()

	return errorType == expectedType
}
