/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"time"

	"github.com/arch-spatula/TIL-CLI/utilFn/markdownReadAndWriter"
	"github.com/spf13/cobra"
)

type Weekday int

const (
	Sunday Weekday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

// draftCmd represents the draft command
var draftCmd = &cobra.Command{
	Use:              "draft [OPTIONS]",
	TraverseChildren: true,
	Short:            "draft (today/tomorrow/retro) 3가지 파일 중 1개 생성",
	Long: `3개의 플래그를 조합해서 사용하기 바랍니다.

- 시점 플래그: today / tomorrow / sun
- 회고 플래그: retro
- 회고 유형 플래그: w / m / q

./TIL-CLI draft 		  -> 2311/TIL231207.md
./TIL-CLI sun retro   -> 2311/TIL231210RetroW.md
./TIL-CLI sun retro q -> 2311/TIL231210RetroQ.md

`,
	Run: func(cmd *cobra.Command, args []string) {
		// ./TIL-CLI draft (today)
		if len(args) == 0 {
			markdownReadAndWriter.WriteMarkdown(time.Now())
			return
		}

		if args[0] != "retro" && args[0] != "today" && args[0] != "tomorrow" && args[0] != "sun" {
			panic("없는 플래그")
		}

		// ./TIL-CLI draft (today) retro (w)
		// ./TIL-CLI draft (today) retro m
		// ./TIL-CLI draft (today) retro q
		if args[0] == "retro" {
			if len(args) <= 1 {
				markdownReadAndWriter.WriteRetro(time.Now(), "W")
				return
			}
			switch args[1] {
			case "w":
				markdownReadAndWriter.WriteRetro(time.Now(), "W")
			case "m":
				markdownReadAndWriter.WriteRetro(time.Now(), "M")
			case "q":
				markdownReadAndWriter.WriteRetro(time.Now(), "Q")
			default:
				panic("없는 플래그입니다.")
			}
		}

		// ./TIL-CLI draft today
		// ./TIL-CLI draft today retro (w)
		// ./TIL-CLI draft today retro m
		// ./TIL-CLI draft today retro q
		if args[0] == "today" {
			if len(args) == 1 {
				markdownReadAndWriter.WriteMarkdown(time.Now())
				return
			} else if len(args) == 2 {
				if args[1] != "retro" {
					panic("없는 명령입니다.")
				}
				markdownReadAndWriter.WriteRetro(time.Now(), "W")
			} else if len(args) == 3 {
				switch args[2] {
				case "w":
					markdownReadAndWriter.WriteRetro(time.Now(), "W")
				case "m":
					markdownReadAndWriter.WriteRetro(time.Now(), "M")
				case "q":
					markdownReadAndWriter.WriteRetro(time.Now(), "Q")
				default:
					panic("없는 플래그입니다.")
				}
			}
		}

		// ./TIL-CLI draft tomorrow
		// ./TIL-CLI draft tomorrow retro (w)
		// ./TIL-CLI draft tomorrow retro m
		// ./TIL-CLI draft tomorrow retro q
		if args[0] == "tomorrow" {
			if len(args) == 1 {
				markdownReadAndWriter.WriteMarkdown(time.Now().AddDate(0, 0, 1))
				return
			} else if len(args) == 2 {
				if args[1] != "retro" {
					panic("없는 명령입니다.")
				}
				markdownReadAndWriter.WriteRetro(time.Now().AddDate(0, 0, 1), "W")
			} else if len(args) == 3 {
				switch args[2] {
				case "w":
					markdownReadAndWriter.WriteRetro(time.Now().AddDate(0, 0, 1), "W")
				case "m":
					markdownReadAndWriter.WriteRetro(time.Now().AddDate(0, 0, 1), "M")
				case "q":
					markdownReadAndWriter.WriteRetro(time.Now().AddDate(0, 0, 1), "Q")
				default:
					panic("없는 플래그입니다.")
				}
			}
		}

		// ./TIL-CLI draft sun
		// ./TIL-CLI draft sun retro (w)
		// ./TIL-CLI draft sun retro m
		// ./TIL-CLI draft sun retro q
		if args[0] == "sun" {
			if len(args) == 1 {
				switch time.Now().Weekday() {
				case time.Sunday:
					markdownReadAndWriter.WriteMarkdown(time.Now())
				default:
					markdownReadAndWriter.WriteMarkdown(time.Now().AddDate(0, 0, 7-int(time.Now().Weekday())))
				}
				return
			} else if len(args) == 2 {
				if args[1] != "retro" {
					panic("없는 명령입니다.")
				}
				switch time.Now().Weekday() {
				case time.Sunday:
					markdownReadAndWriter.WriteRetro(time.Now(), "W")
				default:
					markdownReadAndWriter.WriteRetro(time.Now().AddDate(0, 0, 7-int(time.Now().Weekday())), "W")
				}
			} else if len(args) == 3 {
				switch args[2] {
				case "w":
					switch time.Now().Weekday() {
					case time.Sunday:
						markdownReadAndWriter.WriteRetro(time.Now(), "W")
					default:
						markdownReadAndWriter.WriteRetro(time.Now().AddDate(0, 0, 7-int(time.Now().Weekday())), "W")
					}
				case "m":
					switch time.Now().Weekday() {
					case time.Sunday:
						markdownReadAndWriter.WriteRetro(time.Now(), "M")
					default:
						markdownReadAndWriter.WriteRetro(time.Now().AddDate(0, 0, 7-int(time.Now().Weekday())), "M")
					}
					markdownReadAndWriter.WriteRetro(time.Now().AddDate(0, 0, 1), "M")
				case "q":
					switch time.Now().Weekday() {
					case time.Sunday:
						markdownReadAndWriter.WriteRetro(time.Now(), "Q")
					default:
						markdownReadAndWriter.WriteRetro(time.Now().AddDate(0, 0, 7-int(time.Now().Weekday())), "Q")
					}
				default:
					panic("없는 플래그입니다.")
				}
			}
		}

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
