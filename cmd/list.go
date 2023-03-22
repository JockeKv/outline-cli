package cmd

import (
	"fmt"
	"outline/pkg/api"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List documents",
	Long:  `List documents`,
	Run: func(_ *cobra.Command, _ []string) {

		if err := conf.Read(); err != nil {
			fmt.Println("No config found, login first.")
			return
		}
		client := api.Client{Config: &conf}
		// err := conf.Login()
		// if err != nil {
		// 	fmt.Printf("Failed to log in: %v", err)
		// }
		list, err := client.Documents().List("", nil, nil)
		if err != nil {
			fmt.Println(err)
		}
		for _, doc := range *list {
			fmt.Println(*doc.Title + "\t" + *doc.Id)

		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
