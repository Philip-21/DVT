package json

import (
	"testing"
)

var testEmails = []struct {
	ValidEmail   string
	InvalidEmail string
	InvalidData  bool
}{
	{
		ValidEmail:  "harry@gmail.com",
		InvalidData: false,
	},
	{
		InvalidEmail: "lesleygmail.com",
		InvalidData:  true,
	},
	{
		ValidEmail:  "dennis@gmail.com",
		InvalidData: false,
	},
	{
		ValidEmail:   "Thiery@gmail.com",
		InvalidEmail: "lesleygmail.com",
		InvalidData:  true,
	},
	{
		ValidEmail:   "dennis@gmail.com",
		InvalidEmail: "invalid-email",
		InvalidData:  true,
	},
}

func TestTransformEmailString(t *testing.T) {
	for _, ts := range testEmails {
		_, err := EmailToString(ts.ValidEmail)
		if err != nil {
			t.Log(err)
			return
		}
		if ts.InvalidData == true {
			t.Log("Invalid data inputed ")
		}
		_, err = EmailToString(ts.InvalidEmail)
		if err != nil {
			t.Log(err)
			return
		}
		if ts.InvalidData == true {
			t.Log("Invalid data inputed ")
		}
	}
	t.Log("Test Passed")
}
