/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"sort"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list all verbs",
	Long:  `This command lists all verbs in the phrasal verbs application.`,
	Run: func(cmd *cobra.Command, args []string) {
		list()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}

func list() {
	sort.Sort(PhrasalVerbs(data))
	text, err := PrettyPrintJSONstream(data)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(text)
	fmt.Printf("%d data in Phrasal verbs\n", len(data))
}

func PrettyPrintJSONstream(data interface{}) (string, error) {
	buffer := new(bytes.Buffer)
	encoder := json.NewEncoder(buffer)
	encoder.SetIndent("", "\t")

	err := encoder.Encode(data)
	if err != nil {
		return "", err
	}

	return buffer.String(), nil
}
