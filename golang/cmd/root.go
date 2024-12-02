package cmd

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strings"

	"github.com/spf13/cobra"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "generated code example",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() string {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	var cmdFlag = rootCmd.Flag("day")
	if cmdFlag == nil || len(cmdFlag.Value.String()) == 0 {
		fmt.Println("No proper day has been set")
		os.Exit(127)
	}
	return cmdFlag.Value.String()
}

func init() {
	cobra.OnInitialize()

	rootCmd.PersistentFlags().StringP("day", "d", "", "Day of the AoC 2024")
}

func ReadInputFile(daySelector string) string {
	filename := fmt.Sprintf("inputs/day_%s.txt", daySelector)
	content, err := os.ReadFile(filename)

	if err != nil {
		log.Fatal(err)
	}
	return string(content)
}

func ReadInputFileGood(daySelector string) []string {
	filename := fmt.Sprintf("inputs/day_%s.txt", daySelector)
	content, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	var lines = strings.Split(string(content), "\n")
	return lines
}

func SortColumn(column []int) []int {
	//var sortedColumn []int
	sort.Slice(column, func(i, j int) bool {
		return column[i] < column[j]
	})
	return column
}
