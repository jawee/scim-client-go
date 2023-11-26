package configuration

import (
	"fmt"
	"log"
	"os"
	"path"
)

type FileConfigurationProvider struct {
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


