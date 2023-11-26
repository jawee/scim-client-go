package configuration

type ConfigurationProvider interface {
    GetConfigurationJson() ([]byte, error)
}

