/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// searchCmd represents the search command
var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search verb in data",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		key, _ := cmd.Flags().GetString("key")

		if key == "" {
			fmt.Println("Please enter valid key.")

			return
		}

		//TODO add regex search
		//TODO fulltext search

		_, ok := index[key]

		if ok {
			verb := data[index[key]]
			text, _ := PrettyPrintJSONstream(verb)
			fmt.Println(text)
			return
		}

		fmt.Printf("%s not found\n", key)
	},
}

func init() {
	rootCmd.AddCommand(searchCmd)
	searchCmd.Flags().StringP("key", "k", "", "search key for query verb's name")
}
