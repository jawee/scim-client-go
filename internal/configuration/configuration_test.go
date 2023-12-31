package configuration

import "testing"

type MockConfigurationProvider struct {
}

func (m *MockConfigurationProvider) GetConfigurationJson() ([]byte, error) {
    return []byte("{\"destinationconfig\": {\"scimapiurl\":\"asdfafijewora\", \"scimapitoken\":\"asdfafijewora\"}}"), nil
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

    if config.DestinationConfig.ScimUrl != "asdfafijewora" {
        t.Errorf("ScimUrl is not set correctly")
    }

    if config.DestinationConfig.ScimToken != "asdfafijewora" {
        t.Errorf("ScimToken is not set correctly")
    }
}
