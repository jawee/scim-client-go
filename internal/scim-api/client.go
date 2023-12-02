package scimapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/jawee/scim-client-go/internal/models"
)

type ExternalId string
const ERROR_EXTERNAL_ID = ""

const TOKEN = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYmYiOjE2OTk2OTQzMDQsImV4cCI6MTY5OTcwMTUwNCwiaXNzIjoiTWljcm9zb2Z0LlNlY3VyaXR5LkJlYXJlciIsImF1ZCI6Ik1pY3Jvc29mdC5TZWN1cml0eS5CZWFyZXIifQ.ySF9zrSZGaIgVvdxJ4LSkYURPJfdEFkU77Q7WYhYn4I"

const API_URL = "http://localhost:6000/scim"

func getToken() (string, error) {
    tokenUrl := API_URL + "/token"

    var body io.Reader

    request, err := http.NewRequest(http.MethodGet, tokenUrl, body)
    if err != nil {
        return "", err 
    }
    
    client := http.Client{}
    resp, err := client.Do(request)
    if err != nil {
        log.Printf("Error: %s\n", err)
        return "", err
    }

    resBody, err := io.ReadAll(resp.Body)
    if err != nil { 
        log.Printf("Error: %s\n", err)
        return ERROR_EXTERNAL_ID, err
    }

	// fmt.Printf("getToken: status code: %d\n", resp.StatusCode)
 //    fmt.Printf("getToken: response: %s\n", resBody)
	//
    var getTokenResponse getTokenResponse
    err = json.Unmarshal(resBody, &getTokenResponse)
    if err != nil { 
        log.Printf("Unmarshal Error: %s\n", err)
        return ERROR_EXTERNAL_ID, err
    }

    return getTokenResponse.Token, nil;
}

type getTokenResponse struct {
	Token string `json:"token"`
}

func userResponseToUser(resp User) (models.User) {
    email := ""
    user := models.User {
        Id: resp.ExternalID,
        UserName: resp.UserName,
        Email: email,
        PhoneNumber: "",
        Department: resp.EnterpriseUser.Department,
        FirstName: resp.Name.GivenName,
        LastName: resp.Name.FamilyName,
        Active: resp.Active,
        ExternalId: resp.ID,
    }

    return user
}

func getExistingUser(token, userName string) (*models.User, error) {
    url := fmt.Sprintf("%s/users/?filter=userName+eq+%s", API_URL, userName)

    body, err := makeRequest(token, url, http.MethodGet, nil)
    if err != nil {
        return nil, err
    }

    var usersResponse GetUsersResponse
    err = json.Unmarshal(body, &usersResponse)
    if err != nil {
        return nil, err
    }

    if usersResponse.TotalResults == 0 {
        return nil, nil
    }

    if usersResponse.TotalResults > 1 {
        return nil, fmt.Errorf("userName %s returned more than 1 result", userName)
    }

    existingUser := userResponseToUser(usersResponse.Resources[0])

    return &existingUser, nil
}

func makeRequest(token, url, method string, content interface{})  ([]byte, error) {
    var body io.Reader

    if content != nil {
        contentBytes, err := json.Marshal(content)
        if err != nil {
            return nil, err
        }

        body = bytes.NewBuffer(contentBytes)
    }

    request, err := http.NewRequest(method, url, body)
    if err != nil {
        log.Printf("Error: %s\n", err)
        return nil, err
    }
    request.Header.Set("Authorization", "Bearer " + token)
    request.Header.Set("Content-Type", "application/json")

    client := http.Client{ }

    resp, err := client.Do(request)
    if err != nil {
        log.Printf("Error: %s\n", err)
        return nil, err
    }

    resBody, err := io.ReadAll(resp.Body)
    if err != nil { 
        log.Printf("Error: %s\n", err)
        return nil, err
    }

    return resBody, nil;
}

func userToScimUser(user *models.User) User {
    email := Email {
        "work", true, user.Email,
    }
    scimUser := User {
        EnterpriseUser: EnterpriseUser{}, 
        Active: user.Active,
        DisplayName: user.UserName,
        Emails: []Email{ email },
        Meta: Meta{
            ResourceType: "User",
            Created: time.Now(),
            LastModified: time.Now(),
        },
        Name: Name{
            Formatted: fmt.Sprintf("%s %s", user.FirstName, user.LastName),
            FamilyName: user.LastName,
            GivenName: user.FirstName,
        },
        UserName: user.UserName,
        ExternalID: user.Id,
        ID: "",
        Schemas: []string{
            "urn:ietf:params:scim:schemas:extension:enterprise:2.0:User",
            "urn:ietf:params:scim:schemas:core:2.0:User",
        },
    }

    return scimUser
}

func createUser(token string, user *models.User) (string, error) {
    scimUser := userToScimUser(user)
    url := fmt.Sprintf("%s/users", API_URL)
    resBytes, err := makeRequest(token, url, http.MethodPost, scimUser)
    if err != nil {
        return "", err
    }

    var createdUser User
    err = json.Unmarshal(resBytes, &createdUser)
    if err != nil {
        return "", err
    }

    return user.Id, nil
}

func HandleUser(newUser *models.User, oldUser *models.User) (ExternalId, error) {
    token, err := getToken()
    if err != nil {
        log.Printf("Error: %s\n", err)
        return ERROR_EXTERNAL_ID, err
    }

    existingUser, err := getExistingUser(token, newUser.UserName)
    if err != nil {
        log.Printf("Error: %s\n", err)
        return ERROR_EXTERNAL_ID, err
    }


    if existingUser == nil {
        log.Printf("User doesn't exist. Creating. \n")
        _, err = createUser(token, newUser)
        if err != nil {
            log.Printf("ERROR: Create user failed. %s\n", err)
            return ERROR_EXTERNAL_ID, err
        }
    }


    // log.Printf("ExistingUser:\n %s\n", structAsString(user))

    // printExistingUsers(token)
    return ERROR_EXTERNAL_ID, fmt.Errorf("error");
}

func printExistingUsers(token string) {
    requestURL := fmt.Sprintf("%s/users", API_URL)

    body, err := makeRequest(token, requestURL, http.MethodGet, nil)
    if err != nil {
        log.Printf("Error: %s\n", err)
        return
    }
    var getUsersResp GetUsersResponse
    err = json.Unmarshal(body, &getUsersResp)
    if err != nil {
        log.Printf("Error: %s\n", err)
        return
    }

    fmt.Printf("GotUsers:\n %s\n", structAsString(getUsersResp))
}

func structAsString(model any) string {
    empJSON, err := json.MarshalIndent(model, "", " ")
    if err != nil {
        return "Something went wrong in structAsString"
    }
    str := fmt.Sprintf("%s", string(empJSON))
    return str
}
