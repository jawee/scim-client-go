package configuration

import "testing"

type MockConfigurationProvider struct {
}

func (m *MockConfigurationProvider) GetConfigurationJson() ([]byte, error) {
    return []byte("{\"scimapiurl\":\"asdfafijewora\", \"scimtoken\":\"asdfafijewora\"}"), nil
}

func TestNewConfiguration(t *testing.T) {

    provider := new(MockConfigurationProvider)

    config, err := New(provider)

    if err != nil {
        t.Errorf("Error creating configuration: %s", err)
    }
    if config == nil {
        t.Errorf("Configuration is nil")
    }

    if config.ScimUrl != "asdfafijewora" {
        t.Errorf("ClientId is not set correctly")
    }

    if config.ScimToken != "asdfafijewora" {
        t.Errorf("ClientSecret is not set correctly")
    }
}
