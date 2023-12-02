package scimapi
type PatchRequest struct {
    Schemas    []string     `json:"schemas"`
    Operations []Operations `json:"Operations"`
}

type Operations struct {
    Op    string `json:"op"`
    Path  string `json:"path"`
    Value string `json:"value"`
}
