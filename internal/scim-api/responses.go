package scimapi

import "time"

type getUsersResponse struct {
    Schemas      []string    `json:"schemas,omitempty"`
    TotalResults int         `json:"totalResults,omitempty"`
    ItemsPerPage int         `json:"itemsPerPage,omitempty"`
    StartIndex   int         `json:"startIndex,omitempty"`
    Resources    []scimUser `json:"Resources,omitempty"`
}

type scimEmail struct {
    Type    string `json:"type,omitempty"`
    Primary bool   `json:"primary,omitempty"`
    Value   string `json:"value,omitempty"`
}

type scimPhoneNumber struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

type scimMeta struct {
    ResourceType string    `json:"resourceType,omitempty"`
    Created      time.Time `json:"created,omitempty"`
    LastModified time.Time `json:"lastModified,omitempty"`
}

type scimName struct {
    Formatted  string `json:"formatted,omitempty"`
    FamilyName string `json:"familyName,omitempty"`
    GivenName  string `json:"givenName,omitempty"`
}

type scimEnterpriseUser struct {
    Manager    scimManager `json:"manager,omitempty"`
    Department string  `json:"department,omitempty"`
}

type scimUser struct {
    EnterpriseUser scimEnterpriseUser `json:"urn:ietf:params:scim:schemas:extension:enterprise:2.0:User,omitempty"`
    Active         bool           `json:"active,omitempty"`
    DisplayName    string         `json:"displayName,omitempty"`
    Emails         []scimEmail        `json:"emails,omitempty"`
    Meta           scimMeta           `json:"meta,omitempty"`
    Name           scimName           `json:"name,omitempty"`
    UserName       string         `json:"userName,omitempty"`
    ExternalId     string         `json:"externalId,omitempty"`
    Id             string         `json:"id,omitempty"`
    PhoneNumbers   []scimPhoneNumber  `json;:"phonenumbers,omitempty"`
    Schemas        []string       `json:"schemas,omitempty"`
}

type scimPostResponse struct {
    EnterpriseUser scimEnterpriseUser `json:"urn:ietf:params:scim:schemas:extension:enterprise:2.0:User,omitempty"`
    Active         bool           `json:"active,omitempty"`
    DisplayName    string         `json:"displayName,omitempty"`
    Emails         []scimEmail       `json:"emails,omitempty"`
    Meta           scimMeta           `json:"meta,omitempty"`
    Name           scimName           `json:"name,omitempty"`
    UserName       string         `json:"userName,omitempty"`
    ExternalId     string         `json:"externalId,omitempty"`
    Id             string         `json:"id,omitempty"`
    Schemas        []string       `json:"schemas,omitempty"`
}

type scimManager struct {
    Value string `json:"value,omitempty"`
}

