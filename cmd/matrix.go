/*
Copyright © 2021 Doug Hellmann <doug@doughellmann.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/dhellmann/playfair/cipher"
)

// matrixCmd represents the matrix command
var matrixCmd = &cobra.Command{
	Use:   "matrix",
	Short: "Generate the cipher matrix for a keyword",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("matrix requires exactly one keyword argument")
		}
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		m, err := cipher.New(args[0])
		if err != nil {
			return errors.Wrap(err, "failed to create matrix")
		}
		fmt.Printf("%s\n", m.Matrix())
		return nil
	},
}

func init() {
	rootCmd.AddCommand(matrixCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// matrixCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// matrixCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
