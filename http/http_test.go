package http

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
)

func Test_Fetch(t *testing.T) {
	// Test case 1: successful HTTP GET request
	url1 := "https://jsonplaceholder.typicode.com/todos/1"
	expectedBody1 := map[string]interface{}{
		"userId":    float64(1),
		"id":        float64(1),
		"title":     "delectus aut autem",
		"completed": false,
	}
	body1, err1 := Fetch(url1)
	if err1 != nil {
		t.Errorf("Fetch(%q) returned an error: %v", url1, err1)
	}
	var actualBody1 map[string]interface{}
	if err := json.Unmarshal(body1, &actualBody1); err != nil {
		t.Errorf("Fetch(%q) returned invalid JSON body: %v", url1, err)
	}
	if !reflect.DeepEqual(actualBody1, expectedBody1) {
		t.Errorf("Fetch(%q) returned unexpected body: got %v, want %v", url1, actualBody1, expectedBody1)
	}

	// Test case 2: non-existent URL
	url2 := "https://example.com/non-existent"
	expectedError2 := fmt.Errorf("unexpected status code: 404")
	_, err2 := Fetch(url2)
	if err2 == nil {
		t.Errorf("Fetch(%q) did not return an error, expected %v", url2, expectedError2)
	}
	if err2.Error() != expectedError2.Error() {
		t.Errorf("Fetch(%q) returned unexpected error: got %v, want %v", url2, err2, expectedError2)
	}

	// Test case 3: server returns non-OK HTTP status code
	url4 := "https://jsonplaceholder.typicode.com/todos/invalid"
	expectedError4 := fmt.Errorf("unexpected status code: 404")
	_, err4 := Fetch(url4)
	if err4 == nil {
		t.Errorf("Fetch(%q) did not return an error, expected %v", url4, expectedError4)
	}
	if err4.Error() != expectedError4.Error() {
		t.Errorf("Fetch(%q) returned unexpected error: got %v, want %v", url4, err4, expectedError4)
	}
}
