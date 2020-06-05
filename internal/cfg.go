package internal

import (
    "encoding/json"
    "errors"
    "io/ioutil"
    "os"
    "os/user"
    "path"
)

const (
    // ServiceName is used as the service name when saving passwords into the OS keychain.
    ServiceName = "ezmail"
    // AccessGroup is used as the access group when saving passwords into the OS keychain.
    AccessGroup = "1.dzx.cz.me"
)

// These variables should be initialized by the linker. They should not be initialized in code.
var (
    // BuildRev is used for storing SHA of the repo revision.
    BuildRev string
    // CommitDate is used for storing commit date of the repo revision.
    CommitDate string
)

// Config is the configuration that will be stored on users computer and persisted between command runs.
type Config struct {
    Accounts Accounts `json:"accounts"`
}

// LoadConfig loads the configuration from the configuration file.
func LoadConfig() (*Config, error) {
    confPath, err := configPath()
    if err != nil {
        return nil, err
    }

    data, err := ioutil.ReadFile(confPath)
    // probably first run or user deleted the config file
    if errors.Is(err, os.ErrNotExist) {
        return &Config{}, nil
    }
    if err != nil {
        return nil, err
    }

    var cfg Config
    err = json.Unmarshal(data, &cfg)
    if err != nil {
        return nil, err
    }
    return &cfg, nil
}

// SaveConfig saves the configuration.
func SaveConfig(c *Config) error {
    jsonC, err := json.Marshal(c)
    if err != nil {
        return err
    }
    confPath, err := configPath()
    if err != nil {
        return err
    }
    return ioutil.WriteFile(confPath, jsonC, os.ModePerm)
}

// configPath returns the path to the configuration file.
func configPath() (string, error) {
    cfgFile := ".ezmail"
    usr, err := user.Current()
    if err != nil {
        return "", err
    }
    return path.Join(usr.HomeDir, cfgFile), nil
}
