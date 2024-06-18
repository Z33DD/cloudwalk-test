/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	i "cloudwalk-test/internal"

	"github.com/spf13/cobra"
)

// parseCmd represents the parse command
var parseCmd = &cobra.Command{
	Use:   "parse",
	Short: "Parse a log file to get some analytics",
	Long:  `lorem ipsum dolor sit amet`,
	Run: func(cmd *cobra.Command, args []string) {
		fileFlag, err := cmd.Flags().GetString("file")
		if err != nil {
			fmt.Println("Error getting file flag")
			panic(err)
		}
		logParser := i.LogParser{FilePath: fileFlag}
		logParser.Parse()

		fmt.Println("parse called")
	},
}

func init() {
	parseCmd.Flags().StringP("file", "f", "", "Specify the log file to parse")
	rootCmd.AddCommand(parseCmd)
}
