package api

import (
	"testing"
)

// Test the origin validation
func TestOriginValidator(t *testing.T) {
	testcases := []struct {
		origin string
		allow  bool
	}{
		// Should be allowed
		{"https://trezor.io", true},
		{"https://foo.trezor.io", true},
		{"https://bar.foo.trezor.io", true},
		// Should be denied
		{"https://faketrezor.io", false},
		{"https://foo.faketrezor.io", false},
		{"https://foo.trezor.ioo", false},
		{"http://foo.trezor.io", false},
		// Localhost 8xxx and 5xxx should be allowed for local development
		{"https://localhost:8000", true},
		{"http://localhost:8000", true},
		{"http://localhost:8999", true},
		{"https://localhost:5000", true},
		{"http://localhost:5000", true},
		{"http://localhost:5999", true},
		// SatoshiLabs dev servers should be allowed
		{"https://sldev.cz", true},
		{"https://foo.sldev.cz", true},
		{"https://bar.foo.sldev.cz", true},
		// Should be denied
		{"https://fakesldev.cz", false},
		{"https://foo.fakesldev.cz", false},
		{"https://foo.sldev.czz", false},
		{"http://foo.trezor.sldev.cz", false},
		// Other ports denied
		{"http://localhost", false},
		{"http://localhost:1234", false},

		{"http://192.168.0.55", true},
		{"http://192.168.0.55:80", true},
		{"http://192.168.0.143:5999", true},
		{"http://192.168.0.143:8999", true},
		{"https://test.bishangex.com", true},
		{"https://test.btc.so:899", true},
		{"https://test.socoin.cc:9000", true},
		{"https://test.b2sit.xyz", true},
		{"https://www.b2sit.xyz:10086", true},
		{"https://b2sim.xyz:10086", true},
		{"https://test.b1dev.xyz", true},
	}
	validator, err := corsValidator()
	if err != nil {
		t.Fatal(err)
	}
	for _, tc := range testcases {
		allow := validator(tc.origin)
		if allow != tc.allow {
			t.Errorf("Origin %q: expected %v, got %v", tc.origin, tc.allow, allow)
		}
	}
}
