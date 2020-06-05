package accounts

import (
    "fmt"

    "github.com/spf13/cobra"

    "github.com/matoous/ezmail/internal"
    "github.com/matoous/ezmail/internal/input"
)

func addAccountCmd(cfg *internal.Config) *cobra.Command {
    cmd := &cobra.Command{
        Use:     "add",
        Short:   "Add email account",
        Example: "  $ ezmail accounts add",
        Run: func(cmd *cobra.Command, args []string) {
            username, err := input.Prompt("Your email:")
            if err != nil {
                fmt.Printf("ERR: get username: %s", err)
                return
            }

            pass, err := input.PasswordPrompt("Your password:")
            if err != nil {
                fmt.Printf("ERR: get username: %s", err)
                return
            }

            server, err := input.Prompt("SMTP server [smtp.gmail.com:587]:")
            if err != nil {
                fmt.Printf("ERR: get username: %s", err)
                return
            }
            if server == "" {
                server = "smtp.gmail.com:587"
            }

            isDef := false
            if len(cfg.Accounts) == 0 {
                isDef = true
            }

            acc := internal.NewAccount(username, server, isDef)
            if err := acc.SavePassword(pass); err != nil {
                fmt.Printf("ERR: saving password for account: %s", err)
                return
            }
            cfg.Accounts = append(cfg.Accounts, acc)
            err = internal.SaveConfig(cfg)
            if err != nil {
                cmd.PrintErr(err)
                return
            }
        },
    }

    return cmd
}
