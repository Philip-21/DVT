package validate

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateDatetime(t *testing.T) {
	// Test valid timestamps
	tests := []struct {
		timestamp string
		valid     bool
	}{
		{"2023-07-10T16:22:41+00:00", true},
		{"2023-07-10T16:22:41+01:00", true},
		{"2023-07-10T16:22:41+02:00", true},
		{"2023-07-10T16:22:41-06:00", true},
		{"2023-07-10T16:22:41+03:00", false}, // Invalid timezone offset
		{"invalid-timestamp", false},         // Completely invalid format
	}
	for _, test := range tests {
		//test the ValidateDateTime(), that returns the actual ValidateDateTime()
		timezone, valid, err := ValidateDateTime(test.timestamp)
		if valid != test.valid {
			t.Errorf("expected validity: %v, got: %v, error: %v", test.valid, valid, err)
		}
		if test.valid && err != nil {
			t.Errorf("expected no error, got: %v", err)
		}
		if !test.valid && err == nil {
			t.Errorf("expected an error for invalid timestamp: %s", test.timestamp)
		}
		assert.NotNil(t, timezone)
	}

}
