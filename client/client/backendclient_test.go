package client

import (
	"testing"
)

func Test(t *testing.T) {
	var backend, err = NewFromUrlAddress("invalid url")

	if err == nil {
		var users, err = backend.GetUserData()
		if err == nil && users == nil {
			t.Error("error is expected, but err var is nil or users are filled")
		}
	}
}