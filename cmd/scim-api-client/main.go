package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"

	"github.com/jawee/scim-client-go/internal/flags"
	"github.com/jawee/scim-client-go/internal/models"
	"github.com/jawee/scim-client-go/internal/readers"
	"github.com/jawee/scim-client-go/internal/services"
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
    go startWebServer()

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
        fmt.Printf("ERROR: %s\n", err)
        os.Exit(1)
    }

    configFile, err := os.ReadFile(configPath + "/config.json")
    if err != nil {
        fmt.Printf("ERROR: %s\n", err)
        os.Exit(1)
    }
    fmt.Printf("%s\n", string(configFile))
    // return
    reader := readers.MemoryReader{}

    dbUsers := getDbUsers()

    services.ExecuteSync(reader, dbUsers)
}

func getDbUsers() (map[string]models.UserHistory) {
    users := []models.UserHistory{
        {
            UserName: "some.user@company.name",
            ErrorMessages: nil,
            LastSync: time.Now().Add(time.Duration(-120)),
        },
        {
            UserName: "other.user@company.name",
            ErrorMessages: nil,
            LastSync: time.Now().Add(time.Duration(-120)),
        },
        {
            UserName: "third.user@company.name",
            ErrorMessages: nil,
            LastSync: time.Now().Add(time.Duration(-120)),
        },
        {
            UserName: "fourth.user@company.name",
            ErrorMessages: nil,
            LastSync: time.Now().Add(time.Duration(-120)),
        },
    }

    m := map[string]models.UserHistory{}
    for _, v := range users {
        m[v.UserName] = v
    }
    return m
}

// func getDbUsers() (map[string]models.User, error) {
//     users := []models.User{
//         {
//             Id: "1",
//             UserName: "some.user@company.name",
//             Email: "some.user@company.name",
//             Department: "clown",
//             PhoneNumber: "12345678",
//             FirstName: "Some",
//             LastName: "User",
//             Active: true,
//             ExternalId: "",
//         },
//         {
//             Id: "2",
//             UserName: "other.user@company.name",
//             Email: "other.user@company.name",
//             Department: "jester",
//             PhoneNumber: "87654321",
//             FirstName: "Other",
//             LastName: "User",
//             Active: true,
//             ExternalId: "",
//         },
//         {
//             Id: "3",
//             UserName: "third.user@company.name",
//             Email: "third.user@company.name",
//             Department: "jester",
//             PhoneNumber: "87654321",
//             FirstName: "Third",
//             LastName: "User",
//             Active: true,
//             ExternalId: "",
//         },
//         {
//             Id: "4",
//             UserName: "fourth.user@company.name",
//             Email: "fourth.user@company.name",
//             Department: "",
//             PhoneNumber: "",
//             FirstName: "Fourth",
//             LastName: "User",
//             Active: true,
//             ExternalId: "",
//         },
//     }
//
//     m := map[string]models.User{}
//     for _, v := range users {
//         m[v.Id] = v
//     }
//     return m, nil
// }

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


func startWebServer() {
    var err error
    if _, err = template.New("index").Parse(Index); err != nil {
        log.Printf("%s\n", err)
    }
    if err != nil {
        log.Fatalf("%s\n", err)
    }

    http.HandleFunc("/", getHandler)

    log.Printf("Listening on port 8080\n")
    log.Fatal(http.ListenAndServe(":8080", nil))
}

type indexparameters struct {
    Configurations []string
}

func getHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method == "GET" {
        t, err := template.New("index").Parse(Index);
        if (err != nil) {
            fmt.Fprintf(w, "template error %s\n", err)
        }
        idxParams := indexparameters{ Configurations: []string{"A","B"}}
        t.Execute(w, idxParams)
    } else if r.Method == "PUT" {
        log.Printf("Got put\n")
    }
}



var Index = `<!DOCTYPE html>
<html lang="en">
    <head>
        <title>SCIM Client Configurator</title>
        <meta charset="UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1">
        <script src="https://unpkg.com/htmx.org@1.9.4" integrity="sha384-zUfuhFKKZCbHTY6aRR46gxiqszMk5tcHjsVFxnUo8VMus4kHGVdIYVbOYYNlKmHV" crossorigin="anonymous"></script>
        <link href="css/style.css" rel="stylesheet">
    </head>
    <body>
        <div class="container">
            <h1>Hello World</h1>
            <select name="make" hx-get="/models" hx-target="#models" hx-indicator=".htmx-indicator">
                {{range .Configurations}}
                <option value="{{.}}">{{.}}</option>
                {{end}}
            </select>
        </div>
    </body>
</html>`
