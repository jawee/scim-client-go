package scimapi

type PostUserRequest struct {
    UserName       string         `json:"userName"`
    Active         bool           `json:"active"`
    DisplayName    string         `json:"displayName"`
    Schemas        []string       `json:"schemas"`
    ExternalID     string         `json:"externalId"`
    Name           Name           `json:"name"`
    Emails         []Emails       `json:"emails"`
    EnterpriseUser EnterpriseUser `json:"urn:ietf:params:scim:schemas:extension:enterprise:2.1:User"`
}

type PatchRequest struct {
    Schemas    []string     `json:"schemas"`
    Operations []Operations `json:"Operations"`
}

type Operations struct {
    Op    string `json:"op"`
    Path  string `json:"path"`
    Value string `json:"value"`
}
