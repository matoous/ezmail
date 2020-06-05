package internal

// Accounts is list of accounts that provides additional functionality over plain slice.
type Accounts []*Account

// ByID retrieves account by its ID.
func (acs Accounts) ByID(id string) *Account {
    for i := range acs {
        if acs[i].ID == id {
            return acs[i]
        }
    }
    return nil
}

// Default retrieves the default account.
func (acs Accounts) Default() *Account {
    for i := range acs {
        if acs[i].IsDefault {
            return acs[i]
        }
    }
    return nil
}
