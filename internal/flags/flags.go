package flags

import (
	"fmt"
	"strings"
)

type FlagType int

const (
    Invalid FlagType = iota
    ConfigDir
    Input
    Delta
)

type Flag struct {
   Type FlagType
   Value string
}

func getFlagType(t string) FlagType {
    var flagType FlagType;
    switch t {
    case "--configDir":
        flagType = ConfigDir
    case "-c":
        flagType = ConfigDir
    case "-i":
        flagType = Input
    case "--input":
        flagType = Input
    case "-d":
        flagType = Delta
    case "--delta":
        flagType = Delta
    default:
        flagType = Invalid 
    }

    return flagType
}

var numberOfArgumentsMap = map[FlagType]int{
    ConfigDir: 1,
    Input: 1,
    Delta: 0,
}

func ParseFlags(args []string) ([]Flag, error) {
    res := []Flag{}

    for i := 0; i < len(args); {
        flagType := getFlagType(args[i])
        if flagType == Invalid {
            return nil, fmt.Errorf("Invalid flag %s", args[i])
        }

        arguments := numberOfArgumentsMap[flagType]

        if arguments == 0 {
            flag := Flag{Type: flagType}
            res = append(res, flag)
        }
        if arguments == 1 {
            if i+1 >= len(args) {
                return nil, fmt.Errorf("Missing argument for %s", args[i])
            }
            if strings.HasPrefix(args[i+1], "-") {
                return nil, fmt.Errorf("Missing argument for %s", args[i])
            }
            flag := Flag{Type: flagType, Value: args[i+1]}
            res = append(res, flag)
        }

        i += arguments + 1
    }

    return res, nil
}

