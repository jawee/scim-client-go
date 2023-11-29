package readers

import "github.com/jawee/scim-client-go/internal/models"

type UsersReader interface {
    GetUsers() ([]models.User, error)
}
