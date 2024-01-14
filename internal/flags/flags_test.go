package flags

import (
	"testing"
)

// var ErrNotFound = errors.New("not found")
func TestParseFlagsErrors(t *testing.T) {
    testCases := []struct {
        args     []string
        expected string
    }{
        {args: []string{"-c"}, expected: "Expects an even number of arguments" },
        {args: []string{"-a", "b"}, expected: "Invalid flag -a" },
    }

    for _, v := range testCases {
        _, err := ParseFlags(v.args)
        if err == nil {
            t.Fatalf("Got err nil, expected '%s'\n", v.expected)
        }

        if err.Error() != v.expected {
            t.Fatalf("Got err '%s', expected '%s'\n", err, v.expected)
        }
    }
}
