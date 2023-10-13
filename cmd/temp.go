/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// tempCmd represents the temp command
var tempCmd = &cobra.Command{
	Use:   "temp",
	Short: "A brief description of your command",
	Long: `템플릿에 해당하는 마크다운을 생성합니다.

템플릿은 single source of truth에 해당합니다.

템플릿이 없으면 기본 생성해줄 것입니다.`,
	Run: func(cmd *cobra.Command, args []string) {
		if _, err := os.Stat("template.md"); errors.Is(err, os.ErrNotExist) {
			// path/to/whatever does not exist
			fmt.Println("template.md이 존재하지 않아 임시로 만들어두겠습니다.")
		}
	},
}

func init() {
	rootCmd.AddCommand(tempCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// tempCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// tempCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
