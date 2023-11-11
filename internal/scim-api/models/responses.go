package models

import "time"

type GetUsersResponse struct {
    Schemas      []string    `json:"schemas"`
    TotalResults int         `json:"totalResults"`
    ItemsPerPage int         `json:"itemsPerPage"`
    StartIndex   int         `json:"startIndex"`
    Resources    []Resources `json:"Resources"`
}

type Emails struct {
    Type    string `json:"type"`
    Primary bool   `json:"primary"`
    Value   string `json:"value"`
}

type Meta struct {
    ResourceType string    `json:"resourceType"`
    Created      time.Time `json:"created"`
    LastModified time.Time `json:"lastModified"`
}

type Name struct {
    Formatted  string `json:"formatted"`
    FamilyName string `json:"familyName"`
    GivenName  string `json:"givenName"`
}

type EnterpriseUser struct {
    Manager    Manager `json:"manager"`
    Department string  `json:"department"`
}

type Resources struct {
    EnterpriseUser EnterpriseUser `json:"urn:ietf:params:scim:schemas:extension:enterprise:2.1:User"`
    Active         bool           `json:"active"`
    DisplayName    string         `json:"displayName"`
    Emails         []Emails       `json:"emails"`
    Meta           Meta           `json:"meta"`
    Name           Name           `json:"name"`
    UserName       string         `json:"userName"`
    ExternalID     string         `json:"externalId"`
    ID             string         `json:"id"`
    Schemas        []string       `json:"schemas"`
}

type PostResponse struct {
    EnterpriseUser EnterpriseUser `json:"urn:ietf:params:scim:schemas:extension:enterprise:2.1:User"`
    Active         bool           `json:"active"`
    DisplayName    string         `json:"displayName"`
    Emails         []Emails       `json:"emails"`
    Meta           Meta           `json:"meta"`
    Name           Name           `json:"name"`
    UserName       string         `json:"userName"`
    ExternalID     string         `json:"externalId"`
    ID             string         `json:"id"`
    Schemas        []string       `json:"schemas"`
}

type Manager struct {
    Value string `json:"value"`
}

