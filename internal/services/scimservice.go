package services

import (
	"log"

	"github.com/jawee/scim-client-go/internal/models"
	"github.com/jawee/scim-client-go/internal/readers"
	scimapi "github.com/jawee/scim-client-go/internal/scim-api"
)

func makeMap(sourceUsers []models.User) map[string]models.User {
    usersMap := map[string]models.User{}
    for _, v := range sourceUsers {
        usersMap[v.UserName] = v
    }
    return usersMap
}

func ExecuteSync(reader readers.UsersReader, usersHistory map[string]models.UserHistory) {
    sourceUsers, err := reader.GetUsers()
    if err != nil {
        log.Printf("%s\n", err)
        return
    }

    sourceUsersMap := makeMap(sourceUsers)

    toHandle := []models.User{}
    toDelete := []string{}

    for _, v := range sourceUsers {
        toHandle = append(toHandle, v)
    }

    for k, v := range usersHistory {
        if _, ok := sourceUsersMap[k]; !ok {
            toDelete = append(toDelete, v.UserName)
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

    for _, userName := range toDelete {
        succ, err := scimapi.DeleteUser(userName)
        if err != nil { 
            log.Printf("Delete error for user %s: %s\n", userName, err)
        } else {
            log.Printf("Delete result for user %s: %v\n", userName, succ)
        }
    }
}
