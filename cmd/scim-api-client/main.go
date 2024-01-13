package main

import (
	"fmt"
	"log"
	"os"

	"github.com/jawee/scim-client-go/internal/models"
	"github.com/jawee/scim-client-go/internal/readers"
	scimapi "github.com/jawee/scim-client-go/internal/scim-api"
)

func usage() {
	fmt.Printf("Usage:\n    scim-client [flags]\n")
	fmt.Printf("Flags:\n")
    fmt.Printf("    %s        %s\n", "-h, --help", "Print help")
    fmt.Printf("    %s      %s\n", "-c, --config", "Path to config directory")
    fmt.Printf("    %s       %s\n", "-i, --input", "Source of users to import")
}

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
        return Flag{}, fmt.Errorf("Invalid flag")
    }

    flag := Flag{
        Type: flagType,
        Value: v,
    }

    return flag, nil
}
func ParseFlags(args []string) ([]Flag, error) {
    res := []Flag{}
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

func getConfigPath(flags []Flag) (string, error) {
    defaultPath, err := os.UserConfigDir()
    if err != nil {
        return "", err
    }
    customPath := "" 
    for _, v := range flags {
        if v.Type == Config {
            customPath = v.Value
        }
    }

    if customPath == "" {
        customPath = defaultPath + "/scimclient"
    }

    _, err = os.Stat(customPath)
    if os.IsNotExist(err) {
        fmt.Printf("ERROR: directory does not exist\n")
        return "", err
    }

    if err != nil {
        return "", err
    }

    return customPath, nil
}

func main() {
    args := os.Args[1:]
    if len(args) == 0 {
        fmt.Printf("ERROR: No flags provided\n")
        usage()
        os.Exit(1)
    }

    flags, err := ParseFlags(args)
    if err != nil {
        fmt.Printf("ERROR: %s\n", err)
        usage()
        os.Exit(1)
    }

    if len(flags) == 0 {
        fmt.Printf("ERROR: No flags parsed\n")
        usage()
        os.Exit(1)
    }
    // fmt.Printf("Len: %d, args: %v\n", len(args), args)

    configPath, err := getConfigPath(flags)
    if err != nil {
        fmt.Printf("ERROR: %s\n", err)
        os.Exit(1)
    }
    fmt.Printf("%s\n", configPath)
    return
    reader := readers.MemoryReader{}

    sourceUsers, err := reader.GetUsers()
    if err != nil {
        log.Printf("%s\n", err)
        return
    }

    sourceUsersMap := makeMap(sourceUsers)
    dbUsers, err := getDbUsers()
    if err != nil {
    }

    if err != nil {
        log.Printf("%s\n", err)
        return
    }

    toHandle := []models.User{}
    toDelete := []models.User{}

    for _, v := range sourceUsers {
        toHandle = append(toHandle, v)
    }

    for k, v := range dbUsers {
        if _, ok := sourceUsersMap[k]; !ok {
            toDelete = append(toDelete, v)
        }
    }

    for _, user := range toHandle {
        id, err := scimapi.HandleUser(&user)

        if err != nil {
            log.Printf("%s\n", err)
            return
        }
        log.Printf("ExternalId: %s\n", id)
    }

    for _, user := range toDelete {
        succ, err := scimapi.DeleteUser(&user)
        if err != nil { 
            log.Printf("Delete error for user %s: %s\n", user.UserName, err)
        } else {
            log.Printf("Delete result for user %s: %v\n", user.UserName, succ)
        }
    }
}

func makeMap(sourceUsers []models.User) map[string]models.User {
    usersMap := map[string]models.User{}
    for _, v := range sourceUsers {
        usersMap[v.Id] = v
    }
    return usersMap
}

func getDbUsers() (map[string]models.User, error) {
    users := []models.User{
        {
            Id: "1",
            UserName: "some.user@company.name",
            Email: "some.user@company.name",
            Department: "clown",
            PhoneNumber: "12345678",
            FirstName: "Some",
            LastName: "User",
            Active: true,
            ExternalId: "",
        },
        {
            Id: "2",
            UserName: "other.user@company.name",
            Email: "other.user@company.name",
            Department: "jester",
            PhoneNumber: "87654321",
            FirstName: "Other",
            LastName: "User",
            Active: true,
            ExternalId: "",
        },
        {
            Id: "3",
            UserName: "third.user@company.name",
            Email: "third.user@company.name",
            Department: "jester",
            PhoneNumber: "87654321",
            FirstName: "Third",
            LastName: "User",
            Active: true,
            ExternalId: "",
        },
        {
            Id: "4",
            UserName: "fourth.user@company.name",
            Email: "fourth.user@company.name",
            Department: "",
            PhoneNumber: "",
            FirstName: "Fourth",
            LastName: "User",
            Active: true,
            ExternalId: "",
        },
    }

    m := map[string]models.User{}
    for _, v := range users {
        m[v.Id] = v
    }
    return m, nil
}
func getDbUserById(id string) (*models.User, error) {
    users := []models.User {
        {
            Id: "1",
            UserName: "some.user@company.name",
            Email: "some.user@company.name",
            Department: "clown",
            PhoneNumber: "12345678",
            FirstName: "Some",
            LastName: "User",
            Active: true,
            ExternalId: "",
        },
        {
            Id: "2",
            UserName: "other.user@company.anem",
            Email: "other.user@company.name",
            Department: "jester",
            PhoneNumber: "87654321",
            FirstName: "Other",
            LastName: "User",
            Active: true,
            ExternalId: "",
        },
    }

    var res *models.User
    for _, v := range users {
        if v.Id == id {
            res = &v
            break
        }
    }

    if res == nil {
        return nil, fmt.Errorf("Couldn't find user\n")
    }

    return res, nil
}
