package utils

import "testing"

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
