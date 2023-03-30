/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "delete verb in data",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		key, _ := cmd.Flags().GetString("key")

		if key == "" {
			fmt.Println("Please enter valid key")
			return
		}

		_, ok := index[key]

		if !ok {
			fmt.Println(key, " not found")
			return
		}

		err := deleteRecord(key)

		if err != nil {
			fmt.Println("Delete failed ", err)
			return
		}
		fmt.Printf("%s has been deleted. \n", key)

	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
	deleteCmd.Flags().StringP("key", "k", "", "search key for query verb's name")
}

func deleteRecord(key string) error {
	// remove form data
	data[index[key]] = data[len(data)-1]
	data = data[:len(data)-1]
	//remove index
	delete(index, key)

	// rebuild index
	createIndex()
	// save to file
	err := saveCSVFile(CSVFILE)

	if err != nil {
		return err
	}
	return nil
}
