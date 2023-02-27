package mytestsdk

import "testing"

func TestGenerateRandomString(t *testing.T) {
	result := GenerateRandomString(10)

	if len(result) != 10 {
		t.Errorf("Generated string has incorrect length: %d", len(result))
	}
}
