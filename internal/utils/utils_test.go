package utils

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIndexOfStringInSliceHappy(t *testing.T) {
	setupStrArry := []string{"hello", "is", "it", "me", "you're", "looking", "for"}
	result := IndexOfStringInSlice("hello", setupStrArry...)
	if result != 0 {
		t.Errorf("lionel richie is sad you couldn't find him at index 0")
	}
}

func TestIndexOfStringInSliceUnhappy(t *testing.T) {
	setupStrArry := []string{"cuz", "you", "got", "that"}
	result := IndexOfStringInSlice("cheeseburgers", setupStrArry...)
	if result != -1 {
		t.Errorf("wait wut...we actually found some cheeseburgers...something wrong in all the right ways...")
	}
}

func TestMustGitProjectRootDirectoryDumbHappyTest(t *testing.T) {
	result := MustGitProjectRootDirectory()
	if len(result) == 0 {
		t.Errorf("uh oh, guess we're not in a project root running test, or something broke?" +
			" -could get crafty and ensure that where this file was, we always navigated to this directory to check _shrug_ hehh")
	}
}

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
