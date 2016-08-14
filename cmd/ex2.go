// Copyright © 2016 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// ex2Cmd represents the ex2 command
var ex2Cmd = &cobra.Command{
	Use:   "ex2",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		sum := fib(0, 1)
		fmt.Printf("Sum: %d\n", sum)
	},
}

func fib(a, b uint64) uint64 {
	var sum uint64 = 0

	if b%2 == 0 {
		sum += b
	}
	result := a + b
	if result > 4000000 {
		return sum
	} else {
		sum += fib(b, result)
		return sum
	}
}

func init() {
	RootCmd.AddCommand(ex2Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// ex2Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// ex2Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
