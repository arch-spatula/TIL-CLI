/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// draftCmd represents the draft command
var draftCmd = &cobra.Command{
	Use:              "draft [OPTIONS]",
	TraverseChildren: true,
	Short:            "draft -(today/tomorrow/retro) 3가지 파일 중 1개 생성",
	Long: `

- today:    오늘 TIL 템플릿을 생성합니다. 2311/TIL231112.md
- tomorrow: 내일 TIL 템플릿을 생성합니다. 2311/TIL231113.md
- retro:    회고 TIL 템플릿을 생성합니다. 2311/TIL231112Retro.md`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("draft called", args)
	},
}

func init() {
	rootCmd.AddCommand(draftCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// draftCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// draftCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
