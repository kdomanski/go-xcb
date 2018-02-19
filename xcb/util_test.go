package xcb

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseDisplay(t *testing.T) {
	defer func(oldValue string) {
		os.Setenv("DISPLAY", oldValue)
	}(os.Getenv("DISPLAY"))
	os.Setenv("DISPLAY", "")

	for _, testcase := range []struct {
		input            string
		expectedHostname string
		expectedDisplay  int
		expectedScreen   int
		expectedError    bool
	}{
		{":0", "", 0, 0, false},
		{":1", "", 1, 0, false},
		{":0.1", "", 0, 1, false},

		{"x.org:0", "x.org", 0, 0, false},
		{"expo:0", "expo", 0, 0, false},
		{"bigmachine:1", "bigmachine", 1, 0, false},
		{"hydra:0.1", "hydra", 0, 1, false},

		{"198.112.45.11:0", "198.112.45.11", 0, 0, false},
		{"198.112.45.11:0.1", "198.112.45.11", 0, 1, false},

		{":::0", "::", 0, 0, false},
		{"1:::0", "1::", 0, 0, false},
		{"::1:0", "::1", 0, 0, false},
		{"::1:0.1", "::1", 0, 1, false},
		{"::127.0.0.1:0", "::127.0.0.1", 0, 0, false},
		{"::ffff:127.0.0.1:0", "::ffff:127.0.0.1", 0, 0, false},
		{"2002:83fc:d052::1:0", "2002:83fc:d052::1", 0, 0, false},
		{"2002:83fc:d052::1:0.1", "2002:83fc:d052::1", 0, 1, false},
		{"[::]:0", "[::]", 0, 0, false},
		{"[1::]:0", "[1::]", 0, 0, false},
		{"[::1]:0", "[::1]", 0, 0, false},
		{"[::1]:0.1", "[::1]", 0, 1, false},
		{"[::127.0.0.1]:0", "[::127.0.0.1]", 0, 0, false},
		{"[::ffff:127.0.0.1]:0", "[::ffff:127.0.0.1]", 0, 0, false},
		{"[2002:83fc:d052::1]:0", "[2002:83fc:d052::1]", 0, 0, false},
		{"[2002:83fc:d052::1]:0.1", "[2002:83fc:d052::1]", 0, 1, false},

		{"myws::0", "myws:", 0, 0, false},
		{"big::1", "big:", 1, 0, false},
		{"hydra::0.1", "hydra:", 0, 1, false},

		{"", "", 0, 0, true},
		{":", "", 0, 0, true},
		{"::", "", 0, 0, true},
		{":::", "", 0, 0, true},
		{":.", "", 0, 0, true},
		{":a", "", 0, 0, true},
		{":a.", "", 0, 0, true},
		{":0.", "", 0, 0, true},
		{":.a", "", 0, 0, true},
		{":.0", "", 0, 0, true},
		{":0.a", "", 0, 0, true},
		{":0.0.", "", 0, 0, true},
	} {
		t.Run(testcase.input, func(t *testing.T) {
			hostname, display, screen, err := ParseDisplay(testcase.input)
			if testcase.expectedError {
				assert.NotNil(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, testcase.expectedHostname, hostname)
				assert.Equal(t, testcase.expectedDisplay, display)
				assert.Equal(t, testcase.expectedScreen, screen)
			}
		})
	}
}
