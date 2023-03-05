package utils

import "testing"

func Test_MD5Hash(t *testing.T) {
	// Test data and expected result
	testData := []byte("test data")
	expectedHash := "eb733a00c0c9d336e65691a37ab54293"

	// Call the function with the test data
	actualHash := MD5Hash(testData)

	// Compare the actual result with the expected result
	if actualHash != expectedHash {
		t.Errorf("MD5Hash(%v) = %v, expected %v", testData, actualHash, expectedHash)
	}
}
