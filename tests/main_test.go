package main

import (
	"encoding/base64"
	"testing"
)

func TestBase64Decode(t *testing.T) {
	// Test case 1: Valid Base64 string
	input1 := "SGVsbG8gV29ybGQh" // "Hello World!"
	expected1 := "Hello World!"
	result1, err1 := base64.StdEncoding.DecodeString(input1)
	if err1 != nil || string(result1) != expected1 {
		t.Errorf("Test case 1 failed. Expected: %s, Got: %s, Error: %v", expected1, result1, err1)
	}

	// Test case 2: Invalid Base64 string
	input2 := "InvalidBase64String"
	_, err2 := base64.StdEncoding.DecodeString(input2)
	if err2 == nil {
		t.Errorf("Test case 2 failed. Expected an error for invalid Base64 string, but got nil.")
	}
}
