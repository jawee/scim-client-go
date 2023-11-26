package configuration

import (
	"encoding/json"
	"log"
)

type DestinationConfig struct {
    ScimUrl string `json:"scimapiurl"`
    ScimToken string `json:"scimapitoken"`
}

type FileReaderConfig struct {
    FilePath string `json:"filepath"`
}

type Config struct { 
    DestinationConfig DestinationConfig `json:"destinationconfig,omitempty"`
    FileReaderConfig FileReaderConfig `json:"filereaderconfig"`
}

func New(configProvider ConfigurationProvider) (*Config, error) {
    log.Println("New. Loading configuration")    

    bytes, err := configProvider.GetConfigurationJson()

    if err != nil {
        log.Printf("Error getting configuration from provider: %s\n", err)
        return nil, err
    }

    var configuration *Config
    err = json.Unmarshal(bytes, &configuration) 
    if err != nil {  
        log.Printf("Error unmarshalling configuration: %s\n", err)
        return nil, err
    }

    // if configuration.ClientId == "" || configuration.ClientSecret == "" {
    //     log.Println("ClientId or ClientSecret is empty")
    //     return nil, err
    // }

    // streamers := strings.Replace(configuration.StreamersString, " ", "", -1)
    // configuration.Streamers = strings.Split(streamers, ",")

    return configuration, nil
}
