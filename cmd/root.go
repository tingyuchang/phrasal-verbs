/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"github.com/gocarina/gocsv"
	"io"
	"os"

	"github.com/spf13/cobra"
)

// JSONFILE resides in the current directory
var CSVFILE = "./data.csv"

var index map[string]int
var data = PhrasalVerbs{}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "phrasal-verbs",
	Short: "A simple CLI tool for phrasal verbs dictionary ",
	Long:  `Support insert, delete, update, search ... feature and using csv to store data`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := setCSVFile()
	if err != nil {
		fmt.Println(err)
		return
	}

	err = readCSVFile(CSVFILE)

	if err != nil && err != io.EOF {
		fmt.Println(err)
		return
	}

	createIndex()

	cobra.CheckErr(rootCmd.Execute())
}

func init() {}

func setCSVFile() error {
	// don't support switch data source
	//filePath := os.Getenv("")

	_, err := os.Stat(CSVFILE)
	if err != nil {
		fmt.Println("Creating: ", CSVFILE)
		f, err := os.Create(CSVFILE)
		if err != nil {
			f.Close()
			return err
		}
		defer f.Close()
	}
	fileInfo, err := os.Stat(CSVFILE)

	mode := fileInfo.Mode()

	if !mode.IsRegular() {
		return fmt.Errorf("%s not a regualr file", CSVFILE)
	}
	return nil
}

func readCSVFile(filePath string) error {
	_, err := os.Stat(filePath)
	if err != nil {
		return err
	}

	f, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer f.Close()

	err = gocsv.UnmarshalFile(f, &data)
	if err != nil {
		return err
	}
	return nil
}

func createIndex() {
	index = make(map[string]int)
	for i, v := range data {
		key := v.Name
		index[key] = i
	}
}
