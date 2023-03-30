/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"github.com/gocarina/gocsv"
	"github.com/spf13/cobra"
	"os"
)

// insertCmd represents the insert command
var insertCmd = &cobra.Command{
	Use:   "insert",
	Short: "insert new data",
	Long:  `this command inserts new phrasal verb into application`,
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")

		if name == "" {
			fmt.Println("Please enter valid words")
			return
		}

		description, _ := cmd.Flags().GetString("description")

		if description == "" {
			fmt.Println("Please enter valid description")
			return
		}

		example1, _ := cmd.Flags().GetString("example")

		verb := InitVerb(name, description, example1, "")

		if verb == nil {
			fmt.Println("Not a valid record ", verb)
			return
		}

		err := insert(verb)

		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("%s has been added into file\n", verb.Name)
	},
}

func init() {
	rootCmd.AddCommand(insertCmd)
	insertCmd.Flags().StringP("name", "n", "", "phrasal verb value")
	insertCmd.Flags().StringP("description", "d", "", "phrasal verb definition")
	insertCmd.Flags().StringP("example", "e", "", "example sentence for phrasal verbs.")
}

func insert(p *PhrasalVerb) error {
	// check exist or not
	_, ok := index[p.Name]
	if ok {
		return fmt.Errorf("%s already exists. ", p.Name)
	}

	// add to cache
	data = append(data, *p)
	index[p.Name] = len(data) - 1

	// save to file
	err := saveCSVFile(CSVFILE)
	if err != nil {
		return err
	}
	return nil
}

func saveCSVFile(filePath string) error {
	f, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer f.Close()

	err = gocsv.MarshalFile(data, f)

	if err != nil {
		return err
	}

	return nil
}
