package readers

import (
	"bytes"
	"testing"
)

func TestIdColumnMissing(t *testing.T) {
    ibuf := bytes.NewBufferString(`Email,FirstName,LastName,MobilePhone,Department`)
    _, err := ReadFile(ibuf)
    if err == nil {
        t.Fatalf("Expected error, got %v\n", err)
    }
}

func TestReadser(t *testing.T) {
    ibuf := bytes.NewBufferString(`Id,Email,FirstName,LastName,MobilePhone,Department
    1,some.user@compay.com,some,user,12345678,Sales`)
    users, err := ReadFile(ibuf)
    if err != nil {
        t.Fatalf("Expected no error, got %v\n", err)
    }

    if len(users) != 1 {
        t.Fatalf("Expected length 2, got %d\n", len(users))
    }

    user := users[0]

    if user.Id != "1" {
        t.Errorf("Expected Id to be '1', got '%s'\n", user.Id)
    }
    if user.UserName != "some.user@compay.com" {
        t.Errorf("Expected UserName to be 'some.user@compay.com', got '%s'\n", user.UserName)
    }
    if user.Email != "some.user@compay.com" {
        t.Errorf("Expected Email to be 'some.user@compay.com', got '%s'\n", user.Email)
    }
    if user.FirstName != "some" {
        t.Errorf("Expected FirstName to be 'some', got '%s'\n", user.FirstName)
    }
    if user.LastName != "user" {
        t.Errorf("Expected LastName to be 'user', got '%s'\n", user.LastName)
    }
    if user.PhoneNumber != "12345678" {
        t.Errorf("Expected PhoneNumber to be '12345678', got '%s'\n", user.PhoneNumber)
    }
    if user.Department != "Sales" {
        t.Errorf("Expected Department to be 'Sales', got '%s'\n", user.Department)
    }
}

func TestReadTwoUsers(t *testing.T) {
    ibuf := bytes.NewBufferString(`Id,Email,FirstName,LastName,MobilePhone,Department
    1,some.user@compay.com,some,user,12345678,Sales 
    2,another.user@compay.com,another,user,87654321,Sales`)
    users, err := ReadFile(ibuf)
    if err != nil {
        t.Fatalf("Expected no error, got %v\n", err)
    }

    if len(users) != 2 {
        t.Fatalf("Expected length 2, got %d\n", len(users))
    }
}


