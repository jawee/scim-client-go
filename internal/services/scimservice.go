package services

import (
	"log"
	"time"

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

func ExecuteSync(reader readers.UsersReader, usersHistory map[string]models.UserHistory) ([]models.UserHistory) {
    sourceUsers, err := reader.GetUsers()
    if err != nil {
        log.Printf("%s\n", err)
        return nil
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

    result := []models.UserHistory{}
    for _, user := range toHandle {
        hist := createHistory(user.UserName)
        id, err := scimapi.HandleUser(&user)

        if err != nil {
            log.Printf("%s\n", err)
            hist.ErrorMessage = err.Error()
        }
        log.Printf("ExternalId: %s\n", id)
        result = append(result, hist)
    }

    for _, userName := range toDelete {
        succ, err := scimapi.DeleteUser(userName)
        if err != nil { 
            log.Printf("Delete error for user %s: %s\n", userName, err)
            hist := createHistory(userName)
            hist.ErrorMessage = err.Error()
            result = append(result, hist)
        } else {
            log.Printf("Delete result for user %s: %v\n", userName, succ)
        }
    }

    return result
}

func createHistory(userName string) models.UserHistory {
    hist := models.UserHistory {
        UserName: userName,
        LastSync: time.Now().UTC(),
    }

    return hist
}
