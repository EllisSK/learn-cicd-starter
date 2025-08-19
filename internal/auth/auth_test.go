package auth

import (
	"net/http"
	"reflect"
	"testing"
)

func TestGetApiKey1(t *testing.T) {
	actualValue, _ := GetAPIKey(http.Header{})
	expectedValue := ""
	if !reflect.DeepEqual(expectedValue, actualValue) {
		t.Fatalf("expected: %v, got: %v", expectedValue, actualValue)
	}
}

func TestGetApiKey2(t *testing.T) {
	headers := http.Header{}
	headers.Add("Authorization", "ApiKey xyz123")
	actualValue, _ := GetAPIKey(headers)
	expectedValue := "xyz123"
	if reflect.DeepEqual(expectedValue, actualValue) {
		t.Fatalf("expected: %v, got: %v", expectedValue, actualValue)
	}
}
