package main

import (
	"log"

	"github.com/jawee/scim-client-go/internal/models"
	"github.com/jawee/scim-client-go/internal/scim-api"
)

func main() {
    user := models.User {
        Id: "asdf",
        UserName: "some.user@company.name",
        Email: "some.user@company.name",
        Department: "clown",
        PhoneNumber: "12345689",
        FirstName: "Some",
        LastName: "User",
        Active: true,
        ExternalId: "",
    }
    id, err := scimapi.HandleUser(&user)

    if err != nil {
        log.Printf("%s\n", err)
        return
    }

    log.Printf("ExternalId: %s\n", id)
}
