package configuration

import (
	"encoding/json"
	"fmt"
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

func findConfigurationFile() (*os.File, error) {
    // userCfgDir, err := os.UserConfigDir();
    cwd, err := os.Getwd()
    if err != nil {
        log.Printf("FileConfigurationProvider.findConfigurationFile. Couldn't get cwd\n")
        return nil, err
    }

    configDir := "/configs"
    path := path.Join(cwd, configDir, "config.json")
    log.Printf("FileConfigurationProvider.findConfigurationFile. Checking if %s exists\n", path)

    file, err := os.Open(path)
    if err != nil {
        log.Printf("FileConfigurationProvider.findConfigurationFile. Error opening configuration file: %s\n", err)
        return nil, err
    }

    log.Printf("FileConfigurationProvider.findConfigurationFile. File found at %s.\n", path)
    return file, nil
}

func (f *FileConfigurationProvider) GetConfigurationJson() ([]byte, error) {
    file, err := findConfigurationFile()
    if err != nil {
        log.Printf("Couldn't find configuration file. %s\n", err)
        return nil, fmt.Errorf("Couldn't find configuration file. %s\n", err)
    }

    bytes := make([]byte, 1024)

    readTotal, err := file.Read(bytes)

    if err != nil {
        log.Printf("FileConfigurationProvider.GetConfigurationJson. Error reading configuration file: %s\n", err)
        return nil, err
    }
    
    return bytes[:readTotal], nil
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
