package readers

import "github.com/jawee/scim-client-go/internal/models"

type ScimReader interface {
    GetUsers() ([]models.User, error)
}
