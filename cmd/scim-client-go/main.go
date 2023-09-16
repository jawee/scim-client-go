package main

import (
	"fmt"
	"os"

	"github.com/jawee/scim-client-go/internal/configuration"
)

func main() {
    // time.Sleep(5* time.Second);

    configProvider := new(configuration.FileConfigurationProvider)
    configuration, err := configuration.New(configProvider)

    if err != nil {
        // log.Println(err)
        os.Exit(1)
    }

    if configuration == nil {
        // log.Println("Configuration is nil")
        os.Exit(1)
    }    

    user1 := User {
        Id: "asdf",
        UserName: "some.user@company.name",
        Email: "some.user@company.name",
        Department: "clown",
        PhoneNumber: "12345678",
        FirstName: "Some",
        LastName: "User",
        Active: true,
        ExternalId: "",
    }

    user2 := User {
        Id: "asdf",
        UserName: "some@user.name",
        Email: "some@user.name",
        Department: "clown",
        PhoneNumber: "12345678",
        FirstName: "Some",
        LastName: "User",
        Active: true,
        ExternalId: "",
    }

    res := getDiff(user1, user2)

    fmt.Printf("Properties to be updated: %v\n", res)
}

type Attribute string

const (
    UserNameAttribute Attribute = "UserName"
    IdAttribute Attribute = "Id"
    DepartmentAttribute Attribute = "Department"
    PhoneNumberAttribute Attribute = "PhoneNumber"
    EmailAttribute Attribute = "Email"
    FirstNameAttribute Attribute =  "FirstName"
    LastNameAttribute Attribute =  "LastName"
    ActiveAttribute Attribute =  "Active"
)

func getDiff(user1, user2 User) []Attribute {
    result := []Attribute{}
    if user1.Id != user2.Id {
        result = append(result, "Id")
    }

    if user1.UserName != user2.UserName {
        result = append(result, UserNameAttribute)
    }

    if user1.Email != user2.Email {
        result = append(result, EmailAttribute)
    }
    
    if user1.Department != user2.Department {
        result = append(result, DepartmentAttribute)
    }

    if user1.PhoneNumber != user2.PhoneNumber {
        result = append(result, PhoneNumberAttribute)
    }

    if user1.FirstName != user2.FirstName {
        result = append(result, FirstNameAttribute)
    }

    if user1.LastName != user2.LastName {
        result = append(result, LastNameAttribute)
    }

    if user1.Active != user2.Active {
        result = append(result, ActiveAttribute)
    }

    return result
}

type User struct {
    Id string
    UserName string
    Email string
    PhoneNumber string
    Department string
    FirstName string
    LastName string
    Active bool
    ExternalId string
}
