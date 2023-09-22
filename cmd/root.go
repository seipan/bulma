// MIT License

// Copyright (c) 2023 Yamasaki Shotaro

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:

// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package cmd

import (
	"log"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "bluma",
	Short: "CLI tool to parse OpenAPI and stress test each endpoint.",
	Long:  `CLI tool to parse OpenAPI and stress test each endpoint..`,
	Run: func(cmd *cobra.Command, args []string) {
		path, err := cmd.Flags().GetString("path")
		if err != nil {
			log.Println(err)
		}
		base, err := cmd.Flags().GetString("base")
		if err != nil {
			log.Println(err)
		}
		freq, err := cmd.Flags().GetInt("frequency")
		if err != nil {
			log.Println(err)
		}
		duration, err := cmd.Flags().GetDuration("duration")
		if err != nil {
			log.Println(err)
		}
		err = ParseAndAttack(cmd.Context(), base, path, freq, duration)
		if err != nil {
			log.Println(err)
		}
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringP("path", "p", "", "FilePath for Parsing OpenAPI")
	rootCmd.Flags().StringP("base", "b", "", "BaseURL for stress test")
	rootCmd.Flags().IntP("frequency", "f", 1, "stress test frequency")
	rootCmd.Flags().DurationP("duration", "d", 1, "stress test duration")
}
