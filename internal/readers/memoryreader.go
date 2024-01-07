package readers

import "github.com/jawee/scim-client-go/internal/models"

type MemoryReader struct {}
func (m *MemoryReader) GetUsers() ([]models.User, error) {
    users := getSourceUsers()

    return users, nil
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
