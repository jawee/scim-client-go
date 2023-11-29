package scimapi

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

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

func getExistingUser(token, userName string) (*models.User, error) {
    url := fmt.Sprintf("%s/users/?filter=userName+eq+%s", API_URL, userName)

    body, err := makeRequest(token, url, http.MethodGet)
    if err != nil {
        return nil, err
    }

    log.Printf("getExistingUser: %s\n", string(body))

    var userResponse GetUsersResponse
    err = json.Unmarshal(body, &userResponse)
    if err != nil {
        return nil, err
    }
    return nil, nil
}

func makeRequest(token, url, method string)  ([]byte, error) {
    var body io.Reader
    request, err := http.NewRequest(method, url, body)
    if err != nil {
        log.Printf("Error: %s\n", err)
        return nil, err
    }
    request.Header.Set("Authorization", "Bearer " + token)

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

func HandleUser(newUser *models.User, oldUser *models.User) (ExternalId, error) {
    token, err := getToken()
    if err != nil {
        log.Printf("Error: %s\n", err)
        return ERROR_EXTERNAL_ID, err
    }
    // log.Printf("Got token: '%s'\n", token)

    _, err = getExistingUser(token, newUser.UserName)
    if err != nil {
        log.Printf("Error: %s\n", err)
        return ERROR_EXTERNAL_ID, err
    }

    // log.Printf("ExistingUser:\n %s\n", structAsString(user))

    // printExistingUsers(token)
    return ERROR_EXTERNAL_ID, fmt.Errorf("error");
}

func printExistingUsers(token string) {
    requestURL := fmt.Sprintf("%s/users", API_URL)

    body, err := makeRequest(token, requestURL, http.MethodGet)
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
