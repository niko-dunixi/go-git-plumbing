package utils

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShouldTrace(t *testing.T) {
	strptr := func(s string) *string {
		return &s
	}
	testcases := []struct {
		environmentValue *string
		expectedValue    bool
	}{
		{
			environmentValue: nil,
			expectedValue:    false,
		},
		{
			environmentValue: strptr(""),
			expectedValue:    false,
		},
		{
			environmentValue: strptr("false"),
			expectedValue:    false,
		},
		{
			environmentValue: strptr("true"),
			expectedValue:    true,
		},
		{
			environmentValue: strptr("TRUE"),
			expectedValue:    true,
		},
		{
			environmentValue: strptr("1"),
			expectedValue:    true,
		},
		{
			environmentValue: strptr("2"),
			expectedValue:    true,
		},
		{
			environmentValue: strptr("3"),
			expectedValue:    false,
		},
		{
			environmentValue: strptr("0"),
			expectedValue:    false,
		},
	}

	defer func() {
		os.Unsetenv("GIT_TRACE")
	}()

	for _, currentTestcase := range testcases {
		os.Unsetenv("GIT_TRACE")
		inputValue := "<nil>"
		if environmentValue := currentTestcase.environmentValue; environmentValue != nil {
			inputValue = *environmentValue
			os.Setenv("GIT_TRACE", *environmentValue)
		}
		actualTraceValue := ShouldTrace()
		assert.Equal(t, actualTraceValue, currentTestcase.expectedValue,
			"input %s should have resulted in %t but was %t", inputValue, currentTestcase.expectedValue, actualTraceValue,
		)
	}
}
