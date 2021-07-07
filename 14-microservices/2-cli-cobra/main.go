package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"strings"
)

var (
	localRootFlag   bool
	persistRootFlag bool
	times           int
	rootCmd         = &cobra.Command{
		Use:   "example",
		Short: "An example cobra program",
		Long:  "The long example command description",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("hello from the root command")
		},
	}
	echoCmd = &cobra.Command{
		Use:   "echo [string to echo]",
		Short: "prints given strings to stdout",
		Long:  "The long echo command description",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Echo: " + strings.Join(args, " "))
		},
	}
	timesCmd = &cobra.Command{
		Use:   "times [string to echo]",
		Short: "prints given strings to stdout multiple times",
		Long:  "The long times command description",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			for i := 0; i < times; i++ {
				fmt.Println("Echo: " + strings.Join(args, " "))
			}
		},
	}
)

func init() {
	rootCmd.PersistentFlags().BoolVarP(&persistRootFlag, "persistFlag", "p", false, "a persist root flag")
	rootCmd.Flags().BoolVarP(&localRootFlag, "localFlag", "l", false, "a local root flag")
	timesCmd.Flags().IntVarP(&times, "times", "t", 1, "numbers of times to echo to stdout")

	rootCmd.AddCommand(echoCmd)
	echoCmd.AddCommand(timesCmd)
}

func main() {
	// запуск:
	//
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
