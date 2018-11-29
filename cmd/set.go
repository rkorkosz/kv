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
	"errors"
	"strings"

	"github.com/rkorkosz/kv/common"
	"github.com/spf13/cobra"
)

func validate(args []string) error {
	if len(args) < 1 {
		return errors.New("key and value needs to be provided")
	}
	if len(args) == 1 && !strings.Contains(args[0], "=") {
		return errors.New("value must be in form `key=value`")
	}
	if len(args) > 2 {
		return errors.New("you should provide only key and value")
	}
	return nil
}

// setCmd represents the store command
var setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set value",
	RunE: func(cmd *cobra.Command, args []string) error {
		err := validate(args)
		if err != nil {
			return err
		}
		var env [2]string
		if len(args) == 1 {
			splitted := strings.Split(args[0], "=")
			env[0] = splitted[0]
			env[1] = splitted[1]
		} else {
			env[0] = args[0]
			env[1] = args[1]
		}
		store := common.NewStore()
		defer store.Close()
		return store.Set(env)
	},
}

func init() {
	rootCmd.AddCommand(setCmd)
}
