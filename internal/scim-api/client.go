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

func HandleUser(newUser *models.User, oldUser *models.User) (ExternalId, error) {
    token, err := getToken()
    if err != nil {
        log.Printf("Error: %s\n", err)
        return ERROR_EXTERNAL_ID, err
    }
    log.Printf("Got token: '%s'\n", token)
    // requestURL := fmt.Sprintf("http://DESKTOP-QQQIEJC.local:59629/scim/users")
    // requestURL := fmt.Sprintf("http://localhost:5000/scim/token")
    requestURL := fmt.Sprintf("%s/users", API_URL)
	// resp, err := http.Get(requestURL)
 //    if err != nil {
 //        log.Printf("Error: %s\n", err)
 //        return ERROR_EXTERNAL_ID, err
 //    }
    
    var body io.Reader
    request, err := http.NewRequest(http.MethodGet, requestURL, body)
    if err != nil {
        log.Printf("Error: %s\n", err)
        return ERROR_EXTERNAL_ID, err
    }
    request.Header.Set("Authorization", "Bearer " + token)

    client := http.Client{ }

    resp, err := client.Do(request)
    if err != nil {
        log.Printf("Error: %s\n", err)
        return ERROR_EXTERNAL_ID, err
    }

    resBody, err := io.ReadAll(resp.Body)
    if err != nil { log.Printf("Error: %s\n", err)
        return ERROR_EXTERNAL_ID, err
    }

    fmt.Printf("client: got response!\n")
	fmt.Printf("client: status code: %d\n", resp.StatusCode)
    fmt.Printf("client: response: %s\n", resBody)

    var getUsersResp GetUsersResponse
    err = json.Unmarshal(resBody, &getUsersResp)
    if err != nil { log.Printf("Error: %s\n", err)
        return ERROR_EXTERNAL_ID, err
    }

    fmt.Printf("json: %v\n", getUsersResp)
    // var request = http.NewRequest("GET", "http://localhost:59628/scim/users")
    return ERROR_EXTERNAL_ID, fmt.Errorf("error");
}
