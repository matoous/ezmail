package cli

import (
	"fmt"
	"io/ioutil"
	"net/smtp"
	"os"
	"strings"

	"github.com/jaytaylor/html2text"
	"github.com/jordan-wright/email"
	"github.com/spf13/cobra"

	"github.com/matoous/ezmail/cli/accounts"
	"github.com/matoous/ezmail/internal"
)

type CLI struct {
	RootCmd *cobra.Command
	Config  *internal.Config
}

func rootCmd(cfg *internal.Config) *cobra.Command {
	var subject string
	var account string
	var html bool
	var cc []string
	var bcc []string

	cmd := &cobra.Command{
		Use:   "ezmail",
		Short: "Send email from terminal with eas",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			if len(cfg.Accounts) == 0 {
				cmd.PrintErr("no accounts")
				return
			}

			var ac *internal.Account
			if account == "" {
				ac = cfg.Accounts.Default()
			} else {
				ac = cfg.Accounts.ByID(account)
			}

			// get password for the selected user from OS keychain
			pw, err := ac.Password()
			if err != nil {
				cmd.PrintErr(err)
				return
			}

			from := ac.Username
			to := args

			in, err := ioutil.ReadAll(os.Stdin)
			if err != nil {
				cmd.PrintErr(err)
				return
			}

			e := email.NewEmail()
			e.From = from
			e.To = to
			e.Subject = subject
			e.Text = in
			if html {
				text, _ := html2text.FromString(string(in), html2text.Options{PrettyTables: true})
				e.Text = []byte(text)
				e.HTML = in
			}

			auth := smtp.PlainAuth("", ac.Username, string(pw), strings.Split(ac.SMTPServer, ":")[0])
			err = e.Send(ac.SMTPServer, auth)
			if err != nil {
				cmd.PrintErrf("sending email: %s", err)
				return
			}
			fmt.Println("📧 email successfully sent")
		},
	}

	cmd.Flags().StringVarP(&subject, "subject", "s", "", "email subject")
	cmd.Flags().StringVarP(&account, "account", "a", "", "account to use for sending, leave empty for default account")
	cmd.Flags().StringSliceVarP(&cc, "CC", "c", []string{}, "recipients in CC")
	cmd.Flags().StringSliceVarP(&bcc, "BCC", "b", []string{}, "recipients in BCC")
	cmd.Flags().BoolVar(&html, "html", false, "marks the input as html, ezmail will also try to convert it to text")

	return cmd
}

func New() (*CLI, error) {
	cfg, err := internal.LoadConfig()
	if err != nil {
		panic(err)
	}

	cli := &CLI{
		Config:  cfg,
		RootCmd: rootCmd(cfg),
	}

	cli.RootCmd.AddCommand(&cobra.Command{
		Use:   "version",
		Short: "Print the version number of ezmail",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("Ezmail mail sending utility %s (%s)\n", internal.BuildRev, internal.CommitDate)
		},
	})

	cli.RootCmd.AddCommand(
		accounts.Command(cfg),
	)

	return cli, nil
}

func (c *CLI) Execute() {
	if err := c.RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
