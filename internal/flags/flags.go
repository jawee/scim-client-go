package flags

import "fmt"

type FlagType int

const (
    Invalid FlagType = iota
    Config
    Input
)

type Flag struct {
   Type FlagType
   Value string
}

func getFlag(t string, v string) (Flag, error) {
    var flagType FlagType;
    switch t {
    case "--config":
        flagType = Config
    case "-c":
        flagType = Config
    case "-i":
        flagType = Input
    case "--input":
        flagType = Input
    default:
        flagType = Invalid 
    }

    if flagType == Invalid {
        return Flag{}, fmt.Errorf("Invalid flag %s", t)
    }

    flag := Flag{
        Type: flagType,
        Value: v,
    }

    return flag, nil
}

func ParseFlags(args []string) ([]Flag, error) {
    res := []Flag{}
    //TODO should accept flags without arguments
    if len(args) % 2 != 0 {
        return nil, fmt.Errorf("Expects an even number of arguments")
    }
    for i := 0; i < len(args); i += 2 {
        flag, err := getFlag(args[i], args[i+1])
        if err != nil {
            return nil, err
        }
        res = append(res, flag)
    }

    return res, nil
}

