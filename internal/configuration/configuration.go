package configuration

import (
	"encoding/json"
	"log"
	"os"
	"path"
)

type ConfigurationProvider interface {
    GetConfigurationJson() ([]byte, error)
}

type FileConfigurationProvider struct {
}

type Config struct { 
    ScimUrl string `json:"scimapiurl"`
    ScimToken string `json:"scimapitoken"`
}

func (f *FileConfigurationProvider) GetConfigurationJson() ([]byte, error) {
    configDir := "/config"
    log.Printf("FileConfigurationProvider.GetConfigurationJson. Loading configuration from %s/config.json\n", configDir)
    path := path.Join(configDir, "config.json")
    file, err := os.Open(path)

    if err != nil {
        log.Printf("FileConfigurationProvider.GetConfigurationJson. Error opening configuration file: %s", err)
        return nil, err
    }
    bytes := make([]byte, 1024)

    readTotal, err := file.Read(bytes)

    if err != nil {
        log.Printf("FileConfigurationProvider.GetConfigurationJson. Error reading configuration file: %s", err)
        return nil, err
    }
    
    return bytes[:readTotal], nil
}


func New(configProvider ConfigurationProvider) (*Config, error) {
    log.Println("New. Loading configuration")    

    bytes, err := configProvider.GetConfigurationJson()

    if err != nil {
        log.Printf("Error getting configuration from provider: %s", err)
        return nil, err
    }

    var configuration *Config
    err = json.Unmarshal(bytes, &configuration) 
    if err != nil {  
        log.Printf("Error unmarshalling configuration: %s", err)
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
