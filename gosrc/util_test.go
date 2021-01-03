package gosrc

import "testing"

func TestCheckTcpConnect(t *testing.T) {
	var urlTests = []struct {
		expected bool   // expected result
		in       string // input
	}{
		{true, "1024"},
		{true, "2048"},
		{false, "80"},
		{false, "8080"},
	}

	for _, tt := range urlTests {
		actual := CheckTcpConnect("0.0.0.0", tt.in)
		if (nil == actual) != tt.expected {
			t.Errorf("CheckTcpConnect(%s) = %t; expected %t", tt.in, actual, tt.expected)
		}
	}
}
