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
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
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
