/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
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
			panic(err)
		}
		outputFlag, err := cmd.Flags().GetString("output")
		if err != nil {
			panic(err)
		}
		logParser := i.LogParser{FilePath: fileFlag}
		logParser.Parse(outputFlag)
	},
}

func init() {
	parseCmd.Flags().StringP("file", "f", "", "Specify the log file to parse")
	parseCmd.MarkFlagRequired("file")
	parseCmd.Flags().StringP("output", "o", "", "Specify the output file")
	parseCmd.MarkFlagRequired("output")
	rootCmd.AddCommand(parseCmd)
}
