package cmd

import (
	"fmt"
	"outline/pkg/api"
	"outline/pkg/editor"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

var draftCmd = &cobra.Command{
	Use:   "draft",
	Short: "Create or edit a draft",
	Long:  `Create or edit a draft`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := conf.Read(); err != nil {
			fmt.Println("No config found, login first.")
			return
		}
		client := api.Client{Config: &conf}
		err := client.Login()
		if err != nil {
			fmt.Printf("Failed to log in: %v", err)
			return
		}

		if len(args) < 1 {
			list, err := client.Documents().Drafts(nil, nil)
			if err != nil {
				fmt.Println(err)
				return
			}
			if len(*list) > 0 {
				for _, doc := range *list {
					fmt.Println(*doc.Title)
				}
			} else {
				fmt.Println("No drafts")
				return
			}
		} else if doc, err := client.Documents().Info(args[0]); err == nil {
			err := setDraftFlags(doc, cmd)
			if err != nil {
				fmt.Printf("Could not update draft: %v", err)
				return
			}
			res, err := editor.EditDoc(doc)
			if err != nil {
				fmt.Printf("Could not edit document: %v", err)
				return
			}

			if res != nil {
				_, err := client.Documents().Update(doc)
				if err != nil {
					fmt.Printf("Could not update document: %v", err)
					return
				}
			}
		} else {
			title := strings.Join(args, " ")
			text := ""
			id := "new_draft"
			doc := api.Document{
				Title: &title,
				Text:  &text,
				Id:    &id,
			}
			err := setDraftFlags(&doc, cmd)
			if err != nil {
				fmt.Printf("Could not create draft: %v", err)
				return
			}

			if doc.CollectionId == nil {
				fmt.Println("You need to specify a collection")
				return
			}

			res, err := editor.EditDoc(&doc)
			if err != nil {
				fmt.Printf("Could not edit document: %v", err)
				return
			}

			if res != nil {
				_, err := client.Documents().Create(&doc)
				if err != nil {
					fmt.Printf("Could not update document: %v", err)
					return
				}
			}
		}

	},
	ValidArgsFunction: completeDrafts,
}

func init() {
	rootCmd.AddCommand(draftCmd)

	draftCmd.Flags().StringP("title", "t", fmt.Sprintf("Draft (%s)", time.Now().String()), "The title of the document")
	draftCmd.Flags().StringP("collection", "c", "", "The collection of the document")
	draftCmd.Flags().BoolP("publish", "p", false, "Publish the document")

	draftCmd.RegisterFlagCompletionFunc("collection", completeCollection)
}

func completeDrafts(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	if err := conf.Read(); err != nil {
		return []string{"error"}, cobra.ShellCompDirectiveDefault
	}
	client := api.Client{Config: &conf}
	docs, err := client.Documents().Drafts(nil, nil)
	if err != nil {
		return []string{"error"}, cobra.ShellCompDirectiveDefault
	}
	argsList := []string{}
	for _, doc := range *docs {
		name := strings.ToLower(strings.ReplaceAll(*doc.Title, " ", "-"))
		argsList = append(argsList, fmt.Sprintf("%s-%s\t%s", name, *doc.UrlId, *doc.UpdatedAt))
	}
	return argsList, cobra.ShellCompDirectiveDefault
}

func completeCollection(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	if err := conf.Read(); err != nil {
		return []string{"error"}, cobra.ShellCompDirectiveDefault
	}
	client := api.Client{Config: &conf}
	cols, err := client.Collections().List()
	if err != nil {
		return []string{"error"}, cobra.ShellCompDirectiveDefault
	}
	argsList := []string{}
	for _, col := range *cols {
		argsList = append(argsList, *col.Name)
	}
	return argsList, cobra.ShellCompDirectiveDefault
}

func setDraftFlags(doc *api.Document, cmd *cobra.Command) error {
	client := api.Client{Config: &conf}
	if cmd.Flag("publish").Changed {
		pub := time.Now()
		doc.PublishedAt = &pub
	}
	if cmd.Flag("title").Changed {
		*doc.Title = cmd.Flag("title").Value.String()
	}
	if cmd.Flag("collection").Changed {
		cols, err := client.Collections().List()
		if err != nil {
			return fmt.Errorf("could not get collections")
		}
		var res *api.Collection
		for _, col := range *cols {
			if *col.Name == cmd.Flag("collection").Value.String() {
				res = &col
				break
			}
		}
		if res == nil {
			return fmt.Errorf("no such collection")
		}
		doc.CollectionId = res.Id
	}
	return nil
}
