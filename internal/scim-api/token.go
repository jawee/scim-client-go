package scimapi

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

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
