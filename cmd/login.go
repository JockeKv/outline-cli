package cmd

import (
	"fmt"
	"outline/pkg/api"
	"outline/pkg/config"

	"github.com/spf13/cobra"
)

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Login to an HedgeDoc account",
	Long: `For initial use you must specify the server with --server and token with --token.
If you just want to make a login request again, you can omit these flags.`,
	Run: func(cmd *cobra.Command, _ []string) {

		if err := conf.Read(); err != nil {
			fmt.Println("No config found, creating default config..")
			if cmd.Flag("server").Value.String() == "" {
				fmt.Println("You need to set a server on initial run.")
				return
			}
			if cmd.Flag("token").Value.String() == "" {
				fmt.Println("You need to set a token on initial run.")
				return
			}
			conf = config.Config{
				Host:  cmd.Flag("server").Value.String(),
				Token: cmd.Flag("token").Value.String(),
			}

			err = conf.Write()
			if err != nil {
				fmt.Printf("Failed to save configuration: %v", err)
				return
			}
		}

		client := api.Client{Config: &conf}
		err := client.Login()
		if err != nil {
			fmt.Printf("Failed to login: %v", err)
			return
		}
		fmt.Println("Login successful!")
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)

	loginCmd.Flags().StringP("server", "s", "", "Set the server URL")
	loginCmd.Flags().StringP("token", "c", "", "Access token")
}
