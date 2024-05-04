package monitor

import (
	"context"
	"testing"
	"time"
)

func TestPing(t *testing.T) {
	tests := []struct {
		URL string
		Up  bool
	}{
		{"encore.dev", true},
		{"google.com", true},
		// Test both with and without "https://"
		{"httpbin.org/status/200", true},
		{"https://httpbin.org/status/200", true},

		// 4xx and 5xx should considered down.
		{"httpbin.org/status/400", false},
		{"https://httpbin.org/status/500", false},
		// Invalid URLs should be considered down.
		{"invalid://scheme", false},
	}

	for _, test := range tests {
		ctx, _ := context.WithTimeout(context.Background(), 4*time.Second)
		resp, err := Ping(ctx, test.URL)
		if err != nil {
			t.Errorf("url %s: unexpected error: %v", test.URL, err)
		} else if resp.Up != test.Up {
			t.Errorf("url %s: got up=%v, want %v", test.URL, resp.Up, test.Up)
		}
	}
}