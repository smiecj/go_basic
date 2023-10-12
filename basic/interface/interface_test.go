package interface_

import (
	"io"
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
)

var (
	e *MyError
	w io.Writer
)

type MyError struct{}

func (e *MyError) Error() string {
	return ""
}

func GetMyError() error {
	return e
}

func TestNilInterface(t *testing.T) {
	println(GetMyError().Error())
	require.True(t, GetMyError() == (*MyError)(nil)) // true
	// require.True(t, GetMyError() == nil) // will return false
	require.True(t, w == nil)

	e := GetMyError()
	errVal := reflect.ValueOf(e)
	if errVal.Kind() == reflect.Pointer {
		require.True(t, errVal.IsNil())
	}
}
