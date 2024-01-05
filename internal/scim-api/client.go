package scimapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
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
    var email string
    for _, v := range resp.Emails {
        if v.Type == "work" {
            email = v.Value
        }
    }

    var phoneNumber string
    for _, v := range resp.PhoneNumbers {
        if v.Type == "work" {
            phoneNumber = v.Value
        }
    }
    user := models.User {
        Id: resp.ExternalId,
        UserName: resp.UserName,
        Email: email,
        PhoneNumber: phoneNumber,
        Department: resp.EnterpriseUser.Department,
        FirstName: resp.Name.GivenName,
        LastName: resp.Name.FamilyName,
        Active: resp.Active,
        ExternalId: resp.Id,
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
    phoneNumber := PhoneNumber {
        Type: "work",
        Value: user.PhoneNumber,
    }
    scimUser := User {
        EnterpriseUser: EnterpriseUser{
            Department: user.Department,
        }, 
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
        PhoneNumbers: []PhoneNumber { phoneNumber, },
        UserName: user.UserName,
        ExternalId: user.Id,
        Id: "",
        Schemas: []string{
            "urn:ietf:params:scim:schemas:extension:enterprise:2.0:User",
            "urn:ietf:params:scim:schemas:core:2.0:User",
        },
    }

    return scimUser
}

func createUser(token string, user *models.User) (ExternalId, error) {
    scimUser := userToScimUser(user)
    url := fmt.Sprintf("%s/users", API_URL)
    log.Printf("Sending: %s\n", structAsString(scimUser))
    resBytes, err := makeRequest(token, url, http.MethodPost, scimUser)
    if err != nil {
        return "", err
    }

    log.Printf("Received: %s\n", string(resBytes))
    var createdUser User
    err = json.Unmarshal(resBytes, &createdUser)
    if err != nil {
        return "", err
    }

    return ExternalId(createdUser.Id), nil
}

func HandleUser(newUser *models.User) (ExternalId, error) {
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
        externalId, err := createUser(token, newUser)
        if err != nil {
            log.Printf("ERROR: Create user failed. %s\n", err)
            return ERROR_EXTERNAL_ID, err
        }

        existingUser = newUser
        existingUser.ExternalId = string(externalId)
        return ExternalId(existingUser.ExternalId), nil;
    }

    changes, err := diffUsers(existingUser, newUser)

    patchOperations := []Operations{}
    for _, v := range changes {
        log.Printf("Changes in fieldname '%s'\n", v.FieldName)
        path := getPath(v.FieldName)
        if path == "" {
            log.Printf("No path for fieldname '%s'\n", v.FieldName)
            continue;
        }
        op := Operations{
            Op: "replace",
            Path: path,
            Value: v.Value,
        }
        patchOperations = append(patchOperations, op)
    }

    patchUser(token, existingUser.ExternalId, patchOperations)

    // log.Printf("ExistingUser:\n %s\n", structAsString(user))

    // printExistingUsers(token)
    return ExternalId(existingUser.ExternalId), nil;
}

func patchUser(token string, id string, patchOperations []Operations) {
    if len(patchOperations) == 0 {
        return;
    }

    url := fmt.Sprintf("%s/users/%s", API_URL, id)
    request := PatchRequest {
        Operations: patchOperations,
        Schemas: []string{"urn:ietf:params:scim:api:messages:2.0:PatchOp", },
    }

    
    // bytes, err := json.Marshal(request)
    // requestStr := string(bytes)
    log.Printf("Making patch request: %s\n", structAsString(request))

    res, err := makeRequest(token, url, http.MethodPatch, request)
    if err != nil {
        log.Printf("Patch failed: %s\n", err)
        return 
    }

    log.Printf("%s\n", string(res))
}

func getPath(fieldName string) string {
    s := ""
    lowerFieldName := strings.ToLower(fieldName)
    switch lowerFieldName {
        case "username": 
            s = "username"
        case "email":
            s = "emails[type eq \"work\"].Value"
        case "firstname":
            s = "name.givenName"
        case "lastname":
            s = "name.familyName"
        case "phonenumber":
            s = "phoneNumbers[type eq \"work\"].Value"
        case "active":
            s = "active"
        case "department":
            s = "urn:ietf:params:scim:schemas:extension:enterprise:2.0:User:department"
    }
    return s;
}

type changes struct {
    FieldName string
    Value string
}

func diffUsers(existingUser *models.User, newUser *models.User) ([]changes, error) {
    fields := []changes{}
    if existingUser.Email != newUser.Email {
        fields = append(fields, changes{ FieldName: "email", Value: newUser.Email })
    }

    if existingUser.FirstName != newUser.FirstName {
        fields = append(fields, changes{ FieldName: "firstname", Value: newUser.FirstName })
    }

    if existingUser.LastName != newUser.LastName {
        fields = append(fields, changes{ FieldName: "lastname", Value: newUser.LastName })
    }

    if existingUser.PhoneNumber != newUser.PhoneNumber {
        fields = append(fields, changes{ FieldName: "phonenumber", Value: newUser.PhoneNumber })
    }
    if existingUser.Active != newUser.Active {
        fields = append(fields, changes{ FieldName: "active", Value: fmt.Sprintf("%v", newUser.Active) })
    }
    if existingUser.Department != newUser.Department {
        log.Printf("%s != %s\n", existingUser.Department, newUser.Department)
        fields = append(fields, changes{ FieldName: "department", Value: newUser.Department, })
    }
    return fields, nil
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
