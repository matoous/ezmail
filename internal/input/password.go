package input

import (
    "fmt"
    "syscall"

    "golang.org/x/crypto/ssh/terminal"
)

// PasswordPrompt prompts the user for password obfuscating the input.
func PasswordPrompt(prompt string) ([]byte, error) {
    fmt.Print(prompt)
    pw, err := terminal.ReadPassword(syscall.Stdin)
    if err != nil {
        return nil, err
    }
    fmt.Println()
    return pw, nil
}
