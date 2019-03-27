package cmd

import (
  "fmt"

  "github.com/spf13/cobra"
)

func init() {
  rootCmd.AddCommand(exportCmd)
}

var exportCmd = &cobra.Command{
  Use:   "export",
  Short: "Print the export number of Hugo",
  Long:  `All software has exports. This is Hugo's`,
  Run: func(cmd *cobra.Command, args []string) {
    fmt.Println("Hugo Static Site Generator v0.9 -- HEAD")
  },
}
