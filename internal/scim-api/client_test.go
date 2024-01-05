package scimapi

import (
	"testing"

	"github.com/jawee/scim-client-go/internal/models"
)

func TestDiffFirstName(t *testing.T) {
    existingUser := models.User {
        FirstName: "a",
    }
    newUser := models.User {
        FirstName: "b",
    }

    changes, err := diffUsers(&existingUser, &newUser)

    if err != nil {
        t.Fatalf("Got error '%s'\n", err)
    }
    if len(changes) != 1 {
        t.Fatalf("Got '%d' changes, expected '1'\n", len(changes))
    }
    change := changes[0]

    if change.FieldName != "firstname" {
        t.Fatalf("Got FieldName '%s', expected 'firstname'\n", change.FieldName)
    }
    if change.Value != "b" {
        t.Fatalf("Got Value '%s', expected 'b'\n", change.Value)
    }
}

func TestDiffLastName(t *testing.T) {
    existingUser := models.User {
        LastName: "a",
    }
    newUser := models.User {
        LastName: "b",
    }

    changes, err := diffUsers(&existingUser, &newUser)

    if err != nil {
        t.Fatalf("Got error '%s'\n", err)
    }
    if len(changes) != 1 {
        t.Fatalf("Got '%d' changes, expected '1'\n", len(changes))
    }
    change := changes[0]

    if change.FieldName != "lastname" {
        t.Fatalf("Got FieldName '%s', expected 'lastname'\n", change.FieldName)
    }
    if change.Value != "b" {
        t.Fatalf("Got Value '%s', expected 'b'\n", change.Value)
    }
}

func TestDiffPhoneNumber(t *testing.T) {
    existingUser := models.User {
        PhoneNumber: "a",
    }
    newUser := models.User {
        PhoneNumber: "b",
    }

    changes, err := diffUsers(&existingUser, &newUser)

    if err != nil {
        t.Fatalf("Got error '%s'\n", err)
    }
    if len(changes) != 1 {
        t.Fatalf("Got '%d' changes, expected '1'\n", len(changes))
    }
    change := changes[0]

    if change.FieldName != "phonenumber" {
        t.Fatalf("Got FieldName '%s', expected 'phonenumber'\n", change.FieldName)
    }
    if change.Value != "b" {
        t.Fatalf("Got Value '%s', expected 'b'\n", change.Value)
    }
}

func TestDiffActive(t *testing.T) {
    existingUser := models.User {
        Active: true,
    }
    newUser := models.User {
        Active: false,
    }

    changes, err := diffUsers(&existingUser, &newUser)

    if err != nil {
        t.Fatalf("Got error '%s'\n", err)
    }
    if len(changes) != 1 {
        t.Fatalf("Got '%d' changes, expected '1'\n", len(changes))
    }
    change := changes[0]

    if change.FieldName != "active" {
        t.Fatalf("Got FieldName '%s', expected 'active'\n", change.FieldName)
    }
    if change.Value != "false" {
        t.Fatalf("Got Value '%s', expected 'b'\n", change.Value)
    }
}
