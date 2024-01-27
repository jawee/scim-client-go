package flags

import (
	"testing"
)

func TestParseFlagsErrors(t *testing.T) {
    testCases := []struct {
        args     []string
        expected string
    }{
        {args: []string{"-a", "b"}, expected: "Invalid flag -a" },
        {args: []string{"-c"}, expected: "Missing argument for -c" },
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

func TestParseFlagsSuccess(t *testing.T) {
    testCases := []struct {
        args     []string
        expected map[FlagType]Flag
    }{
        {args: []string{"-d"}, expected: map[FlagType]Flag{ Delta: {Type: Delta} }}, 
        {args: []string{"--delta"}, expected: map[FlagType]Flag{ Delta: {Type: Delta} }}, 
        {args: []string{"-c", "/path/to/config"}, expected: map[FlagType]Flag{ ConfigDir: {Type: ConfigDir, Value: "/path/to/config"} }}, 
        {args: []string{"--configDir", "/path/to/config"}, expected: map[FlagType]Flag{ ConfigDir: {Type: ConfigDir, Value: "/path/to/config"} }}, 
        {args: []string{"-c", "/path/to/config", "-d"}, expected: map[FlagType]Flag{ 
                ConfigDir: {Type: ConfigDir, Value: "/path/to/config"},
                Delta: {Type: Delta },
            },
        }, 
    }

    for _, v := range testCases {
        res, err := ParseFlags(v.args)
        if err != nil {
            t.Fatalf("Got err '%s'\n", err)
        }

        for _, f := range res {
            expected, ok := v.expected[f.Type]
            if !ok {
                t.Fatalf("Couldn't find expected flag of type '%v'\n", f.Type)
            }
            if expected.Value != f.Value {
                t.Fatalf("Got value '%s', expected '%s'\n", f.Value, expected.Value)
            }
        }
    }
}
