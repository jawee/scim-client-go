package main

import (
	"fmt"
	"log"

	"github.com/jawee/scim-client-go/internal/models"
	"github.com/jawee/scim-client-go/internal/readers"
	scimapi "github.com/jawee/scim-client-go/internal/scim-api"
)

func main() {
    //TODO: make configureable through args and/or config file
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
        //TODO: handle delete
        log.Printf("ToDelete: %v\n", user)
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
            UserName: "other.user@company.anem",
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
            UserName: "third.user@company.anem",
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
            UserName: "fourth.user@company.anem",
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
