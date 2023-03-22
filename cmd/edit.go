package cmd

import (
	"fmt"
	"outline/pkg/api"
	"outline/pkg/editor"
	"strings"
	"unicode"

	"github.com/spf13/cobra"
	"golang.org/x/text/runes"
)

var editCmd = &cobra.Command{
	Use:   "edit",
	Short: "Edit a document",
	Long:  `Edit a document`,
	Run: func(_ *cobra.Command, args []string) {
		if len(args) < 2 {
			fmt.Println("Provide ID for document to edit.")
			return
		}
		if err := conf.Read(); err != nil {
			fmt.Println("No config found, login first.")
			return
		}
		client := api.Client{Config: &conf}
		err := client.Login()
		if err != nil {
			fmt.Printf("Failed to log in: %v", err)
		}

		doc, err := client.Documents().Info(args[1])
		if err != nil {
			fmt.Printf("Could not get document: %v", err)
			return
		}

		res, err := editor.EditDoc(doc)
		if err != nil {
			fmt.Printf("Could not edit document: %v", err)
		}

		if res != nil {
			_, err := client.Documents().Update(doc)
			if err != nil {
				fmt.Printf("Could not update document: %v", err)
			}
		}
	},
	ValidArgsFunction: completeDocs,
}

func init() {
	rootCmd.AddCommand(editCmd)
}

func completeDocs(_ *cobra.Command, args []string, _ string) ([]string, cobra.ShellCompDirective) {
	if err := conf.Read(); err != nil {
		return []string{"error"}, cobra.ShellCompDirectiveDefault
	}
	client := api.Client{Config: &conf}
	cols, err := client.Collections().List()
	if err != nil {
		return []string{"error"}, cobra.ShellCompDirectiveDefault
	}
	colMap := map[string]string{}
	for _, col := range *cols {
		colMap[*col.Name] = *col.Id
	}
	if len(args) == 0 {
		argsList := []string{""}
		for k := range colMap {
			argsList = append(argsList, k)
		}
		return argsList, cobra.ShellCompDirectiveDefault
	} else if len(args) == 1 {
		col := colMap[args[0]]
		docs, err := client.Documents().List(col, nil, nil)
		if err != nil {
			return []string{"error"}, cobra.ShellCompDirectiveDefault
		}
		argsList := []string{}
		for _, doc := range *docs {
			// name := strings.ToLower(strings.ReplaceAll(*doc.Title, " ", "-"))
			// FIXME: Should create URL-friendly titles for non-english characters
			// Should work with Swedish and Japanese
			argsList = append(argsList, fmt.Sprintf("%s\t%s", *doc.UrlId, *doc.Title))
		}
		return argsList, cobra.ShellCompDirectiveDefault
	}
	return []string{""}, cobra.ShellCompDirectiveError
}

// TODO: URLify titles?
func urlify(c rune) rune {
	runes.In(unicode.Common)
	return c
}

func titleToUrl(title string) string {
	title = strings.ToLower(title) // To lowercase
	title = strings.ReplaceAll(title, " ", "-")
	title = strings.ReplaceAll(title, "å", "a")
	title = strings.ReplaceAll(title, "a", "a")

	return title
}
