// Copyright © 2018 Rafal Korkosz <korkosz.rafal@gmail.com>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package cmd

import (
	"bufio"
	"errors"
	"os"
	"strings"

	"github.com/rkorkosz/kv/common"
	"github.com/spf13/cobra"
)

// parseCmd represents the parse command
var parseCmd = &cobra.Command{
	Use:   "parse",
	Short: "Parse and store env file",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("provide file to parse")
		}
		file, err := os.Open(args[0])
		if err != nil {
			return err
		}
		defer file.Close()
		store := common.NewStore()
		defer store.Close()
		var envs [][2]string
		var row string
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			row = scanner.Text()
			splitted := strings.Split(row, "=")
			envs = append(envs, [2]string{splitted[0], splitted[1]})
		}
		return store.SetMany(envs)
	},
}

func init() {
	rootCmd.AddCommand(parseCmd)
}
