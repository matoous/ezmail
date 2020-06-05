package internal

import (
    "github.com/keybase/go-keychain"
    "github.com/matoous/go-nanoid"
)

// Account is single email account entry that will be saved in the configuration file.
type Account struct {
    ID         string `json:"id"`
    Username   string `json:"username"`
    SMTPServer string `json:"smtp_server"`
    IsDefault  bool   `json:"is_default"`
}

// NewAccount creates new account.
func NewAccount(username, server string, isDefault bool) *Account {
    id, _ := gonanoid.ID(8)
    return &Account{
        ID:         id,
        Username:   username,
        SMTPServer: server,
        IsDefault:  isDefault,
    }
}

// SavePassword saves password for given account into the OS keychain.
func (ac *Account) SavePassword(pw []byte) error {
    item := keychain.NewGenericPassword(ServiceName, ac.Username, "", pw, AccessGroup)
    item.SetSynchronizable(keychain.SynchronizableNo)
    item.SetAccessible(keychain.AccessibleWhenUnlocked)
    err := keychain.AddItem(item)
    if err == keychain.ErrorDuplicateItem {
        return err
    }
    return nil
}

// Password retrieves the accounts password from the OS keychain.
func (ac *Account) Password() ([]byte, error) {
    pw, err := keychain.GetGenericPassword(ServiceName, ac.Username, "", AccessGroup)
    if err != nil {
        return nil, err
    }
    return pw, nil
}

// DeletePassword deletes the password entry for given account from the OS keychain.
func (ac *Account) DeletePassword() error {
    return keychain.DeleteGenericPasswordItem(ServiceName, ac.Username)
}
