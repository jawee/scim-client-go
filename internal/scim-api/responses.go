package scimapi

import "time"

type GetUsersResponse struct {
    Schemas      []string    `json:"schemas,omitempty"`
    TotalResults int         `json:"totalResults,omitempty"`
    ItemsPerPage int         `json:"itemsPerPage,omitempty"`
    StartIndex   int         `json:"startIndex,omitempty"`
    Resources    []User `json:"Resources,omitempty"`
}

type Email struct {
    Type    string `json:"type,omitempty"`
    Primary bool   `json:"primary,omitempty"`
    Value   string `json:"value,omitempty"`
}

type Meta struct {
    ResourceType string    `json:"resourceType,omitempty"`
    Created      time.Time `json:"created,omitempty"`
    LastModified time.Time `json:"lastModified,omitempty"`
}

type Name struct {
    Formatted  string `json:"formatted,omitempty"`
    FamilyName string `json:"familyName,omitempty"`
    GivenName  string `json:"givenName,omitempty"`
}

type EnterpriseUser struct {
    Manager    Manager `json:"manager,omitempty"`
    Department string  `json:"department,omitempty"`
}

type User struct {
    EnterpriseUser EnterpriseUser `json:"urn:ietf:params:scim:schemas:extension:enterprise:2.1:User,omitempty"`
    Active         bool           `json:"active,omitempty"`
    DisplayName    string         `json:"displayName,omitempty"`
    Emails         []Email        `json:"emails,omitempty"`
    Meta           Meta           `json:"meta,omitempty"`
    Name           Name           `json:"name,omitempty"`
    UserName       string         `json:"userName,omitempty"`
    ExternalID     string         `json:"externalId,omitempty"`
    ID             string         `json:"id,omitempty"`
    Schemas        []string       `json:"schemas,omitempty"`
}

type PostResponse struct {
    EnterpriseUser EnterpriseUser `json:"urn:ietf:params:scim:schemas:extension:enterprise:2.1:User,omitempty"`
    Active         bool           `json:"active,omitempty"`
    DisplayName    string         `json:"displayName,omitempty"`
    Emails         []Email       `json:"emails,omitempty"`
    Meta           Meta           `json:"meta,omitempty"`
    Name           Name           `json:"name,omitempty"`
    UserName       string         `json:"userName,omitempty"`
    ExternalID     string         `json:"externalId,omitempty"`
    ID             string         `json:"id,omitempty"`
    Schemas        []string       `json:"schemas,omitempty"`
}

type Manager struct {
    Value string `json:"value,omitempty"`
}

