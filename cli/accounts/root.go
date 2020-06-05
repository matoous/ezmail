package accounts

import (
    "os"

    "github.com/jedib0t/go-pretty/table"
    "github.com/spf13/cobra"

    "github.com/matoous/ezmail/internal"
)

// Command returns cobra Command for ids subcommand.
func Command(cfg *internal.Config) *cobra.Command {
    cmd := &cobra.Command{
        Use:   "accounts",
        Short: "List, add and remove accounts",
        Run: func(cmd *cobra.Command, args []string) {
            if len(cfg.Accounts) == 0 {
                cmd.PrintErr("no accounts")
                return
            }

            t := table.NewWriter()
            t.SetOutputMirror(os.Stdout)
            t.AppendHeader(table.Row{"ID", "Account", "SMTP Server", "Default"})
            var rows []table.Row
            for i := range cfg.Accounts {
                rows = append(rows, []interface{}{
                    cfg.Accounts[i].ID,
                    cfg.Accounts[i].Username,
                    cfg.Accounts[i].SMTPServer,
                    cfg.Accounts[i].IsDefault,
                })
            }
            t.AppendRows(rows)
            t.Render()
        },
    }

    cmd.AddCommand(addAccountCmd(cfg), removeAccountCmd(cfg))

    return cmd
}
