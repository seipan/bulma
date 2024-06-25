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
	"fmt"
	"os"

	"github.com/seipan/bulma/lib"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "bluma",
	Short: "CLI tool to parse OpenAPI and stress test each endpoint.",
	Long:  `CLI tool to parse OpenAPI and stress test each endpoint..`,
	RunE: func(cmd *cobra.Command, args []string) error {
		path, err := cmd.Flags().GetString("path")
		if err != nil {
			return fmt.Errorf("failed to get path: %w", err)
		}
		base, err := cmd.Flags().GetString("base")
		if err != nil {
			return fmt.Errorf("failed to get base: %w", err)
		}
		freq, err := cmd.Flags().GetInt("frequency")
		if err != nil {
			return fmt.Errorf("failed to get frequency: %w", err)
		}
		duration, err := cmd.Flags().GetDuration("duration")
		if err != nil {
			return fmt.Errorf("failed to get duration: %w", err)
		}
		ignore, err := cmd.Flags().GetStringArray("ignore")
		if err != nil {
			return fmt.Errorf("failed to get ignore paths: %w", err)
		}
		err = lib.ParseAndAttack(cmd.Context(), ignore, base, path, freq, duration)
		if err != nil {
			return fmt.Errorf("failed to parse and attack: %w", err)
		}
		return nil
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
	rootCmd.Flags().StringArrayP("ignore", "i", []string{}, "Ignore Paths for stress test")
	rootCmd.Flags().IntP("frequency", "f", 1, "stress test frequency")
	rootCmd.Flags().DurationP("duration", "d", 1, "stress test duration")
}
