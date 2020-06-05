package input

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

// Prompt prompts the user for input.
func Prompt(prompt string) (string, error) {
    reader := bufio.NewReader(os.Stdin)

    fmt.Printf(prompt)
    data, err := reader.ReadString('\n')
    if err != nil {
        return "", err
    }
    return strings.TrimSpace(data), nil
}
