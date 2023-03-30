/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "update data",
	Long:  `update existing data in the applecation`,
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		if name == "" {
			fmt.Println("Please enter valid verb.")
			return
		}

		_, ok := index[name]

		if !ok {
			fmt.Printf("%s not found\n", name)
			return
		}

		verb := data[index[name]]

		description, _ := cmd.Flags().GetString("description")

		if description != "" {
			verb.Description = description
		}

		example, _ := cmd.Flags().GetString("example")

		if example != "" {
			if verb.Example1 == "" {
				verb.Example1 = example
			} else {
				verb.Example2 = example
			}
		}

		err := update(&verb)

		if err != nil {
			fmt.Printf("Update %s failure\n", verb.Name)
			return
		}

		fmt.Printf("Update %s success!\n", verb.Name)
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
	updateCmd.Flags().StringP("name", "n", "", "phrasal verb value")
	updateCmd.Flags().StringP("description", "d", "", "phrasal verb definition")
	updateCmd.Flags().StringP("example", "e", "", "example sentence for phrasal verbs.")
}

func update(p *PhrasalVerb) error {
	data[index[p.Name]] = *p

	err := saveCSVFile(CSVFILE)
	if err != nil {
		return err
	}
	return nil
}
