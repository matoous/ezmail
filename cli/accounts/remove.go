package accounts

import (
    "github.com/spf13/cobra"

    "github.com/matoous/ezmail/internal"
)

func removeAccountCmd(cfg *internal.Config) *cobra.Command {
    cmd := &cobra.Command{
        Use:     "remove <id>",
        Short:   "Remove account by ID",
        Example: "  $ ezmail accounts remove fjO3Grwl",
        Args:    cobra.ExactArgs(1),
        Run: func(cmd *cobra.Command, args []string) {
            acs := make(internal.Accounts, 0, len(cfg.Accounts)-1)
            for i := range cfg.Accounts {
                if cfg.Accounts[i].ID == args[0] {
                    err := cfg.Accounts[i].DeletePassword()
                    if err != nil {
                        cmd.PrintErr(err)
                        return
                    }
                    continue
                }
                acs = append(acs, cfg.Accounts[i])
            }
            cfg.Accounts = acs
            err := internal.SaveConfig(cfg)
            if err != nil {
                cmd.PrintErr(err)
                return
            }
        },
    }

    return cmd
}
