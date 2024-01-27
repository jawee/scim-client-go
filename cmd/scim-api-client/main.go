package main

import (
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/jawee/scim-client-go/internal/flags"
	"github.com/jawee/scim-client-go/internal/models"
	"github.com/jawee/scim-client-go/internal/readers"
	scimapi "github.com/jawee/scim-client-go/internal/scim-api"
)

func usage() {
	fmt.Printf("Usage:\n    scim-client [flags]\n")
	fmt.Printf("Flags:\n")
    fmt.Printf("    %s        %s\n", "-h, --help", "Print help")
    fmt.Printf("    %s      %s\n", "-c, --configDir", "Path to config directory")
    fmt.Printf("    %s       %s\n", "-i, --input", "Source of users to import")
    fmt.Printf("    %s       %s\n", "-d, --delta", "Delta import, only users included in this run")
}

func getConfigPath(f []flags.Flag) (string, error) {
    defaultPath, err := os.UserConfigDir()
    if err != nil {
        return "", err
    }
    customPath := "" 
    for _, v := range f {
        if v.Type == flags.ConfigDir {
            customPath = v.Value
        }
    }

    if customPath == "" {
        customPath = defaultPath + "/scimclient"
    }

    ex, err := os.Executable()
    if err != nil {
        return "", err
    }

    exPath := filepath.Dir(ex)
    exPath, _ = strings.CutSuffix(exPath, "/")

    cust, foundPrefix := strings.CutPrefix(customPath, "./"); 
    if foundPrefix {
        customPath = path.Join(exPath, cust)
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

    flags, err := flags.ParseFlags(args)
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

    _, err = os.Stat(configPath)
    if os.IsNotExist(err) {
        fmt.Printf("ERROR: directory does not exist\n")
        os.Exit(1)
    }

    if err != nil {
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
