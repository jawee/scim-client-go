package main

import (
	"os"
	"path"
	"path/filepath"
	"strings"
	"testing"

	"github.com/jawee/scim-client-go/internal/flags"
)

func TestGetConfigPathRelative(t *testing.T) {
    configPath := "./config"

    expectedPath := ""
    ex, _ := os.Executable()
    exPath := filepath.Dir(ex)
    exPath, _ = strings.CutSuffix(exPath, "/")

    cust, foundPrefix := strings.CutPrefix(configPath, "./"); 
    if foundPrefix {
        expectedPath = path.Join(exPath, cust)
    }

    path, err := getConfigPath([]flags.Flag{ { Type: flags.ConfigDir, Value: configPath }})
    if err != nil {
        t.Fatalf("Err: '%s'\n", err)
    }

    if path != expectedPath {
        t.Fatalf("Got '%s', expected '%s'\n", path, expectedPath)
    }
}

func TestGetConfigPathAbsolute(t *testing.T) {
    configPath := "/config"

    path, err := getConfigPath([]flags.Flag{ { Type: flags.ConfigDir, Value: configPath }})
    if err != nil {
        t.Fatalf("Err: '%s'\n", err)
    }

    if path != configPath {
        t.Fatalf("Got '%s', expected '%s'\n", path, configPath)
    }
}

func TestGetConfigPathDefault(t *testing.T) {
    defaultPath, err := os.UserConfigDir()
    defaultPath = path.Join(defaultPath, "/scimclient")

    path, err := getConfigPath([]flags.Flag{})
    if err != nil {
        t.Fatalf("Err: '%s'\n", err)
    }

    if path != defaultPath {
        t.Fatalf("Got '%s', expected '%s'\n", path, "./config")
    }
}
