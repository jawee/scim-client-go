package main

import (
	"fmt"
)

func main() {
    // internal.MainFunc()

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

func getDiff(user1, user2 User) []string {
    result := []string{}
    if user1.Id != user2.Id {
        result = append(result, "Id")
    }

    if user1.UserName != user2.UserName {
        result = append(result, "UserName")
    }

    if user1.Email != user2.Email {
        result = append(result, "Email")
    }
    
    if user1.Department != user2.Department {
        result = append(result, "Department")
    }

    if user1.PhoneNumber != user2.PhoneNumber {
        result = append(result, "PhoneNumber")
    }

    if user1.FirstName != user2.FirstName {
        result = append(result, "FirstName")
    }

    if user1.LastName != user2.LastName {
        result = append(result, "LastName")
    }

    if user1.Active != user2.Active {
        result = append(result, "Active")
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
