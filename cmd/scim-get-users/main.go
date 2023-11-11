package main

import (
	"fmt"
	"log"

    
	"github.com/jawee/scim-client-go/internal/models"
)


func main() {
    sourceUsers := getSourceUsers()

    toCreate := make([]models.User, 0)
    toUpdate := make([]models.User, 0)

    for _, v := range sourceUsers {
        _, err := getDbUserById(v.Id)
        if err != nil {
            toCreate = append(toCreate, v)
            continue
        }
        //TODO: Check diff between existing and new
        toUpdate = append(toUpdate, v)
    }

    for _, v := range toCreate {
        log.Printf("ToCreate: %v\n", v)
    }

    for _, v := range toUpdate {
        log.Printf("ToUpdate: %v\n", v)
    }
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

func getSourceUsers() []models.User {
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
            FirstName: "Other",
            LastName: "User",
            Active: true,
            ExternalId: "",
        },
    }
    return users
}

